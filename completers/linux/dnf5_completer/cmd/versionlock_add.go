package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf5_completer/cmd/action"
	"github.com/spf13/cobra"
)

var versionlockAddCmd = &cobra.Command{
	Use:   "add [options] <package-spec>...",
	Short: "add new entry to versionlock configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionlockAddCmd).Standalone()

	versionlockCmd.AddCommand(versionlockAddCmd)

	carapace.Gen(versionlockAddCmd).PositionalAnyCompletion(
		action.ActionInstalledPackages(versionlockAddCmd),
	)
}
