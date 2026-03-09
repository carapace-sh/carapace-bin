package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mark_groupCmd = &cobra.Command{
	Use:   "group <group-id> <package-spec>...",
	Short: "mark package as installed by group",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mark_groupCmd).Standalone()

	mark_groupCmd.Flags().Bool("skip-unavailable", false, "allow skipping packages that are not installed")

	markCmd.AddCommand(mark_groupCmd)

	carapace.Gen(mark_groupCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(mark_groupCmd),
	)
}
