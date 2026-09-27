package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Test connectivity to the configured registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pingCmd).Standalone()

	pingCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	pingCmd.Flags().String("registry", "", "Test a specific registry URL")
	rootCmd.AddCommand(pingCmd)
}
