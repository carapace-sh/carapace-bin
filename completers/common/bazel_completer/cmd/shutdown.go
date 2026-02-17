package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shutdownCmd = &cobra.Command{
	Use:   "shutdown",
	Short: "Stops the %{product} server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shutdownCmd).Standalone()

	shutdownCmd.Flags().String("iff_heap_size_greater_than", "", "Iff non-zero, then shutdown will only shut down the server if the total memory (in MB) consumed by the JVM exceeds this value.")
	rootCmd.AddCommand(shutdownCmd)
}
