package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var pr_newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new review for a branch. If no branch is specified, you will be prompted to select one. If there is only one branch without a review, you will be asked to confirm",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_newCmd).Standalone()

	pr_newCmd.Flags().BoolP("default", "t", false, "Use the default content for the review title and description, skipping any prompts. If the branch contains only a single commit, the commit message will be used")
	pr_newCmd.Flags().BoolP("draft", "d", false, "Whether to create reviews as a draft")
	pr_newCmd.Flags().StringP("file", "F", "", "Read review title and description from file. The first line is the title, the rest is the description")
	pr_newCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	pr_newCmd.Flags().StringP("message", "m", "", "review title and description. The first line is the title, the rest is the description")
	pr_newCmd.Flags().BoolP("run-hooks", "r", false, "Run pre-push hooks (defaults to true)")
	pr_newCmd.Flags().BoolP("skip-force-push-protection", "s", false, "Skip force push protection checks")
	pr_newCmd.Flags().BoolP("with-force", "f", false, "Force push even if it's not fast-forward (defaults to true)")
	prCmd.AddCommand(pr_newCmd)

	carapace.Gen(pr_newCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})
	carapace.Gen(pr_newCmd).PositionalCompletion(
		but.ActionLocalBranches(),
	)
}
