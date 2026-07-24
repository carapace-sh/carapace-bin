package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_metricsCmd = &cobra.Command{
	Use:   "metrics",
	Short: "Print tailscaled's metrics",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_metricsCmd).Standalone()

	debug_metricsCmd.Flags().Bool("watch", false, "watch metrics continuously")
	debugCmd.AddCommand(debug_metricsCmd)
}
