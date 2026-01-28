package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var pkcs8Cmd = &cobra.Command{
	Use:     "pkcs8",
	Short:   "PKCS#8 format private key conversion command",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkcs8Cmd).Standalone()

	pkcs8Cmd.Flags().StringS("in", "in", "", "Input file")
	pkcs8Cmd.Flags().StringS("inform", "inform", "", "Input format (DER or PEM)")
	pkcs8Cmd.Flags().StringS("iter", "iter", "", "Specify the iteration count")
	pkcs8Cmd.Flags().BoolS("nocrypt", "nocrypt", false, "Use or expect unencrypted private key")
	pkcs8Cmd.Flags().BoolS("noiter", "noiter", false, "Use 1 as iteration count")
	pkcs8Cmd.Flags().StringS("out", "out", "", "Output file")
	pkcs8Cmd.Flags().StringS("outform", "outform", "", "Output format (DER or PEM)")
	pkcs8Cmd.Flags().StringS("passin", "passin", "", "Input file pass phrase source")
	pkcs8Cmd.Flags().StringS("passout", "passout", "", "Output file pass phrase source")
	pkcs8Cmd.Flags().StringS("saltlen", "saltlen", "", "Specify the salt length (in bytes) Default: 8 (For PBE1) or 16 (for PBE2)")
	pkcs8Cmd.Flags().BoolS("scrypt", "scrypt", false, "Use scrypt algorithm")
	pkcs8Cmd.Flags().StringS("scrypt_N", "scrypt_N", "", "Set scrypt N parameter")
	pkcs8Cmd.Flags().StringS("scrypt_p", "scrypt_p", "", "Set scrypt p parameter")
	pkcs8Cmd.Flags().StringS("scrypt_r", "scrypt_r", "", "Set scrypt r parameter")
	pkcs8Cmd.Flags().BoolS("topk8", "topk8", false, "Output PKCS8 file")
	pkcs8Cmd.Flags().BoolS("traditional", "traditional", false, "use traditional format private key")
	pkcs8Cmd.Flags().StringS("v1", "v1", "", "Use PKCS#5 v1.5 and cipher")
	pkcs8Cmd.Flags().StringS("v2", "v2", "", "Use PKCS#5 v2.0 and cipher")
	pkcs8Cmd.Flags().StringS("v2prf", "v2prf", "", "Set the PRF algorithm to use with PKCS#5 v2.0")
	common.AddProviderFlags(pkcs8Cmd)
	common.AddRandomStateFlags(pkcs8Cmd)
	rootCmd.AddCommand(pkcs8Cmd)

	carapace.Gen(pkcs8Cmd).FlagCompletion(carapace.ActionMap{
		"in":      carapace.ActionFiles(),
		"inform":  carapace.ActionValues("DER", "PEM"),
		"out":     carapace.ActionFiles(),
		"outform": carapace.ActionValues("DER", "PEM"),
	})
}
