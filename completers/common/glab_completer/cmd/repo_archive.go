package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var repo_archiveCmd = &cobra.Command{
	Use:   "archive <command> [flags]",
	Short: "Get an archive of the repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(repo_archiveCmd).Standalone()

	repo_archiveCmd.Flags().StringP("format", "f", "", "Optional. Specify format if you want a downloaded archive: tar.gz, tar.bz2, tbz, tbz2, tb2, bz2, tar, zip.")
	repo_archiveCmd.Flags().StringP("sha", "s", "", "The commit SHA to download. A tag, branch reference, or SHA can be used. Defaults to the tip of the default branch if not specified.")
	repoCmd.AddCommand(repo_archiveCmd)

	carapace.Gen(repo_archiveCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("tar.gz", "tar.bz2", "tbz", "tbz2", "tb2", "bz2", "tar", "zip"),
		"sha": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.Batch(
				action.ActionBranches(repo_archiveCmd),
				action.ActionTags(repo_archiveCmd),
			).Invoke(c).Merge().ToA() // TODO sha
		}),
	})

	carapace.Gen(repo_archiveCmd).PositionalCompletion(
		action.ActionRepo(repo_archiveCmd),
		carapace.ActionDirectories(),
	)
}
