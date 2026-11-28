package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secureDigestSecureBootloaderCmd = &cobra.Command{
	Use:   "secure-digest-secure-bootloader",
	Short: "Generate secure bootloader digest",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secureDigestSecureBootloaderCmd).Standalone()

	secureDigestSecureBootloaderCmd.Flags().String("iv", "", "128 byte IV file")
	secureDigestSecureBootloaderCmd.Flags().StringP("keyfile", "k", "", "256 bit key for secure boot digest")
	secureDigestSecureBootloaderCmd.Flags().StringP("output", "o", "", "Output file for signed digest image")
	rootCmd.AddCommand(secureDigestSecureBootloaderCmd)

	carapace.Gen(secureDigestSecureBootloaderCmd).FlagCompletion(carapace.ActionMap{
		"iv":      carapace.ActionFiles(),
		"keyfile": carapace.ActionFiles(),
		"output":  carapace.ActionFiles(),
	})

	carapace.Gen(secureDigestSecureBootloaderCmd).PositionalCompletion(carapace.ActionFiles())
}
