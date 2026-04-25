package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var unstackCmd = &cobra.Command{
	Use:     "unstack [branch]",
	Short:   "Delete a stack locally and on GitHub",
	Aliases: []string{"delete"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unstackCmd).Standalone()

	unstackCmd.Flags().Bool("local", false, "Only delete the stack locally")
	rootCmd.AddCommand(unstackCmd)

	carapace.Gen(unstackCmd).PositionalCompletion(
		git.ActionLocalBranches(),
	)
}
