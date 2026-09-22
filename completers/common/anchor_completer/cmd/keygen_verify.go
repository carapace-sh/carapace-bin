package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keygen_verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify a keypair can sign and verify a message",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keygen_verifyCmd).Standalone()

	keygen_verifyCmd.Flags().BoolP("help", "h", false, "Print help")
	keygenCmd.AddCommand(keygen_verifyCmd)
}
