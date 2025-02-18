package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "insert RULE at NUM",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(insertCmd)

	rootCmd.AddCommand(insertCmd)
}
