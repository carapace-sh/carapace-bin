package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Stage changes. With no paths, stages everything",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addCmd).Standalone()

	addCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(addCmd)

	carapace.Gen(addCmd).PositionalAnyCompletion(
		git.ActionChanges(git.ChangeOpts{Unstaged: true}).FilterArgs(),
	)

	carapace.Gen(addCmd).DashAnyCompletion(
		carapace.ActionPositional(addCmd),
	)
}
