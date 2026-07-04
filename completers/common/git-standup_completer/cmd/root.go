package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-standup",
	Short: "Recall the commit history",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-standup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("B", "B", false, "Display the commits in branch groups")
	rootCmd.Flags().StringS("D", "D", "", "The date format displayed in commit history")
	rootCmd.Flags().StringS("F", "F", "", "Show if commit is GPG signed (G) or not (N), or print author date (authordate)")
	rootCmd.Flags().BoolS("L", "L", false, "Enable the inclusion of symbolic links in recursive directory search")
	rootCmd.Flags().StringS("a", "a", "", "The author of commits")
	rootCmd.Flags().StringS("d", "d", "", "The start of commit history")
	rootCmd.Flags().BoolS("f", "f", false, "Fetch the latest commits before showing commit history")
	rootCmd.Flags().BoolS("h", "h", false, "Display help message")
	rootCmd.Flags().StringS("m", "m", "", "The depth of recursive directory search")
	rootCmd.Flags().StringS("n", "n", "", "Limit the number of commits displayed per group")
	rootCmd.Flags().StringS("w", "w", "", "Specify weekday range to limit search to")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"F": carapace.ActionValues("gpg", "authordate"),
		"a": carapace.Batch(
			carapace.ActionValuesDescribed("all", "all authors").Style(style.Blue),
			git.ActionAuthors(),
		).ToA(),
	})
}
