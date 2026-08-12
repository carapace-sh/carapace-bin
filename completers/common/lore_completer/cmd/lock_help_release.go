package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lock_help_releaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Release lock on file(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lock_help_releaseCmd).Standalone()

	lock_helpCmd.AddCommand(lock_help_releaseCmd)
}
