package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf5_completer/cmd/action"
	"github.com/spf13/cobra"
)

var versionlockDeleteCmd = &cobra.Command{
	Use:   "delete [options] <package-spec>...",
	Short: "remove any matching versionlock configuration entries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionlockDeleteCmd).Standalone()

	versionlockCmd.AddCommand(versionlockDeleteCmd)

	carapace.Gen(versionlockDeleteCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(versionlockDeleteCmd),
	)
}
