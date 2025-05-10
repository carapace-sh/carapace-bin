package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var codespace_portsCmd = &cobra.Command{
	Use:   "ports",
	Short: "List ports in a codespace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codespace_portsCmd).Standalone()

	codespace_portsCmd.PersistentFlags().StringP("codespace", "c", "", "Name of the codespace")
	codespace_portsCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	codespace_portsCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	codespace_portsCmd.PersistentFlags().StringP("repo", "R", "", "Filter codespace selection by repository name (user/repo)")
	codespace_portsCmd.PersistentFlags().String("repo-owner", "", "Filter codespace selection by repository owner (username or org)")
	codespace_portsCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	codespaceCmd.AddCommand(codespace_portsCmd)

	carapace.Gen(codespace_portsCmd).FlagCompletion(carapace.ActionMap{
		"codespace":  action.ActionCodespaces(),
		"json":       action.ActionCodespacePortFields().UniqueList(","),
		"repo":       gh.ActionOwnerRepositories(gh.HostOpts{}),
		"repo-owner": gh.ActionOwners(gh.HostOpts{}),
	})
}
