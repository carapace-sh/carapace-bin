package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gpgKeyLookupCmd = &cobra.Command{
	Use:   "gpg-key-lookup",
	Short: "Fetch the actual ed25519 public key value of the underlying key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gpgKeyLookupCmd).Standalone()
	gpgKeyLookupCmd.Flags().BoolP("help", "h", false, "show this help message and exit")

	rootCmd.AddCommand(gpgKeyLookupCmd)
}
