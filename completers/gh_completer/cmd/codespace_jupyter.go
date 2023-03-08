package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var codespace_jupyterCmd = &cobra.Command{
	Use:   "jupyter",
	Short: "Open a codespace in JupyterLab",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codespace_jupyterCmd).Standalone()

	codespace_jupyterCmd.PersistentFlags().StringP("codespace", "c", "", "Name of the codespace")
	codespace_jupyterCmd.PersistentFlags().StringP("repo", "R", "", "Filter codespace selection by repository name (user/repo)")
	codespaceCmd.AddCommand(codespace_jupyterCmd)

	carapace.Gen(codespace_jupyterCmd).FlagCompletion(carapace.ActionMap{
		"codespace": action.ActionCodespaces(),
		"repo":      gh.ActionOwnerRepositories(gh.HostOpts{}),
	})
}
