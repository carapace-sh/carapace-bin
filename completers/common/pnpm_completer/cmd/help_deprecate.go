package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_deprecateCmd = &cobra.Command{
	Use:   "deprecate",
	Short: "Deprecates a version of a package in the registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_deprecateCmd).Standalone()

	helpCmd.AddCommand(help_deprecateCmd)
}
