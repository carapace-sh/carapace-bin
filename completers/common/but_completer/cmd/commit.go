package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit [BRANCH]",
	Short: "Commit changes to a stack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(commitCmd).Standalone()

	commitCmd.Flags().BoolP("create", "c", false, "Whether to create a new branch for this commit. If the branch name given matches an existing branch, that branch will be used instead. If no branch name is given, a new branch with a generated name will be created")
	commitCmd.Flags().BoolP("help", "h", false, "Print help")
	commitCmd.Flags().StringP("message", "m", "", "Commit message")
	commitCmd.Flags().BoolP("only", "o", false, "Only commit assigned files, not unassigned files")
	rootCmd.AddCommand(commitCmd)

	carapace.Gen(commitCmd).PositionalCompletion(
		but.ActionLocalBranches(),
	)
}
