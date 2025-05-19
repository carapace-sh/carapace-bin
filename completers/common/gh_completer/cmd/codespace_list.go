package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var codespace_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List codespaces",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codespace_listCmd).Standalone()

	codespace_listCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	codespace_listCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	codespace_listCmd.Flags().StringP("limit", "L", "", "Maximum number of codespaces to list")
	codespace_listCmd.Flags().StringP("org", "o", "", "The `login` handle of the organization to list codespaces for (admin-only)")
	codespace_listCmd.Flags().StringP("repo", "R", "", "Repository name with owner: user/repo")
	codespace_listCmd.Flags().StringP("repo-deprecated", "r", "", "(Deprecated) Shorthand for --repo")
	codespace_listCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	codespace_listCmd.Flags().StringP("user", "u", "", "The `username` to list codespaces for (used with --org)")
	codespace_listCmd.Flags().BoolP("web", "w", false, "List codespaces in the web browser, cannot be used with --user or --org")
	codespace_listCmd.Flag("repo-deprecated").Hidden = true
	codespaceCmd.AddCommand(codespace_listCmd)

	carapace.Gen(codespace_listCmd).FlagCompletion(carapace.ActionMap{
		"json": action.ActionCodespaceFields().UniqueList(","),
		"org":  gh.ActionOrganizations(gh.HostOpts{}),
		"repo": gh.ActionOwnerRepositories(gh.HostOpts{}),
		"user": gh.ActionUsers(gh.HostOpts{}),
	})
}
