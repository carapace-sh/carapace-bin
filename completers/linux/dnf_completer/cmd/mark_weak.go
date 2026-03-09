package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mark_weakCmd = &cobra.Command{
	Use:   "weak <package-spec>...",
	Short: "mark package as a weak dependency",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mark_weakCmd).Standalone()

	mark_weakCmd.Flags().Bool("skip-unavailable", false, "allow skipping packages that are not installed")

	markCmd.AddCommand(mark_weakCmd)

	carapace.Gen(mark_weakCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(mark_weakCmd),
	)
}
