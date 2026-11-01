package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var metricsCmd = &cobra.Command{
	Use:   "metrics",
	Short: "Show Tailscale metrics",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(metricsCmd).Standalone()

	rootCmd.AddCommand(metricsCmd)
}
