package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/carapace-sh/carapace/pkg/condition"
	"github.com/spf13/cobra"
)

var codespace_cpCmd = &cobra.Command{
	Use:   "cp [-e] [-r] [-- [<scp flags>...]] <sources>... <dest>",
	Short: "Copy files between local and remote file systems",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codespace_cpCmd).Standalone()

	codespace_cpCmd.PersistentFlags().StringP("codespace", "c", "", "Name of the codespace")
	codespace_cpCmd.Flags().BoolP("expand", "e", false, "Expand remote file names on remote shell")
	codespace_cpCmd.Flags().StringP("profile", "p", "", "Name of the SSH profile to use")
	codespace_cpCmd.Flags().BoolP("recursive", "r", false, "Recursively copy directories")
	codespace_cpCmd.PersistentFlags().StringP("repo", "R", "", "Filter codespace selection by repository name (user/repo)")
	codespace_cpCmd.PersistentFlags().String("repo-owner", "", "Filter codespace selection by repository owner (username or org)")
	codespaceCmd.AddCommand(codespace_cpCmd)

	// TODO profile completion
	carapace.Gen(codespace_cpCmd).FlagCompletion(carapace.ActionMap{
		"codespace":  action.ActionCodespaces(),
		"repo":       gh.ActionOwnerRepositories(gh.HostOpts{}),
		"repo-owner": gh.ActionOwners(gh.HostOpts{}),
	})

	carapace.Gen(codespace_cpCmd).PositionalAnyCompletion(
		carapace.Batch(
			carapace.ActionFiles(),
			action.ActionCodespacePath(
				codespace_cpCmd.Flag("codespace").Value.String(),
				codespace_cpCmd.Flag("expand").Changed,
			).Prefix("remote:").UnlessF(condition.CompletingPath),
		).ToA(),
	)
}
