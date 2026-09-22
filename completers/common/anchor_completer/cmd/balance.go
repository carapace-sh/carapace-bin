package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "Get your balance",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(balanceCmd).Standalone()

	balanceCmd.Flags().BoolP("help", "h", false, "Print help")
	balanceCmd.Flags().Bool("lamports", false, "Display balance in lamports instead of SOL")
	rootCmd.AddCommand(balanceCmd)
}
