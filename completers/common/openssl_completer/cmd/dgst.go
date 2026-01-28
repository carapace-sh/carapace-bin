package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var dgstCmd = &cobra.Command{
	Use:     "dgst",
	Short:   "Message Digest calculation",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dgstCmd).Standalone()

	dgstCmd.Flags().BoolS("binary", "binary", false, "Print in binary form")
	dgstCmd.Flags().BoolS("c", "c", false, "Print the digest with separating colons")
	dgstCmd.Flags().BoolS("d", "d", false, "Print debug info")
	dgstCmd.Flags().BoolS("debug", "debug", false, "Print debug info")
	dgstCmd.Flags().BoolS("fips-fingerprint", "fips-fingerprint", false, "Compute HMAC with the key used in OpenSSL-FIPS fingerprint")
	dgstCmd.Flags().BoolS("hex", "hex", false, "Print as hex dump")
	dgstCmd.Flags().StringS("hmac", "hmac", "", "Create hashed MAC with key")
	dgstCmd.Flags().StringS("hmac-env", "hmac-env", "", "Create hashed MAC with key from environment variable")
	dgstCmd.Flags().BoolS("hmac-stdin", "hmac-stdin", false, "Create hashed MAC with key from stdin")
	dgstCmd.Flags().StringS("keyform", "keyform", "", "Key file format (DER/PEM)")
	dgstCmd.Flags().BoolS("list", "list", false, "List digests")
	dgstCmd.Flags().StringS("mac", "mac", "", "Create MAC (not necessarily HMAC)")
	dgstCmd.Flags().StringSliceS("macopt", "macopt", nil, "MAC algorithm parameters in n:v form or key")
	dgstCmd.Flags().StringS("out", "out", "", "Output to filename rather than stdout")
	dgstCmd.Flags().StringS("passin", "passin", "", "Input file pass phrase source")
	dgstCmd.Flags().StringS("prverify", "prverify", "", "Verify a signature using private key")
	dgstCmd.Flags().BoolS("r", "r", false, "Print the digest in coreutils format")
	dgstCmd.Flags().StringS("sign", "sign", "", "Sign digest using private key")
	dgstCmd.Flags().StringS("signature", "signature", "", "File with signature to verify")
	dgstCmd.Flags().StringSliceS("sigopt", "sigopt", nil, "Signature parameter in n:v form")
	dgstCmd.Flags().StringS("verify", "verify", "", "Verify a signature using public key")
	dgstCmd.Flags().StringS("xoflen", "xoflen", "", "Output length for XOF algorithms. To obtain the maximum security strength set this to 32 (or greater) for SHAKE128, and 64 (or greater) for SHAKE256")
	common.AddProviderFlags(dgstCmd)
	common.AddRandomStateFlags(dgstCmd)
	rootCmd.AddCommand(dgstCmd)

	carapace.Gen(dgstCmd).FlagCompletion(carapace.ActionMap{
		"keyform":   carapace.ActionValues("DER", "PEM", "P12"),
		"out":       carapace.ActionFiles(),
		"prverify":  carapace.ActionFiles(),
		"sign":      carapace.ActionFiles(),
		"signature": carapace.ActionFiles(),
		"verify":    carapace.ActionFiles(),
	})

	carapace.Gen(dgstCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
