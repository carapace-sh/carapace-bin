package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var packslipCmd = &cobra.Command{
	Use:   "packslip",
	Short: "Inspect and reset the signers mise has pinned for packslip tools",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(packslipCmd).Standalone()

	packslipCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(packslipCmd)
}
