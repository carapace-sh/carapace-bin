package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var codespace_codeCmd = &cobra.Command{
	Use:   "code",
	Short: "Open a codespace in Visual Studio Code",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codespace_codeCmd).Standalone()

	codespace_codeCmd.PersistentFlags().StringP("codespace", "c", "", "Name of the codespace")
	codespace_codeCmd.Flags().Bool("insiders", false, "Use the insiders version of Visual Studio Code")
	codespace_codeCmd.PersistentFlags().StringP("repo", "R", "", "Filter codespace selection by repository name (user/repo)")
	codespace_codeCmd.PersistentFlags().String("repo-owner", "", "Filter codespace selection by repository owner (username or org)")
	codespace_codeCmd.Flags().BoolP("web", "w", false, "Use the web version of Visual Studio Code")
	codespaceCmd.AddCommand(codespace_codeCmd)

	carapace.Gen(codespace_codeCmd).FlagCompletion(carapace.ActionMap{
		"codespace":  action.ActionCodespaces(),
		"repo":       gh.ActionOwnerRepositories(gh.HostOpts{}),
		"repo-owner": gh.ActionOwners(gh.HostOpts{}),
	})
}
