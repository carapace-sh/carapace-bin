package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var novaCustomFilterCmd = &cobra.Command{
	Use:   "nova:custom-filter <name>",
	Short: "Create a new custom filter",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(novaCustomFilterCmd).Standalone()

	rootCmd.AddCommand(novaCustomFilterCmd)
}
