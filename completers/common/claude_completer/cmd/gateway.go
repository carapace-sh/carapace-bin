package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "Run the enterprise auth/telemetry gateway",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gatewayCmd).Standalone()

	gatewayCmd.Flags().String("config", "", "Path to gateway YAML config")
	gatewayCmd.Flags().BoolP("help", "h", false, "Display help for command")
	rootCmd.AddCommand(gatewayCmd)
}
