package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var airdropCmd = &cobra.Command{
	Use:   "airdrop",
	Short: "Request an airdrop of SOL",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(airdropCmd).Standalone()

	airdropCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(airdropCmd)
}
