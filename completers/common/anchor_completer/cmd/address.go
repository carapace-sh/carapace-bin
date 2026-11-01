package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var addressCmd = &cobra.Command{
	Use:   "address",
	Short: "Get your public key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addressCmd).Standalone()

	addressCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(addressCmd)
}
