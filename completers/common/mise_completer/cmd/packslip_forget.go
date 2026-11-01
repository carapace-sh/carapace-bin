package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var packslip_forgetCmd = &cobra.Command{
	Use:   "forget",
	Short: "Forget a project's pinned signer so the next accepted release sets it",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(packslip_forgetCmd).Standalone()

	packslip_forgetCmd.Flags().BoolP("help", "h", false, "Print help")
	packslipCmd.AddCommand(packslip_forgetCmd)
}
