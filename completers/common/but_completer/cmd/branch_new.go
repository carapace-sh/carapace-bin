package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_newCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new branch in the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_newCmd).Standalone()

	branch_newCmd.Flags().StringP("anchor", "a", "", "Anchor point - either a commit ID or branch name to create the new branch from")
	branch_newCmd.Flags().BoolP("help", "h", false, "Print help")
	branchCmd.AddCommand(branch_newCmd)

	carapace.Gen(branch_newCmd).FlagCompletion(carapace.ActionMap{
		"anchor": carapace.ActionValues(), // TODO git.ActionRefs(git.RefOption{}.Default())?
	})
}
