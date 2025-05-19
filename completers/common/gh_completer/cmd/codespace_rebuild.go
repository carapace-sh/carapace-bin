package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var codespace_rebuildCmd = &cobra.Command{
	Use:   "rebuild",
	Short: "Rebuild a codespace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codespace_rebuildCmd).Standalone()

	codespace_rebuildCmd.PersistentFlags().StringP("codespace", "c", "", "Name of the codespace")
	codespace_rebuildCmd.Flags().Bool("full", false, "Perform a full rebuild")
	codespace_rebuildCmd.PersistentFlags().StringP("repo", "R", "", "Filter codespace selection by repository name (user/repo)")
	codespace_rebuildCmd.PersistentFlags().String("repo-owner", "", "Filter codespace selection by repository owner (username or org)")
	codespaceCmd.AddCommand(codespace_rebuildCmd)

	carapace.Gen(codespace_rebuildCmd).FlagCompletion(carapace.ActionMap{
		"codespace": action.ActionCodespaces(),
		"repo":      gh.ActionOwnerRepositories(gh.HostOpts{}),
	})
}
