package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securefile_createCmd = &cobra.Command{
	Use:     "create <name> <path>",
	Short:   "Upload a new secure file to a project.",
	Aliases: []string{"upload"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securefile_createCmd).Standalone()

	securefileCmd.AddCommand(securefile_createCmd)

	carapace.Gen(securefile_createCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
