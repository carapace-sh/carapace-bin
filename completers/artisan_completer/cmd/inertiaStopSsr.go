package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inertiaStopSsrCmd = &cobra.Command{
	Use:   "inertia:stop-ssr",
	Short: "Stop the Inertia SSR server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inertiaStopSsrCmd).Standalone()

	rootCmd.AddCommand(inertiaStopSsrCmd)
}
