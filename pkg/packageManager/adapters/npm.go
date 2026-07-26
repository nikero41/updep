package adapters

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os/exec"
	"strings"
	"sync"

	"updep/pkg/dependency"
	"updep/pkg/packageManager/events"

	"github.com/go-playground/validator/v10"
)

type Npm struct {
	name string
}

func NewNpm() *Npm {
	return &Npm{name: "npm"}
}

var validate = validator.New(validator.WithRequiredStructEnabled())

type JSONPackage struct {
	Wanted   string `validate:"required,semver" json:"wanted"`
	Latest   string `validate:"required,semver" json:"latest"`
	Current  string `validate:"required,semver" json:"current"`
	Homepage string `validate:"omitempty,url"   json:"homepage"`
}

func (pm Npm) Name() string { return pm.name }

func (pm Npm) GetOutdated() ([]dependency.Dependency, error) {
	output, err := exec.Command("npm", "outdated", "--json", "--long").Output()

	// npm outdated returns exit status 1 if there are outdated packages
	if err != nil && err.Error() != "exit status 1" {
		return nil, err
	}

	var outdated map[string]JSONPackage
	err = json.NewDecoder(strings.NewReader(string(output))).Decode(&outdated)
	if err != nil {
		return nil, err
	}

	for _, d := range outdated {
		err = validate.Struct(d)
		if err != nil {
			return nil, err
		}
	}

	outdatedDeps := make([]dependency.Dependency, len(outdated))

	var index int
	for name, value := range outdated {
		d, err := dependency.New(
			name,
			value.Wanted,
			value.Latest,
			value.Current,
			value.Homepage,
		)
		if err != nil {
			return nil, fmt.Errorf("invalid package versions: %v %w", value, err)
		}

		outdatedDeps[index] = *d
		index++
	}

	return outdatedDeps, nil
}

func (pm Npm) Update(deps []dependency.Dependency, outputChan chan<- events.PmOutputEvent) {
	args := make([]string, len(deps)+3)
	for i, d := range deps {
		args[i] = d.Name
		if d.Target == dependency.Latest {
			args[i] += "@latest"
		}
	}
	args = append([]string{"--color=always", "install"}, args...)
	slog.Info("install args:", "args", args)

	cmd := exec.Command("npm", args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go scan(&wg, outputChan, stdout)
	go scan(&wg, outputChan, stderr)
	go func() {
		wg.Wait()
		outputChan <- events.PmOutputEvent{
			Err:  nil,
			Done: true,
		}
		close(outputChan)
	}()

	_ = cmd.Start()
}

func scan(wg *sync.WaitGroup, c chan<- events.PmOutputEvent, r io.ReadCloser) {
	defer wg.Done()
	s := bufio.NewScanner(r)

	for s.Scan() {
		c <- events.PmOutputEvent{Output: s.Text()}
	}

	if err := s.Err(); err != nil {
		c <- events.PmOutputEvent{
			Err:  s.Err(),
			Done: false,
		}
	}
}

func (pm Npm) String() string {
	return pm.Name()
}
