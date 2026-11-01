package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var pr_newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new review for a branch, force-pushing it first",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_newCmd).Standalone()

	pr_newCmd.Flags().BoolP("default", "t", false, "Use the default content for the review title and description, skipping any prompts. If the branch contains only a single commit, the commit message will be used")
	pr_newCmd.Flags().BoolP("draft", "d", false, "Create the review as a draft")
	pr_newCmd.Flags().StringP("file", "F", "", "Read review title and description from FILE, or from stdin when FILE is -. The first line is the title, the rest is the description")
	pr_newCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	pr_newCmd.Flags().StringP("message", "m", "", "Review title and description: the first line is the title, the rest is the description. A non-interactive run needs -m, -F, or -t")
	pr_newCmd.Flags().Bool("no-hooks", false, "Bypass pre-push hooks")
	pr_newCmd.Flags().Bool("no-verify", false, "Bypass pre-push hooks")
	pr_newCmd.Flags().BoolP("skip-force-push-protection", "s", false, "Skip force push protection checks")
	pr_newCmd.Flags().BoolP("with-force", "f", false, "Force push even if it's not fast-forward. Always on; the flag is kept for compatibility")
	pr_newCmd.Flag("with-force").Hidden = true
	prCmd.AddCommand(pr_newCmd)

	carapace.Gen(pr_newCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})
	carapace.Gen(pr_newCmd).PositionalCompletion(
		but.ActionLocalBranches(),
	)
}
