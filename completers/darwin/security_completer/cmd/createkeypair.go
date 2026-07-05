package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/security"
	"github.com/spf13/cobra"
)

var createKeypairCmd = &cobra.Command{
	Use:   "create-keypair",
	Short: "Create an asymmetric key pair",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createKeypairCmd).Standalone()
	createKeypairCmd.Flags().StringP("algorithm", "a", "", "Algorithm: rsa, dh, dsa, or fee (default: rsa)")
	createKeypairCmd.Flags().BoolP("allow-all", "A", false, "Allow any application to access without warning")
	createKeypairCmd.Flags().StringP("days", "d", "", "Make key valid for specified days from today")
	createKeypairCmd.Flags().StringP("keychain", "k", "", "Use specified keychain instead of default")
	createKeypairCmd.Flags().StringP("size", "s", "", "Keysize in bits (default: 512)")
	createKeypairCmd.Flags().StringP("tool", "T", "", "Specify an application which may access this key")
	createKeypairCmd.Flags().StringP("valid-from", "f", "", "Make key valid from specified date")
	createKeypairCmd.Flags().StringP("valid-to", "t", "", "Make key valid to specified date")
	rootCmd.AddCommand(createKeypairCmd)
	carapace.Gen(createKeypairCmd).FlagCompletion(carapace.ActionMap{
		"algorithm": carapace.ActionValues("rsa", "dh", "dsa", "fee"),
		"keychain":  security.ActionKeychains(),
	})
}
