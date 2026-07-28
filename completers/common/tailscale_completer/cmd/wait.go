package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waitCmd = &cobra.Command{
	Use:   "wait",
	Short: "Wait for Tailscale interface/IPs to be ready for binding",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waitCmd).Standalone()

	waitCmd.Flags().String("timeout", "", "how long to wait before giving up (0 means wait indefinitely)")
	rootCmd.AddCommand(waitCmd)
}
