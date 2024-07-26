package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var novaPartitionCmd = &cobra.Command{
	Use:   "nova:partition <name>",
	Short: "Create a new metric (partition) class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(novaPartitionCmd).Standalone()

	rootCmd.AddCommand(novaPartitionCmd)
}
