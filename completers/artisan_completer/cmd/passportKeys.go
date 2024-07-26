package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var passportKeysCmd = &cobra.Command{
	Use:   "passport:keys [--force] [--length [LENGTH]]",
	Short: "Create the encryption keys for API authentication",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(passportKeysCmd).Standalone()

	passportKeysCmd.Flags().Bool("force", false, "Overwrite keys they already exist")
	passportKeysCmd.Flags().String("length", "", "The length of the private key")
	rootCmd.AddCommand(passportKeysCmd)
}
