package packagemanager

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	Npm = iota
	Yarn
	Pnpm
	Bun
)

type PackageManager int

func New() PackageManager {
	// search for package-lock.json, yarn.lock, pnpm.lock, bun.lock
	return PackageManager(Npm)
}

func (pm PackageManager) Update(packages []Package) error {
	return errors.New("not implemented")
}

type JSONPackage struct {
	Name    string
	Wanted  string `json:"wanted"`
	Latest  string `json:"latest"`
	Current string `json:"current"`
}

func (pm PackageManager) GetOutdated() (map[string]JSONPackage, error) {
	// timer := time.NewTimer(time.Second * 1)
	// <-timer.C
	// var err error

	// output, err := exec.Command("npm", "outdated", "--json").Output()
	output, err := os.ReadFile("output.json")
	if err != nil {
		fmt.Println("🪚 err:", err)
	}

	var outdated map[string]JSONPackage
	err = json.NewDecoder(strings.NewReader(string(output))).Decode(&outdated)
	if err != nil {
		return nil, err
	}

	return outdated, nil
}

var output string = `
	{
	  "react-native-reanimated": {
    "current": "4.1.0",
    "wanted": "4.1.3",
    "latest": "4.1.3"
  },
  "react-native-screens": {
    "current": "4.16.0",
    "wanted": "4.17.1",
    "latest": "4.17.1"
  },
  "react-native-svg": {
    "current": "15.13.0",
    "wanted": "15.14.0",
    "latest": "16.14.0"
  },
  "react-native-worklets": {
    "current": "0.5.1",
    "wanted": "0.5.1",
    "latest": "0.6.1"
  }
}`
