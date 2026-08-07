package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securefile_updateCmd = &cobra.Command{
	Use:     "update <name> <path>",
	Short:   "Update a secure file in a project.",
	Aliases: []string{"overwrite"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securefile_updateCmd).Standalone()

	securefile_updateCmd.Flags().BoolP("yes", "y", false, "Skip the confirmation prompt.")
	securefileCmd.AddCommand(securefile_updateCmd)
}
