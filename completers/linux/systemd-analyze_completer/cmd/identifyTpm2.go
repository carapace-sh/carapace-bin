package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var identifyTpm2Cmd = &cobra.Command{
	Use:   "identify-tpm2",
	Short: "Show TPM2 vendor information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(identifyTpm2Cmd).Standalone()

	rootCmd.AddCommand(identifyTpm2Cmd)
}
