package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var submodule_initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the submodules recorded in the index",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(submodule_initCmd).Standalone()

	submoduleCmd.AddCommand(submodule_initCmd)

	carapace.Gen(submodule_initCmd).PositionalAnyCompletion(
		git.ActionSubmodulePaths().FilterArgs(),
	)
}
