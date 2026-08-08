package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf5_completer/cmd/action"
	"github.com/spf13/cobra"
)

var versionlockExcludeCmd = &cobra.Command{
	Use:   "exclude [options] <package-spec>...",
	Short: "add new exclude entry to versionlock configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionlockExcludeCmd).Standalone()

	versionlockCmd.AddCommand(versionlockExcludeCmd)

	carapace.Gen(versionlockExcludeCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(versionlockExcludeCmd),
	)
}
