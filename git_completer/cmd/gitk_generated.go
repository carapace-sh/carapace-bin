package cmd

import (
	"github.com/spf13/cobra"
)

var gitkCmd = &cobra.Command{
	Use: "gitk",
	Short: "The Git repository browser",
    Run: func(cmd *cobra.Command, args []string) {
    },
}

func init() {

    rootCmd.AddCommand(gitkCmd)
}
