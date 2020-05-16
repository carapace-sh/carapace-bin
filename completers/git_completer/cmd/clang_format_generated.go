package cmd

import (
	"github.com/spf13/cobra"
)

var clang_formatCmd = &cobra.Command{
	Use:   "clang-format",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {

	rootCmd.AddCommand(clang_formatCmd)
}
