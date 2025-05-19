package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var codespace_selectCmd = &cobra.Command{
	Use:    "select",
	Short:  "Select a Codespace",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codespace_selectCmd).Standalone()

	codespace_selectCmd.PersistentFlags().StringP("codespace", "c", "", "Name of the codespace")
	codespace_selectCmd.Flags().StringP("file", "f", "", "Output file path")
	codespace_selectCmd.PersistentFlags().StringP("repo", "R", "", "Filter codespace selection by repository name (user/repo)")
	codespace_selectCmd.PersistentFlags().String("repo-owner", "", "Filter codespace selection by repository owner (username or org)")
	codespaceCmd.AddCommand(codespace_selectCmd)

	carapace.Gen(codespace_selectCmd).FlagCompletion(carapace.ActionMap{
		"codespace":  action.ActionCodespaces(),
		"file":       carapace.ActionFiles(),
		"repo":       gh.ActionOwnerRepositories(gh.HostOpts{}),
		"repo-owner": gh.ActionOwners(gh.HostOpts{}),
	})
}
