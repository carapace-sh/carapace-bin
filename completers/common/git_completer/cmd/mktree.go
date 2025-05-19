package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mktreeCmd = &cobra.Command{
	Use:     "mktree",
	Short:   "Build a tree-object from ls-tree formatted text",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_manipulator].ID,
}

func init() {
	carapace.Gen(mktreeCmd).Standalone()

	mktreeCmd.Flags().Bool("batch", false, "allow creation of more than one tree")
	mktreeCmd.Flags().Bool("missing", false, "allow missing objects")
	mktreeCmd.Flags().Bool("no-batch", false, "do not allow creation of more than one tree")
	mktreeCmd.Flags().Bool("no-missing", false, "do not allow missing objects")
	mktreeCmd.Flags().BoolS("z", "z", false, "input is NUL terminated")
	rootCmd.AddCommand(mktreeCmd)
}
