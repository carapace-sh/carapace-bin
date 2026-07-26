package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-changelog",
	Short: "Generate a changelog from git tags and commit messages",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-changelog",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all", "a", false, "Retrieve all commits")
	rootCmd.Flags().String("final-tag", "", "Newest tag in range")
	rootCmd.Flags().BoolP("help", "h", false, "show help")
	rootCmd.Flags().BoolP("list", "l", false, "Display as list without titles")
	rootCmd.Flags().BoolP("merges-only", "m", false, "Use only merge commits")
	rootCmd.Flags().BoolP("no-merges", "n", false, "Suppress merge commits")
	rootCmd.Flags().BoolP("prune-old", "p", false, "Replace existing changelog entirely")
	rootCmd.Flags().String("start-commit", "", "Oldest commit in range")
	rootCmd.Flags().StringP("start-tag", "s", "", "Oldest tag in range")
	rootCmd.Flags().BoolP("stdout", "x", false, "Write to stdout")
	rootCmd.Flags().StringP("tag", "t", "", "Tag label for most-recent untagged commits")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"final-tag":    git.ActionTags(),
		"start-commit": git.ActionRefs(git.RefOption{}.Default()),
		"start-tag":    git.ActionTags(),
		"tag":          git.ActionTags(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
