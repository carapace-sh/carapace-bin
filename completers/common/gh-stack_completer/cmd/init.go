package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [branches...]",
	Short: "Initialize a new stack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().BoolP("adopt", "a", false, "Track existing branches as part of a stack")
	initCmd.Flags().StringP("base", "b", "", "Trunk branch for stack (defaults to default branch)")
	initCmd.Flags().BoolP("numbered", "n", false, "Use auto-incrementing numbered branch names (requires --prefix)")
	initCmd.Flags().StringP("prefix", "p", "", "Branch name prefix for the stack")
	rootCmd.AddCommand(initCmd)

	carapace.Gen(initCmd).FlagCompletion(carapace.ActionMap{
		"base": git.ActionLocalBranches(), // TODO refs?
	})

	carapace.Gen(initCmd).PositionalAnyCompletion(
		git.ActionLocalBranches().FilterArgs(),
	)
}
