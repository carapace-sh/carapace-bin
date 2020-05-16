package cmd

import (
	"github.com/spf13/cobra"
)

var varCmd = &cobra.Command{
	Use:   "var",
	Short: "Show a Git logical variable",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {

	rootCmd.AddCommand(varCmd)
}
