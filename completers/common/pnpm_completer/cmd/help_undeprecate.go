package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_undeprecateCmd = &cobra.Command{
	Use:   "undeprecate",
	Short: "Removes deprecation from a version of a package in the registry. Only works on already deprecated versions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_undeprecateCmd).Standalone()

	helpCmd.AddCommand(help_undeprecateCmd)
}
