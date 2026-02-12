package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "Print the encryption key for transfer to another machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keyCmd).Standalone()

	keyCmd.Flags().Bool("base64", false, "Switch to base64 output of the key")
	keyCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(keyCmd)
}
