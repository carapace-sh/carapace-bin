package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "List branches, or create / delete one",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branchCmd).Standalone()

	branchCmd.Flags().BoolP("delete", "d", false, "Delete the named branch instead of creating it")
	branchCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(branchCmd)

	carapace.Gen(branchCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{LocalBranches: true}),
	)

	carapace.Gen(branchCmd).DashAnyCompletion(
		carapace.ActionPositional(branchCmd),
	)
}
