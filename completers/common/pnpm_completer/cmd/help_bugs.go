package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_bugsCmd = &cobra.Command{
	Use:   "bugs",
	Short: "Opens the bug tracker URL of a package in the default browser",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_bugsCmd).Standalone()

	helpCmd.AddCommand(help_bugsCmd)
}
