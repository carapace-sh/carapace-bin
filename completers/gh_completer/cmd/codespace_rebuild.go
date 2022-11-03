package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var codespace_rebuildCmd = &cobra.Command{
	Use:   "rebuild",
	Short: "Rebuild a codespace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codespace_rebuildCmd).Standalone()
	codespace_rebuildCmd.Flags().StringP("codespace", "c", "", "name of the codespace")
	codespace_rebuildCmd.Flags().Bool("full", false, "perform a full rebuild")
	codespaceCmd.AddCommand(codespace_rebuildCmd)

	carapace.Gen(codespace_rebuildCmd).FlagCompletion(carapace.ActionMap{
		"codespace": action.ActionCodespaces(),
	})
}
