package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var writeTreeCmd = &cobra.Command{
	Use:     "write-tree",
	Short:   "Create a tree object from the current index",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_manipulator].ID,
}

func init() {
	carapace.Gen(writeTreeCmd).Standalone()

	writeTreeCmd.Flags().Bool("missing-ok", false, "allow missing objects")
	writeTreeCmd.Flags().String("prefix", "", "write tree object for a subdirectory <prefix>")
	rootCmd.AddCommand(writeTreeCmd)
}
