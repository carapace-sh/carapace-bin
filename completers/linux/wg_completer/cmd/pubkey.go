package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pubkeyCmd = &cobra.Command{
	Use:   "pubkey",
	Short: "Reads a private key from stdin and writes a public key to stdout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pubkeyCmd).Standalone()

	rootCmd.AddCommand(pubkeyCmd)
}
