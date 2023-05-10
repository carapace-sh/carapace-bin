package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var codespace_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete codespaces",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codespace_deleteCmd).Standalone()

	codespace_deleteCmd.Flags().Bool("all", false, "Delete all codespaces")
	codespace_deleteCmd.PersistentFlags().StringP("codespace", "c", "", "Name of the codespace")
	codespace_deleteCmd.Flags().Uint16("days", 0, "Delete codespaces older than `N` days")
	codespace_deleteCmd.Flags().BoolP("force", "f", false, "Skip confirmation for codespaces that contain unsaved changes")
	codespace_deleteCmd.Flags().StringP("org", "o", "", "The `login` handle of the organization (admin-only)")
	codespace_deleteCmd.PersistentFlags().StringP("repo", "R", "", "Filter codespace selection by repository name (user/repo)")
	codespace_deleteCmd.PersistentFlags().String("repo-owner", "", "Filter codespace selection by repository owner (username or org)")
	codespace_deleteCmd.Flags().StringP("user", "u", "", "The `username` to delete codespaces for (used with --org)")
	codespaceCmd.AddCommand(codespace_deleteCmd)

	carapace.Gen(codespace_deleteCmd).FlagCompletion(carapace.ActionMap{
		"codespace":  action.ActionCodespaces(),
		"org":        gh.ActionOrganizations(gh.HostOpts{}),
		"repo":       action.ActionOwnerRepositories(codespace_deleteCmd),
		"repo-owner": gh.ActionOwners(gh.HostOpts{}),
		"user":       gh.ActionUsers(gh.HostOpts{}),
	})
}
