package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hasTpm2Cmd = &cobra.Command{
	Use:   "has-tpm2",
	Short: "Report whether TPM2 support is available",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hasTpm2Cmd).Standalone()

	rootCmd.AddCommand(hasTpm2Cmd)
}
