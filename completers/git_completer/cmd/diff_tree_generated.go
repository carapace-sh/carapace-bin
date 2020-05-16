package cmd

import (
	"github.com/spf13/cobra"
)

var diff_treeCmd = &cobra.Command{
	Use:   "diff-tree",
	Short: "Compares the content and mode of blobs found via two tree objects",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {

	rootCmd.AddCommand(diff_treeCmd)
}
