package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configure_synologyCertCmd = &cobra.Command{
	Use:   "synology-cert",
	Short: "Configure Synology with a TLS certificate for your tailnet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configure_synologyCertCmd).Standalone()

	configure_synologyCertCmd.Flags().String("domain", "", "tailnet domain for the certificate")
	configureCmd.AddCommand(configure_synologyCertCmd)
}
