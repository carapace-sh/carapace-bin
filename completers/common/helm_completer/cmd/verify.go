package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var verifyCmd = &cobra.Command{
	Use:     "verify",
	Short:   "verify that a chart at the given path has been signed and is valid",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(verifyCmd).Standalone()
	verifyCmd.Flags().String("keyring", "/home/user/.gnupg/pubring.gpg", "keyring containing public keys")
	rootCmd.AddCommand(verifyCmd)

	carapace.Gen(verifyCmd).FlagCompletion(carapace.ActionMap{
		"keyring": carapace.ActionFiles(),
	})

	carapace.Gen(verifyCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
