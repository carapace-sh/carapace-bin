package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var nvpcrsCmd = &cobra.Command{
	Use:   "nvpcrs",
	Short: "Show additional TPM2 PCRs stored in NV indexes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(nvpcrsCmd).Standalone()

	rootCmd.AddCommand(nvpcrsCmd)
}
