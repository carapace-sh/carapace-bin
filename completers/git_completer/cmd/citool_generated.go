package cmd

import (
	"github.com/spf13/cobra"
)

var citoolCmd = &cobra.Command{
	Use:   "citool",
	Short: "Graphical alternative to git-commit",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {

	rootCmd.AddCommand(citoolCmd)
}
