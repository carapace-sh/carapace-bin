package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var limitCmd = &cobra.Command{
	Use:   "limit",
	Short: "add limit rule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(limitCmd).Standalone()

	rootCmd.AddCommand(limitCmd)
}
