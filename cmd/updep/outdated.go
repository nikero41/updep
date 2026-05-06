package main

import "github.com/spf13/cobra"

var outdatedCmd = &cobra.Command{
	Use: "outdated",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: implement
		println("TODO: outdated")
	},
}

func init() {
	rootCmd.AddCommand(outdatedCmd)
}
