package main

import (
	"errors"

	"github.com/spf13/cobra"

	"updep/pkg/app"

	tea "charm.land/bubbletea/v2"
)

var rootCmd = &cobra.Command{
	Use:     "updep",
	Version: "0.0.1",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		value, err := cmd.Flags().GetString("pm")
		if err != nil {
			return err
		}
		switch value {
		case "npm", "pnpm", "yarn", "bun":
		// TODO: set package manager
		case "":
		default:
			return errors.New("invalid package manager")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		p := tea.NewProgram(app.New())
		defer p.Kill()
		_, err := p.Run()
		return err
	},
}

func init() {
	rootCmd.PersistentFlags().String("pm", "", "the package manager to use")
}
