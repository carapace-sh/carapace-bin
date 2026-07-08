package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-summary",
	Short: "Show repository summary",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-summary",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("dedup-by-email", false, "Merge duplicate authors by email")
	rootCmd.Flags().Bool("full-path", false, "Show full repository path")
	rootCmd.Flags().Bool("help", false, "show help")
	rootCmd.Flags().Bool("line", false, "Summarize by lines instead of commits")
	rootCmd.Flags().Bool("no-merges", false, "Exclude merge commits")
	rootCmd.Flags().String("output-style", "", "Output style: tabular or oneline")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"output-style": carapace.ActionValues("tabular", "oneline"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
