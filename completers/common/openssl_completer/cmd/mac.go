package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var macCmd = &cobra.Command{
	Use:     "mac",
	Short:   "Message Authentication Code Calculation",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macCmd).Standalone()

	macCmd.Flags().BoolS("binary", "binary", false, "Output in binary format (default is hexadecimal)")
	macCmd.Flags().StringS("cipher", "cipher", "", "Cipher")
	macCmd.Flags().StringS("digest", "digest", "", "Digest See 'PARAMETER NAMES' in the EVP_MAC_ docs")
	macCmd.Flags().StringS("in", "in", "", "Input file to MAC (default is stdin)")
	macCmd.Flags().StringSliceS("macopt", "macopt", nil, "MAC algorithm parameters in n:v form")
	macCmd.Flags().StringS("out", "out", "", "Output to filename rather than stdout")
	common.AddProviderFlags(macCmd)
	rootCmd.AddCommand(macCmd)

	carapace.Gen(macCmd).FlagCompletion(carapace.ActionMap{
		"cipher": action.ActionCipherAlgorithms(macCmd),
		"digest": action.ActionDigestAlgorithms(macCmd),
		"in":     carapace.ActionFiles(),
		"macopt": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues(
					"key",
					"hexkey",
					"iv",
					"hexiv",
					"size",
					"custom",
					"digest",
					"cipher",
				).Suffix(":")
			case 1:
				switch c.Parts[0] {
				case "digest":
					return action.ActionDigestAlgorithms(macCmd)
				case "cipher":
					return action.ActionCipherAlgorithms(macCmd)
				default:
					return carapace.ActionValues()
				}
			default:
				return carapace.ActionValues()
			}
		}),
		"out": carapace.ActionFiles(),
	})

	carapace.Gen(macCmd).PositionalCompletion(
		action.ActionMACAlgorithms(macCmd),
	)
}
