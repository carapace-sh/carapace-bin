package cmd

import (
	"github.com/spf13/cobra"
)

var gistCmd = &cobra.Command{
	Use:   "gist <command>",
	Short: "Manage gists",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(gistCmd)
}
