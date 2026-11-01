package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var attestationCmd = &cobra.Command{
	Use:   "attestation <command> [flags]",
	Short: "Manage software attestations. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(attestationCmd).Standalone()

	rootCmd.AddCommand(attestationCmd)
}
