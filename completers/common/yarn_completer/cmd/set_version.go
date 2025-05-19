package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var set_versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Lock the Yarn version used by the project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(set_versionCmd).Standalone()

	set_versionCmd.Flags().Bool("only-if-needed", false, "Only lock the Yarn version if it isn't already locked")
	setCmd.AddCommand(set_versionCmd)

	// TODO positional completion
}
