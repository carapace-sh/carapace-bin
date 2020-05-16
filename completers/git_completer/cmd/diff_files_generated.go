package cmd

import (
	"github.com/spf13/cobra"
)

var diff_filesCmd = &cobra.Command{
	Use:   "diff-files",
	Short: "Compares files in the working tree and the index",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {

	rootCmd.AddCommand(diff_filesCmd)
}
