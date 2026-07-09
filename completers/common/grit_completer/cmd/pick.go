package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var pickCmd = &cobra.Command{
	Use:   "pick",
	Short: "Cherry-pick a commit onto the current branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pickCmd).Standalone()

	pickCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(pickCmd)

	carapace.Gen(pickCmd).PositionalCompletion(
		git.ActionRefRanges(git.RefOption{}.Default()),
	)

	carapace.Gen(pickCmd).DashAnyCompletion(
		carapace.ActionPositional(pickCmd),
	)
}
