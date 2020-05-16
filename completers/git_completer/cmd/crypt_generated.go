package cmd

import (
	"github.com/spf13/cobra"
)

var cryptCmd = &cobra.Command{
	Use:   "crypt",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {

	rootCmd.AddCommand(cryptCmd)
}
