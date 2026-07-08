package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show information about a commit, tag, or branch (defaults to HEAD)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showCmd).Standalone()

	showCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(showCmd)

	carapace.Gen(showCmd).PositionalAnyCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
	)

	carapace.Gen(showCmd).DashAnyCompletion(
		carapace.ActionPositional(showCmd),
	)
}
