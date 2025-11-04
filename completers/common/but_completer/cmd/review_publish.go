package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var review_publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish review requests for active branches in your workspace. By default, publishes reviews for all active branches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(review_publishCmd).Standalone()

	review_publishCmd.Flags().StringP("branch", "b", "", "Publish reviews only for the specified branch")
	review_publishCmd.Flags().BoolP("default", "t", false, "Whether to use just the branch name as the review title, without opening an editor")
	review_publishCmd.Flags().BoolP("help", "h", false, "Print help")
	review_publishCmd.Flags().BoolP("run-hooks", "r", false, "Run pre-push hooks (defaults to true)")
	review_publishCmd.Flags().BoolP("skip-force-push-protection", "s", false, "Skip force push protection checks")
	review_publishCmd.Flags().BoolP("with-force", "f", false, "Force push even if it's not fast-forward (defaults to true)")
	reviewCmd.AddCommand(review_publishCmd)

	carapace.Gen(review_publishCmd).FlagCompletion(carapace.ActionMap{
		"branch": but.ActionLocalBranches(),
	})
}
