package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var codespace_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a codespace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codespace_deleteCmd).Standalone()
	codespace_deleteCmd.Flags().Bool("all", false, "Delete all codespaces")
	codespace_deleteCmd.Flags().StringP("codespace", "c", "", "Name of the codespace")
	codespace_deleteCmd.Flags().Uint16("days", 0, "Delete codespaces older than `N` days")
	codespace_deleteCmd.Flags().BoolP("force", "f", false, "Skip confirmation for codespaces that contain unsaved changes")
	codespace_deleteCmd.Flags().StringP("repo", "r", "", "Delete codespaces for a `repository`")
	codespaceCmd.AddCommand(codespace_deleteCmd)

	carapace.Gen(codespace_deleteCmd).FlagCompletion(carapace.ActionMap{
		"codespace": action.ActionCodespaces(),
		"repo":      action.ActionOwnerRepositories(codespace_deleteCmd),
	})
}
