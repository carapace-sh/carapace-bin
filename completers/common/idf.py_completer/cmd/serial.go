package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/espidf"
	"github.com/spf13/cobra"
)

func addFlashFlags(cmd *cobra.Command) {
	cmd.Flags().BoolP("all", "a", false, "Flash all data")
	cmd.Flags().String("extra-args", "", "Pass extra arguments to esptool separated by space")
	cmd.Flags().Bool("force", false, "Force write, skip security and compatibility checks")
	cmd.Flags().Bool("trace", false, "Enable trace-level output of flasher tool interactions")
	cmd.Flags().BoolP("trust-flash-content", "t", false, "Skip MD5 verification of files which do not need reflashing")
}

func addEfuseFlags(cmd *cobra.Command) {
	cmd.Flags().String("before", "", "What to do before connecting to the chip")
	cmd.Flags().BoolP("debug", "d", false, "Print debug information")
	cmd.Flags().Bool("do-not-confirm", false, "Do not pause for confirmation before permanently writing eFuses")
	cmd.Flags().Bool("virt", false, "Work in virtual mode without connecting to a chip")

	cmd.Flag("virt").Hidden = true

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"before": carapace.ActionValues("default-reset", "no-reset", "no-reset-no-sync", "usb-reset"),
	})
}

func init() {
	appFlashCmd := &cobra.Command{
		Use:   "app-flash",
		Short: "Flash the app only",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(appFlashCmd).Standalone()
	addFlashFlags(appFlashCmd)
	rootCmd.AddCommand(appFlashCmd)

	bootloaderFlashCmd := &cobra.Command{
		Use:   "bootloader-flash",
		Short: "Flash bootloader only",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(bootloaderFlashCmd).Standalone()
	addFlashFlags(bootloaderFlashCmd)
	rootCmd.AddCommand(bootloaderFlashCmd)

	efuseBurnCmd := &cobra.Command{
		Use:   "efuse-burn",
		Short: "Burn eFuse",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(efuseBurnCmd).Standalone()
	addEfuseFlags(efuseBurnCmd)
	rootCmd.AddCommand(efuseBurnCmd)

	efuseBurnKeyCmd := &cobra.Command{
		Use:   "efuse-burn-key",
		Short: "Burn eFuse key",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(efuseBurnKeyCmd).Standalone()
	addEfuseFlags(efuseBurnKeyCmd)
	efuseBurnKeyCmd.Flags().Bool("force-write-always", false, "Write the eFuse even if already written or write protected")
	efuseBurnKeyCmd.Flags().Bool("no-protect-key", false, "Disable default read- and write-protecting of the key")
	efuseBurnKeyCmd.Flags().Bool("show-sensitive-info", false, "Show data to be burned")
	rootCmd.AddCommand(efuseBurnKeyCmd)

	efuseDumpCmd := &cobra.Command{
		Use:   "efuse-dump",
		Short: "Dump eFuse blocks",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(efuseDumpCmd).Standalone()
	addEfuseFlags(efuseDumpCmd)
	efuseDumpCmd.Flags().String("file-name", "", "Saves dump for each block into separate file")
	carapace.Gen(efuseDumpCmd).FlagCompletion(carapace.ActionMap{
		"file-name": carapace.ActionFiles(),
	})
	rootCmd.AddCommand(efuseDumpCmd)

	efuseReadProtectCmd := &cobra.Command{
		Use:   "efuse-read-protect",
		Short: "Read-protect eFuse",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(efuseReadProtectCmd).Standalone()
	addEfuseFlags(efuseReadProtectCmd)
	rootCmd.AddCommand(efuseReadProtectCmd)

	efuseSummaryCmd := &cobra.Command{
		Use:   "efuse-summary",
		Short: "Show eFuse summary",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(efuseSummaryCmd).Standalone()
	addEfuseFlags(efuseSummaryCmd)
	efuseSummaryCmd.Flags().String("format", "", "Summary format")
	carapace.Gen(efuseSummaryCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json", "summary", "value-only"),
	})
	rootCmd.AddCommand(efuseSummaryCmd)

	efuseWriteProtectCmd := &cobra.Command{
		Use:   "efuse-write-protect",
		Short: "Write-protect eFuse",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(efuseWriteProtectCmd).Standalone()
	addEfuseFlags(efuseWriteProtectCmd)
	rootCmd.AddCommand(efuseWriteProtectCmd)

	encryptedAppFlashCmd := &cobra.Command{
		Use:   "encrypted-app-flash",
		Short: "Flash the encrypted app only",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(encryptedAppFlashCmd).Standalone()
	addFlashFlags(encryptedAppFlashCmd)
	rootCmd.AddCommand(encryptedAppFlashCmd)

	encryptedFlashCmd := &cobra.Command{
		Use:   "encrypted-flash",
		Short: "Flash the encrypted project",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(encryptedFlashCmd).Standalone()
	addFlashFlags(encryptedFlashCmd)
	rootCmd.AddCommand(encryptedFlashCmd)

	eraseFlashCmd := &cobra.Command{
		Use:   "erase-flash",
		Short: "Erase entire flash chip",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(eraseFlashCmd).Standalone()
	rootCmd.AddCommand(eraseFlashCmd)

	eraseOtadataCmd := &cobra.Command{
		Use:   "erase-otadata",
		Short: "Erase otadata partition",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(eraseOtadataCmd).Standalone()
	rootCmd.AddCommand(eraseOtadataCmd)

	flashCmd := &cobra.Command{
		Use:   "flash",
		Short: "Flash the project",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(flashCmd).Standalone()
	addFlashFlags(flashCmd)
	rootCmd.AddCommand(flashCmd)

	mergeBinCmd := &cobra.Command{
		Use:   "merge-bin",
		Short: "Merge binaries into a single image",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(mergeBinCmd).Standalone()
	mergeBinCmd.Flags().String("flash-offset", "", "Flash offset where the output file will be flashed")
	mergeBinCmd.Flags().StringP("format", "f", "", "Format of the output file")
	mergeBinCmd.Flags().Bool("md5-disable", false, "Disable MD5 checksum in UF2 output")
	mergeBinCmd.Flags().StringP("output", "o", "", "Output filename")
	mergeBinCmd.Flags().String("pad-to-size", "", "Pad final binary file with FF bytes up to this flash size")
	carapace.Gen(mergeBinCmd).FlagCompletion(carapace.ActionMap{
		"format":      carapace.ActionValues("hex", "raw", "uf2"),
		"output":      carapace.ActionFiles(),
		"pad-to-size": carapace.ActionValues("256KB", "512KB", "1MB", "2MB", "4MB", "8MB", "16MB", "32MB", "64MB", "128MB"),
	})
	rootCmd.AddCommand(mergeBinCmd)

	monitorCmd := &cobra.Command{
		Use:   "monitor",
		Short: "Display serial output",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(monitorCmd).Standalone()
	monitorCmd.Flags().StringP("command-file", "c", "", "Run IDF Monitor in command-stream mode from FILE")
	monitorCmd.Flags().Bool("disable-auto-color", false, "Disable auto coloring logs")
	monitorCmd.Flags().BoolP("encrypted", "E", false, "Enable encrypted flash targets")
	monitorCmd.Flags().Bool("force-color", false, "Always print ANSI for colors")
	monitorCmd.Flags().String("monitor-baud", "", "Baud rate for monitor")
	monitorCmd.Flags().Bool("no-reset", false, "Disable reset on monitor startup")
	monitorCmd.Flags().String("print-filter", "", "Filter monitor output")
	monitorCmd.Flags().String("timestamp-format", "", "Set formatting of timestamps compatible with strftime()")
	monitorCmd.Flags().Bool("timestamps", false, "Print a time stamp in the beginning of each line")
	carapace.Gen(monitorCmd).FlagCompletion(carapace.ActionMap{
		"command-file": carapace.ActionFiles(),
		"monitor-baud": espidf.ActionBaudRates(),
	})
	rootCmd.AddCommand(monitorCmd)

	partitionTableFlashCmd := &cobra.Command{
		Use:   "partition-table-flash",
		Short: "Flash partition table only",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(partitionTableFlashCmd).Standalone()
	addFlashFlags(partitionTableFlashCmd)
	rootCmd.AddCommand(partitionTableFlashCmd)

	readOtadataCmd := &cobra.Command{
		Use:   "read-otadata",
		Short: "Read otadata partition",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(readOtadataCmd).Standalone()
	rootCmd.AddCommand(readOtadataCmd)

	secureDecryptFlashDataCmd := &cobra.Command{
		Use:   "secure-decrypt-flash-data",
		Short: "Decrypt flash data",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(secureDecryptFlashDataCmd).Standalone()
	secureDecryptFlashDataCmd.Flags().StringP("address", "a", "", "Address offset in flash")
	secureDecryptFlashDataCmd.Flags().BoolP("aes-xts", "x", false, "Decrypt data using AES-XTS")
	secureDecryptFlashDataCmd.Flags().String("flash-crypt-conf", "", "Override FLASH_CRYPT_CONF eFuse value")
	secureDecryptFlashDataCmd.Flags().StringP("keyfile", "k", "", "File with flash encryption key")
	secureDecryptFlashDataCmd.Flags().StringP("output", "o", "", "Output file for plaintext data")
	carapace.Gen(secureDecryptFlashDataCmd).FlagCompletion(carapace.ActionMap{
		"keyfile": carapace.ActionFiles(),
		"output":  carapace.ActionFiles(),
	})
	carapace.Gen(secureDecryptFlashDataCmd).PositionalCompletion(carapace.ActionFiles())
	rootCmd.AddCommand(secureDecryptFlashDataCmd)

	secureDigestSecureBootloaderCmd := &cobra.Command{
		Use:   "secure-digest-secure-bootloader",
		Short: "Generate secure bootloader digest",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(secureDigestSecureBootloaderCmd).Standalone()
	secureDigestSecureBootloaderCmd.Flags().String("iv", "", "128 byte IV file")
	secureDigestSecureBootloaderCmd.Flags().StringP("keyfile", "k", "", "256 bit key for secure boot digest")
	secureDigestSecureBootloaderCmd.Flags().StringP("output", "o", "", "Output file for signed digest image")
	carapace.Gen(secureDigestSecureBootloaderCmd).FlagCompletion(carapace.ActionMap{
		"iv":      carapace.ActionFiles(),
		"keyfile": carapace.ActionFiles(),
		"output":  carapace.ActionFiles(),
	})
	carapace.Gen(secureDigestSecureBootloaderCmd).PositionalCompletion(carapace.ActionFiles())
	rootCmd.AddCommand(secureDigestSecureBootloaderCmd)

	secureEncryptFlashDataCmd := &cobra.Command{
		Use:   "secure-encrypt-flash-data",
		Short: "Encrypt flash data",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

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

	secureEncryptNvsPartitionCmd := &cobra.Command{
		Use:   "secure-encrypt-nvs-partition",
		Short: "Encrypt NVS partition",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(secureEncryptNvsPartitionCmd).Standalone()
	secureEncryptNvsPartitionCmd.Flags().StringP("keyfile", "k", "", "File with NVS partition key")
	carapace.Gen(secureEncryptNvsPartitionCmd).FlagCompletion(carapace.ActionMap{
		"keyfile": carapace.ActionFiles(),
	})
	carapace.Gen(secureEncryptNvsPartitionCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
	rootCmd.AddCommand(secureEncryptNvsPartitionCmd)

	secureGenerateFlashEncryptionKeyCmd := &cobra.Command{
		Use:   "secure-generate-flash-encryption-key",
		Short: "Generate flash encryption key",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(secureGenerateFlashEncryptionKeyCmd).Standalone()
	secureGenerateFlashEncryptionKeyCmd.Flags().StringP("keylen", "l", "", "Length of private key digest file (in bits)")
	carapace.Gen(secureGenerateFlashEncryptionKeyCmd).FlagCompletion(carapace.ActionMap{
		"keylen": carapace.ActionValues("192", "256", "512"),
	})
	carapace.Gen(secureGenerateFlashEncryptionKeyCmd).PositionalCompletion(carapace.ActionFiles())
	rootCmd.AddCommand(secureGenerateFlashEncryptionKeyCmd)

	secureGenerateKeyDigestCmd := &cobra.Command{
		Use:   "secure-generate-key-digest",
		Short: "Generate key digest",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(secureGenerateKeyDigestCmd).Standalone()
	secureGenerateKeyDigestCmd.Flags().StringP("keyfile", "k", "", "Public key file for digest generation")
	secureGenerateKeyDigestCmd.Flags().StringP("output", "o", "", "Output file for key digest")
	carapace.Gen(secureGenerateKeyDigestCmd).FlagCompletion(carapace.ActionMap{
		"keyfile": carapace.ActionFiles(),
		"output":  carapace.ActionFiles(),
	})
	rootCmd.AddCommand(secureGenerateKeyDigestCmd)

	secureGenerateNvsPartitionKeyCmd := &cobra.Command{
		Use:   "secure-generate-nvs-partition-key",
		Short: "Generate NVS partition key",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(secureGenerateNvsPartitionKeyCmd).Standalone()
	secureGenerateNvsPartitionKeyCmd.Flags().StringP("encryption-scheme", "s", "", "Encryption scheme to use")
	secureGenerateNvsPartitionKeyCmd.Flags().StringP("hmac-keyfile", "l", "", "File to store the generated HMAC key")
	secureGenerateNvsPartitionKeyCmd.Flags().StringP("keyfile", "k", "", "File to store the generated key")
	carapace.Gen(secureGenerateNvsPartitionKeyCmd).FlagCompletion(carapace.ActionMap{
		"encryption-scheme": carapace.ActionValues("Flash", "HMAC"),
		"hmac-keyfile":      carapace.ActionFiles(),
		"keyfile":           carapace.ActionFiles(),
	})
	rootCmd.AddCommand(secureGenerateNvsPartitionKeyCmd)

	secureGenerateSigningKeyCmd := &cobra.Command{
		Use:   "secure-generate-signing-key",
		Short: "Generate secure boot signing key",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(secureGenerateSigningKeyCmd).Standalone()
	secureGenerateSigningKeyCmd.Flags().StringP("scheme", "s", "", "Scheme of secure boot signing")
	secureGenerateSigningKeyCmd.Flags().String("version", "", "Version of the secure boot signing scheme")
	carapace.Gen(secureGenerateSigningKeyCmd).FlagCompletion(carapace.ActionMap{
		"scheme":  carapace.ActionValues("ecdsa256", "ecdsa384", "rsa3072"),
		"version": carapace.ActionValues("1", "2"),
	})
	carapace.Gen(secureGenerateSigningKeyCmd).PositionalCompletion(carapace.ActionFiles())
	rootCmd.AddCommand(secureGenerateSigningKeyCmd)

	secureSignDataCmd := &cobra.Command{
		Use:   "secure-sign-data",
		Short: "Sign data for secure boot",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(secureSignDataCmd).Standalone()
	secureSignDataCmd.Flags().BoolP("append-signatures", "a", false, "Append signature block(s) to already signed image")
	secureSignDataCmd.Flags().StringP("keyfile", "k", "", "Private key file for signing")
	secureSignDataCmd.Flags().StringP("output", "o", "", "Output file for signed digest image")
	secureSignDataCmd.Flags().String("pub-key", "", "Public key files corresponding to the private key")
	secureSignDataCmd.Flags().String("signature", "", "Pre-calculated signatures")
	secureSignDataCmd.Flags().String("version", "", "Version of the secure boot signing scheme")
	carapace.Gen(secureSignDataCmd).FlagCompletion(carapace.ActionMap{
		"keyfile":   carapace.ActionFiles(),
		"output":    carapace.ActionFiles(),
		"pub-key":   carapace.ActionFiles(),
		"signature": carapace.ActionFiles(),
		"version":   carapace.ActionValues("1", "2"),
	})
	carapace.Gen(secureSignDataCmd).PositionalCompletion(carapace.ActionFiles())
	rootCmd.AddCommand(secureSignDataCmd)

	secureVerifySignatureCmd := &cobra.Command{
		Use:   "secure-verify-signature",
		Short: "Verify secure boot signature",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(secureVerifySignatureCmd).Standalone()
	secureVerifySignatureCmd.Flags().StringP("keyfile", "k", "", "Public key file for verification")
	secureVerifySignatureCmd.Flags().String("version", "", "Version of the secure boot signing scheme")
	carapace.Gen(secureVerifySignatureCmd).FlagCompletion(carapace.ActionMap{
		"keyfile": carapace.ActionFiles(),
		"version": carapace.ActionValues("1", "2"),
	})
	carapace.Gen(secureVerifySignatureCmd).PositionalCompletion(carapace.ActionFiles())
	rootCmd.AddCommand(secureVerifySignatureCmd)
}
