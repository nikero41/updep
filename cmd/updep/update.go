package main

import (
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [flags] [package]...",
	Short: "Update packages",
	Example: `
# Update all packages
updep update

# Update a specific package
updep update package
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: implement
		println("TODO: update")
	},
}

func init() {
	updateCmd.Flags().BoolP("all", "a", false, "update all packages")

	rootCmd.AddCommand(updateCmd)
}
