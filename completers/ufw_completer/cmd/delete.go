package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete RULE",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deleteCmd)

	rootCmd.AddCommand(deleteCmd)
}
