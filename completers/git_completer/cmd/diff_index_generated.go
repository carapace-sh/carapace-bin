package cmd

import (
	"github.com/spf13/cobra"
)

var diff_indexCmd = &cobra.Command{
	Use:   "diff-index",
	Short: "Compare a tree to the working tree or index",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {

	rootCmd.AddCommand(diff_indexCmd)
}
