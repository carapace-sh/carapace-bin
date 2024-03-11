package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var version_checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check that all the relevant packages have been bumped",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(version_checkCmd).Standalone()

	version_checkCmd.Flags().BoolP("interactive", "i", false, "Open an interactive interface used to set version bumps")
	versionCmd.AddCommand(version_checkCmd)
}
