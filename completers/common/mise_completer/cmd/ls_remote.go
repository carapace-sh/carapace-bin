package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var ls_remoteCmd = &cobra.Command{
	Use:   "ls-remote",
	Short: "List runtime versions available for install",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ls_remoteCmd).Standalone()

	ls_remoteCmd.Flags().Bool("all", false, "Show all versions")
	ls_remoteCmd.Flags().Bool("all-versions", false, "Show all versions including non-semver ones")
	rootCmd.AddCommand(ls_remoteCmd)

	carapace.Gen(ls_remoteCmd).PositionalCompletion(
		action.ActionTools(),
	)
}
