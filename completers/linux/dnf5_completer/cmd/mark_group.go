package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf5_completer/cmd/action"
	"github.com/spf13/cobra"
)

var markGroupCmd = &cobra.Command{
	Use:   "group [options] <group_id> <package-spec>...",
	Short: "mark package as installed by a group",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(markGroupCmd).Standalone()

	markCmd.AddCommand(markGroupCmd)

	carapace.Gen(markGroupCmd).PositionalCompletion(
		carapace.ActionValues(),
		action.ActionInstalledPackages(markGroupCmd),
	)
}
