package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mark_userCmd = &cobra.Command{
	Use:   "user <package-spec>...",
	Short: "mark package as user-installed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mark_userCmd).Standalone()

	mark_userCmd.Flags().Bool("skip-unavailable", false, "allow skipping packages that are not installed")

	markCmd.AddCommand(mark_userCmd)

	carapace.Gen(mark_userCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(mark_userCmd),
	)
}
