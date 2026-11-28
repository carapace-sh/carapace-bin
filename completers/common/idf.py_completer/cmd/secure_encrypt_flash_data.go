package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secureEncryptFlashDataCmd = &cobra.Command{
	Use:   "secure-encrypt-flash-data",
	Short: "Encrypt flash data",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secureEncryptFlashDataCmd).Standalone()
	secureEncryptFlashDataCmd.Flags().StringP("address", "a", "", "Address offset in flash")
	secureEncryptFlashDataCmd.Flags().BoolP("aes-xts", "x", false, "Encrypt data using AES-XTS")
	secureEncryptFlashDataCmd.Flags().String("flash-crypt-conf", "", "Override FLASH_CRYPT_CONF eFuse value")
	secureEncryptFlashDataCmd.Flags().StringP("keyfile", "k", "", "File with flash encryption key")
	secureEncryptFlashDataCmd.Flags().StringP("output", "o", "", "Output file for encrypted data")
	carapace.Gen(secureEncryptFlashDataCmd).FlagCompletion(carapace.ActionMap{
		"keyfile": carapace.ActionFiles(),
		"output":  carapace.ActionFiles(),
	})
	carapace.Gen(secureEncryptFlashDataCmd).PositionalCompletion(carapace.ActionFiles())
	rootCmd.AddCommand(secureEncryptFlashDataCmd)
}
