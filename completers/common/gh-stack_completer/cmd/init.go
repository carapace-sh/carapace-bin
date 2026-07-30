package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:     "init [branches...]",
	Short:   "Initialize a new stack",
	GroupID: "stack",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().BoolP("adopt", "a", false, "Deprecated: existing branches are now adopted automatically")
	initCmd.Flags().StringP("base", "b", "", "Trunk branch for stack (defaults to default branch)")
	initCmd.Flag("adopt").Hidden = true
	rootCmd.AddCommand(initCmd)

	carapace.Gen(initCmd).FlagCompletion(carapace.ActionMap{
		"base": git.ActionLocalBranches(), // TODO refs?
	})

	carapace.Gen(initCmd).PositionalAnyCompletion(
		git.ActionLocalBranches().FilterArgs(),
	)
}
