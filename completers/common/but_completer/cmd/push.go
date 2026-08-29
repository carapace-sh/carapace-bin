package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:     "push",
	Short:   "Push changes in a branch to remote",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "server interactions",
}

func init() {
	carapace.Gen(pushCmd).Standalone()

	pushCmd.Flags().Bool("allow-merged", false, "Allow targeting branches and commits that are already merged upstream")
	pushCmd.Flags().BoolP("dry-run", "d", false, "Show what would be pushed without actually pushing")
	pushCmd.Flags().StringSliceP("hashtag", "a", nil, "Add hashtag(s) to change (Gerrit). Can be used multiple times")
	pushCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	pushCmd.Flags().Bool("no-hooks", false, "Bypass pre-push hooks")
	pushCmd.Flags().Bool("no-verify", false, "Bypass pre-push hooks")
	pushCmd.Flags().BoolP("private", "p", false, "Mark change as private (Gerrit)")
	pushCmd.Flags().BoolP("ready", "y", false, "Mark change as ready for review (Gerrit). This is the default state")
	pushCmd.Flags().BoolP("skip-force-push-protection", "s", false, "Skip force push protection checks")
	pushCmd.Flags().StringSlice("tag", nil, "Add hashtag(s) to change (Gerrit). Can be used multiple times")
	pushCmd.Flags().Bool("tb", false, "Use branch name as topic (Gerrit)")
	pushCmd.Flags().StringP("topic", "t", "", "Add custom topic to change (Gerrit). At most one topic can be set")
	pushCmd.Flags().Bool("topic-from-branch", false, "Use branch name as topic (Gerrit)")
	pushCmd.Flags().BoolP("wip", "w", false, "Mark change as work-in-progress (Gerrit). Mutually exclusive with --ready")
	pushCmd.Flags().BoolP("with-force", "f", false, "Force push even if it's not fast-forward")
	pushCmd.Flag("hashtag").Hidden = true
	pushCmd.Flag("private").Hidden = true
	pushCmd.Flag("ready").Hidden = true
	pushCmd.Flag("tag").Hidden = true
	pushCmd.Flag("tb").Hidden = true
	pushCmd.Flag("topic").Hidden = true
	pushCmd.Flag("topic-from-branch").Hidden = true
	pushCmd.Flag("wip").Hidden = true
	rootCmd.AddCommand(pushCmd)

	carapace.Gen(pushCmd).PositionalCompletion(
		but.ActionLocalBranches(),
	)
}
