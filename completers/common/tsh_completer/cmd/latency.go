package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var latencyCmd = &cobra.Command{
	Use:    "latency",
	Short:  "Run latency diagnostics.",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(latencyCmd).Standalone()

	rootCmd.AddCommand(latencyCmd)
}
