package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keygenCmd = &cobra.Command{
	Use:   "keygen",
	Short: "Keypair generation and management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keygenCmd).Standalone()

	keygenCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(keygenCmd)
}
