package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/apk_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/apk_completer/cmd/common"
	"github.com/spf13/cobra"
)

var delCmd = &cobra.Command{
	Use:     "del",
	Short:   "Remove constraints from WORLD and commit changes",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "package installation and removal",
}

func init() {
	carapace.Gen(delCmd).Standalone()

	delCmd.Flags().BoolP("rdepends", "r", false, "Recursively delete all top-level reverse dependencies, too")
	common.AddGlobalFlags(delCmd)
	common.AddCommitFlags(delCmd)
	rootCmd.AddCommand(delCmd)

	carapace.Gen(delCmd).PositionalAnyCompletion(
		action.ActionPackages(delCmd, true).FilterArgs(),
	)
}
