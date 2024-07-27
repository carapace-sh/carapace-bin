package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var novaValueCmd = &cobra.Command{
	Use:   "nova:value <name>",
	Short: "Create a new metric (single value) class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(novaValueCmd).Standalone()

	rootCmd.AddCommand(novaValueCmd)
}
