package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var submodule_deinitCmd = &cobra.Command{
	Use:   "deinit",
	Short: "Unregister the given submodules",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(submodule_deinitCmd).Standalone()
	submodule_deinitCmd.Flags().Bool("all", false, "remove the whole submodule section from the config as well")
	submodule_deinitCmd.Flags().Bool("force", false, "remove working tree even if it contains local modifications")

	submoduleCmd.AddCommand(submodule_deinitCmd)

	carapace.Gen(submodule_deinitCmd).PositionalAnyCompletion(
		git.ActionSubmodulePaths().FilterArgs(),
	)
}
