package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Show changes as a diff. No argument: uncommitted changes; with a commit: the change that commit introduced",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diffCmd).Standalone()

	diffCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(diffCmd)

	carapace.Gen(diffCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)

	carapace.Gen(diffCmd).DashAnyCompletion(
		carapace.ActionPositional(diffCmd),
	)
}
