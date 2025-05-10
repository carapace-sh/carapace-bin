package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var codespace_viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View details about a codespace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codespace_viewCmd).Standalone()

	codespace_viewCmd.PersistentFlags().StringP("codespace", "c", "", "Name of the codespace")
	codespace_viewCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	codespace_viewCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	codespace_viewCmd.PersistentFlags().StringP("repo", "R", "", "Filter codespace selection by repository name (user/repo)")
	codespace_viewCmd.PersistentFlags().String("repo-owner", "", "Filter codespace selection by repository owner (username or org)")
	codespace_viewCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	codespaceCmd.AddCommand(codespace_viewCmd)

	carapace.Gen(codespace_viewCmd).FlagCompletion(carapace.ActionMap{
		"codespace":  action.ActionCodespaces(),
		"json":       gh.ActionCodespaceViewFields().UniqueList(","),
		"repo":       gh.ActionOwnerRepositories(gh.HostOpts{}),
		"repo-owner": gh.ActionOwners(gh.HostOpts{}),
	})
}
