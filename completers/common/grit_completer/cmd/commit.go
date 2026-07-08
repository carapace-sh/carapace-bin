package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Stage every change and record a new commit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(commitCmd).Standalone()

	commitCmd.Flags().BoolP("all", "a", false, "Stage every change first, then commit. This is the default behavior")
	commitCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	commitCmd.Flags().StringP("message", "m", "", "Commit message")
	rootCmd.AddCommand(commitCmd)

	carapace.Gen(commitCmd).PositionalAnyCompletion(
		git.ActionChanges(git.ChangeOpts{Staged: true, Unstaged: true}).FilterArgs(),
	)

	carapace.Gen(commitCmd).DashAnyCompletion(
		carapace.ActionPositional(commitCmd),
	)
}
