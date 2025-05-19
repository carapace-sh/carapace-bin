package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securefile_createCmd = &cobra.Command{
	Use:     "create <fileName> <inputFilePath>",
	Short:   "Create a new project secure file.",
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
