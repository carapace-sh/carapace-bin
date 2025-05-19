package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var codespace_logsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Access codespace logs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codespace_logsCmd).Standalone()

	codespace_logsCmd.PersistentFlags().StringP("codespace", "c", "", "Name of the codespace")
	codespace_logsCmd.Flags().BoolP("follow", "f", false, "Tail and follow the logs")
	codespace_logsCmd.PersistentFlags().StringP("repo", "R", "", "Filter codespace selection by repository name (user/repo)")
	codespace_logsCmd.PersistentFlags().String("repo-owner", "", "Filter codespace selection by repository owner (username or org)")
	codespaceCmd.AddCommand(codespace_logsCmd)

	carapace.Gen(codespace_logsCmd).FlagCompletion(carapace.ActionMap{
		"codespace":  action.ActionCodespaces(),
		"repo":       gh.ActionOwnerRepositories(gh.HostOpts{}),
		"repo-owner": gh.ActionOwners(gh.HostOpts{}),
	})
}
