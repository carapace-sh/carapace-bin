package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:     "commit",
	Short:   "Commit changes to a stack.",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "branching and committing",
}

func init() {
	carapace.Gen(commitCmd).Standalone()

	commitCmd.Flags().StringP("ai", "i", "", "Generate commit message using AI with optional user summary")
	commitCmd.Flags().BoolP("create", "c", false, "Whether to create a new branch for this commit. If the branch name given matches an existing branch, that branch will be used instead. If no branch name is given, a new branch with a generated name will be created")
	commitCmd.Flags().StringP("file", "f", "", "Read commit message from file")
	commitCmd.Flags().StringSliceP("files", "F", nil, "Uncommitted file or hunk CLI IDs to include in the commit. Can be specified multiple times or as comma-separated values. If not specified, all uncommitted changes (or changes staged to the target branch) are committed")
	commitCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	commitCmd.Flags().StringP("message", "m", "", "Commit message")
	commitCmd.Flags().BoolP("no-hooks", "n", false, "Bypass pre-commit hooks")
	commitCmd.Flags().Bool("no-verify", false, "Bypass pre-commit hooks")
	commitCmd.Flags().BoolP("only", "o", false, "Only commit staged files, not unstaged files")
	commitCmd.Flag("no-verify").Hidden = true
	rootCmd.AddCommand(commitCmd)

	carapace.Gen(commitCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
		"files": carapace.Batch(
			git.ActionChanges(git.ChangeOpts{}.Default()),
			but.ActionCliIds(but.CliIdsOpts{Changes: true}),
		).ToA().UniqueList(","), // TODO unique only filters one of change/cliId
	})

	carapace.Gen(commitCmd).PositionalCompletion(
		but.ActionLocalBranches(),
	)
}
