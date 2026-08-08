package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stack_amendCmd = &cobra.Command{
	Use:   "amend",
	Short: "Save more changes to a stacked diff. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_amendCmd).Standalone()

	stack_amendCmd.Flags().BoolP("all", "a", false, "Automatically stage modified and deleted tracked files.")
	stack_amendCmd.Flags().StringP("description", "d", "", "A description of the change.")
	stack_amendCmd.Flags().StringP("message", "m", "", "Alias for the description flag.")
	stack_amendCmd.Flags().Bool("no-verify", false, "Bypass the pre-commit and commit-msg hooks of git-commit(1).")
	stack_amendCmd.Flags().Bool("reword", false, "Only update the commit message without staging any files.")
	stackCmd.AddCommand(stack_amendCmd)
}
