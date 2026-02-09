package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:     "push",
	Short:   "Push changes in a branch to remote.",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "server interactions",
}

func init() {
	carapace.Gen(pushCmd).Standalone()

	pushCmd.Flags().BoolP("dry-run", "d", false, "Show what would be pushed without actually pushing")
	pushCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	pushCmd.Flags().BoolP("run-hooks", "r", false, "Run pre-push hooks")
	pushCmd.Flags().BoolP("skip-force-push-protection", "s", false, "Skip force push protection checks")
	pushCmd.Flags().BoolP("with-force", "f", false, "Force push even if it's not fast-forward")
	rootCmd.AddCommand(pushCmd)

	carapace.Gen(pushCmd).PositionalCompletion(
		but.ActionLocalBranches(),
	)
}
