package main

import (
	"errors"

	"github.com/spf13/cobra"

	"updep/pkg/app"
	packagemanager "updep/pkg/packageManager"

	tea "charm.land/bubbletea/v2"
)

var rootCmd = &cobra.Command{
	Use:     "updep",
	Version: "0.0.1",
	RunE: func(cmd *cobra.Command, args []string) error {
		value, err := cmd.Flags().GetString("pm")
		if err != nil {
			return err
		}
		var pm packagemanager.PackageManager
		switch value {
		case "npm", "pnpm", "yarn", "bun":
			pm, err = packagemanager.New(value)
			if err != nil {
				return err
			}
		case "":
		default:
			return errors.New("invalid package manager")
		}

		p := tea.NewProgram(app.New(pm))
		defer p.Kill()
		_, err = p.Run()
		return err
	},
}

func init() {
	rootCmd.PersistentFlags().String("pm", "", "the package manager to use")
}
