package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secureDecryptFlashDataCmd = &cobra.Command{
	Use:   "secure-decrypt-flash-data",
	Short: "Decrypt flash data",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secureDecryptFlashDataCmd).Standalone()

	secureDecryptFlashDataCmd.Flags().StringP("address", "a", "", "Address offset in flash")
	secureDecryptFlashDataCmd.Flags().BoolP("aes-xts", "x", false, "Decrypt data using AES-XTS")
	secureDecryptFlashDataCmd.Flags().String("flash-crypt-conf", "", "Override FLASH_CRYPT_CONF eFuse value")
	secureDecryptFlashDataCmd.Flags().StringP("keyfile", "k", "", "File with flash encryption key")
	secureDecryptFlashDataCmd.Flags().StringP("output", "o", "", "Output file for plaintext data")
	rootCmd.AddCommand(secureDecryptFlashDataCmd)

	carapace.Gen(secureDecryptFlashDataCmd).FlagCompletion(carapace.ActionMap{
		"keyfile": carapace.ActionFiles(),
		"output":  carapace.ActionFiles(),
	})

	carapace.Gen(secureDecryptFlashDataCmd).PositionalCompletion(carapace.ActionFiles())
}
