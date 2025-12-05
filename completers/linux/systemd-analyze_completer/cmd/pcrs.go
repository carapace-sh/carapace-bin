package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pcrsCmd = &cobra.Command{
	Use:   "pcrs",
	Short: "Show TPM2 PCRs and their names",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pcrsCmd).Standalone()

	rootCmd.AddCommand(pcrsCmd)
}
