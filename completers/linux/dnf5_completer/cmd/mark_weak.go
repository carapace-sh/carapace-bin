package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf5_completer/cmd/action"
	"github.com/spf13/cobra"
)

var markWeakCmd = &cobra.Command{
	Use:   "weak [options] <package-spec>...",
	Short: "mark package as a weak dependency",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(markWeakCmd).Standalone()

	markCmd.AddCommand(markWeakCmd)

	carapace.Gen(markWeakCmd).PositionalAnyCompletion(
		action.ActionInstalledPackages(markWeakCmd),
	)
}
