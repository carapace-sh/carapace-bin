package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var kdfCmd = &cobra.Command{
	Use:     "kdf",
	Short:   "Key Derivation Functions",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kdfCmd).Standalone()

	kdfCmd.Flags().BoolS("binary", "binary", false, "Output in binary format (default is hexadecimal)")
	kdfCmd.Flags().StringS("cipher", "cipher", "", "Cipher")
	kdfCmd.Flags().StringS("digest", "digest", "", "Digest")
	kdfCmd.Flags().StringSliceS("kdfopt", "kdfopt", nil, "KDF algorithm control parameters in n:v form")
	kdfCmd.Flags().StringS("keylen", "keylen", "", "The size of the output derived key")
	kdfCmd.Flags().StringS("mac", "mac", "", "MAC See 'Supported Controls' in the EVP_KDF_ docs")
	kdfCmd.Flags().StringS("out", "out", "", "Output to filename rather than stdout")
	common.AddProviderFlags(kdfCmd)
	rootCmd.AddCommand(kdfCmd)

	carapace.Gen(kdfCmd).FlagCompletion(carapace.ActionMap{
		"cipher": action.ActionCipherAlgorithms(kdfCmd),
		"digest": action.ActionDigestAlgorithms(kdfCmd),
		"kdfopt": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues(
					"key",
					"hexkey",
					"pass",
					"hexpass",
					"salt",
					"hexsalt",
					"info",
					"hexinfo",
					"digest",
					"cipher",
					"mac",
				).Suffix(":")
			case 1:
				switch c.Parts[0] {
				case "digest":
					return action.ActionDigestAlgorithms(kdfCmd)
				case "cipher":
					return action.ActionCipherAlgorithms(kdfCmd)
				case "mac":
					return action.ActionMACAlgorithms(kdfCmd)
				default:
					return carapace.ActionValues()
				}
			default:
				return carapace.ActionValues()
			}
		}),
		"mac": action.ActionMACAlgorithms(kdfCmd),
		"out": carapace.ActionFiles(),
	})

	carapace.Gen(kdfCmd).PositionalCompletion(
		action.ActionKDFAlgorithms(kdfCmd),
	)
}
