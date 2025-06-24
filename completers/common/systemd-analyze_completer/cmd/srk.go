package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var srkCmd = &cobra.Command{
	Use:   "srk",
	Short: "Write TPM2 SRK (to FILE)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(srkCmd).Standalone()

	rootCmd.AddCommand(srkCmd)
}
