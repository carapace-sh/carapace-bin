package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secureGenerateFlashEncryptionKeyCmd = &cobra.Command{
	Use:   "secure-generate-flash-encryption-key",
	Short: "Generate flash encryption key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secureGenerateFlashEncryptionKeyCmd).Standalone()

	secureGenerateFlashEncryptionKeyCmd.Flags().StringP("keylen", "l", "", "Length of private key digest file (in bits)")
	rootCmd.AddCommand(secureGenerateFlashEncryptionKeyCmd)

	carapace.Gen(secureGenerateFlashEncryptionKeyCmd).FlagCompletion(carapace.ActionMap{
		"keylen": carapace.ActionValues("192", "256", "512"),
	})

	carapace.Gen(secureGenerateFlashEncryptionKeyCmd).PositionalCompletion(carapace.ActionFiles())
}
