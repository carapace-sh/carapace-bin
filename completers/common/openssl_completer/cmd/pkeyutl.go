package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var pkeyutlCmd = &cobra.Command{
	Use:     "pkeyutl",
	Short:   "Public key algorithm cryptographic operation command",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkeyutlCmd).Standalone()

	pkeyutlCmd.Flags().BoolS("asn1parse", "asn1parse", false, "parse the output as ASN.1 data to check its DER encoding and print errors")
	pkeyutlCmd.Flags().BoolS("certin", "certin", false, "Input is a cert with a public key")
	pkeyutlCmd.Flags().StringS("config", "config", "", "Load a configuration file (this may load modules)")
	pkeyutlCmd.Flags().BoolS("decap", "decap", false, "Decapsulate shared secret")
	pkeyutlCmd.Flags().BoolS("decrypt", "decrypt", false, "Decrypt input data with private key")
	pkeyutlCmd.Flags().BoolS("derive", "derive", false, "Derive shared secret from own and peer (EC)DH keys")
	pkeyutlCmd.Flags().StringS("digest", "digest", "", "The digest algorithm to use for signing/verifying raw input data. Implies -rawin")
	pkeyutlCmd.Flags().BoolS("encap", "encap", false, "Encapsulate shared secret")
	pkeyutlCmd.Flags().BoolS("encrypt", "encrypt", false, "Encrypt input data with public key")
	pkeyutlCmd.Flags().StringS("engine", "engine", "", "Use engine, possibly a hardware device")
	pkeyutlCmd.Flags().BoolS("engine_impl", "engine_impl", false, "Also use engine given by -engine for crypto operations")
	pkeyutlCmd.Flags().BoolS("hexdump", "hexdump", false, "Hex dump output")
	pkeyutlCmd.Flags().StringS("in", "in", "", "Input file - default stdin")
	pkeyutlCmd.Flags().StringS("inkey", "inkey", "", "Input key, by default private key")
	pkeyutlCmd.Flags().StringS("kdf", "kdf", "", "Use KDF algorithm")
	pkeyutlCmd.Flags().StringS("kdflen", "kdflen", "", "KDF algorithm output length")
	pkeyutlCmd.Flags().StringS("kemop", "kemop", "", "KEM operation specific to the key algorithm")
	pkeyutlCmd.Flags().StringS("keyform", "keyform", "", "Private key format (ENGINE, other values ignored)")
	pkeyutlCmd.Flags().StringS("out", "out", "", "Output file - default stdout")
	pkeyutlCmd.Flags().StringS("passin", "passin", "", "Input file pass phrase source")
	pkeyutlCmd.Flags().StringS("peerform", "peerform", "", "Peer key format (DER/PEM/P12/ENGINE)")
	pkeyutlCmd.Flags().StringS("peerkey", "peerkey", "", "Peer key file used in key derivation")
	pkeyutlCmd.Flags().StringSliceS("pkeyopt", "pkeyopt", nil, "Public key options as opt:value")
	pkeyutlCmd.Flags().StringS("pkeyopt_passin", "pkeyopt_passin", "", "Public key option that is read as a passphrase argument opt:passphrase")
	pkeyutlCmd.Flags().BoolS("pubin", "pubin", false, "Input key is a public key")
	pkeyutlCmd.Flags().BoolS("rawin", "rawin", false, "Indicate that the signature/verification input data is not yet hashed")
	pkeyutlCmd.Flags().BoolS("rev", "rev", false, "Reverse the order of the input buffer")
	pkeyutlCmd.Flags().StringS("secret", "secret", "", "File to store secret on encapsulation")
	pkeyutlCmd.Flags().StringS("sigfile", "sigfile", "", "Signature file (verify operation only)")
	pkeyutlCmd.Flags().BoolS("sign", "sign", false, "Sign input data with private key")
	pkeyutlCmd.Flags().BoolS("verify", "verify", false, "Verify with public key")
	pkeyutlCmd.Flags().BoolS("verifyrecover", "verifyrecover", false, "Verify RSA signature, recovering original signature input data")
	common.AddProviderFlags(pkeyutlCmd)
	common.AddRandomStateFlags(pkeyutlCmd)
	rootCmd.AddCommand(pkeyutlCmd)

	carapace.Gen(pkeyutlCmd).FlagCompletion(carapace.ActionMap{
		"config":   carapace.ActionFiles(),
		"engine":   action.ActionEngines(),
		"in":       carapace.ActionFiles(),
		"inkey":    carapace.ActionFiles(),
		"kdf":      carapace.ActionValues("HKDF", "TLS1-PRF"),
		"keyform":  carapace.ActionValues("ENGINE", "DER", "PEM", "P12"),
		"out":      carapace.ActionFiles(),
		"peerform": carapace.ActionValues("ENGINE", "DER", "PEM", "P12"),
		"peerkey":  carapace.ActionFiles(),
		"secret":   carapace.ActionFiles(),
		"sigfile":  carapace.ActionFiles(),
	})
}
