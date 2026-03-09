package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mark_dependencyCmd = &cobra.Command{
	Use:   "dependency <package-spec>...",
	Short: "mark package as a dependency",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mark_dependencyCmd).Standalone()

	mark_dependencyCmd.Flags().Bool("skip-unavailable", false, "allow skipping packages that are not installed")

	markCmd.AddCommand(mark_dependencyCmd)

	carapace.Gen(mark_dependencyCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(mark_dependencyCmd),
	)
}
