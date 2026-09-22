package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keygen_pubkeyCmd = &cobra.Command{
	Use:   "pubkey",
	Short: "Display the pubkey for a given keypair",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keygen_pubkeyCmd).Standalone()

	keygen_pubkeyCmd.Flags().BoolP("help", "h", false, "Print help")
	keygenCmd.AddCommand(keygen_pubkeyCmd)

	carapace.Gen(keygen_pubkeyCmd).PositionalCompletion(
		carapace.ActionFiles(".json"),
	)
}
