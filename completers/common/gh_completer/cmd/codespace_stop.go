package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var codespace_stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop a running codespace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codespace_stopCmd).Standalone()

	codespace_stopCmd.PersistentFlags().StringP("codespace", "c", "", "Name of the codespace")
	codespace_stopCmd.Flags().StringP("org", "o", "", "The `login` handle of the organization (admin-only)")
	codespace_stopCmd.PersistentFlags().StringP("repo", "R", "", "Filter codespace selection by repository name (user/repo)")
	codespace_stopCmd.PersistentFlags().String("repo-owner", "", "Filter codespace selection by repository owner (username or org)")
	codespace_stopCmd.Flags().StringP("user", "u", "", "The `username` to stop codespace for (used with --org)")
	codespaceCmd.AddCommand(codespace_stopCmd)

	carapace.Gen(codespace_stopCmd).FlagCompletion(carapace.ActionMap{
		"codespace": action.ActionCodespaces(),
		"org":       gh.ActionOrganizations(gh.HostOpts{}),
		"repo":      gh.ActionOwnerRepositories(gh.HostOpts{}),
		"user":      gh.ActionUsers(gh.HostOpts{}),
	})
}
