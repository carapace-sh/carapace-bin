package cmd

import (
	"github.com/spf13/cobra"
)

var repoCmd = &cobra.Command{
	Use:   "repo <command>",
	Short: "Create, clone, fork, and view repositories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(repoCmd)
}
