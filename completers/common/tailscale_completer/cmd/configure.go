package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure the host to enable more Tailscale features",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configureCmd).Standalone()

	rootCmd.AddCommand(configureCmd)
}
