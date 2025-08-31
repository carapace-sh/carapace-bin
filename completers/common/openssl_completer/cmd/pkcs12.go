package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var pkcs12Cmd = &cobra.Command{
	Use:     "pkcs12",
	Short:   "PKCS#12 Data Management",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkcs12Cmd).Standalone()

	pkcs12Cmd.Flags().StringS("CAfile", "CAfile", "", "PEM-format file of CA's")
	pkcs12Cmd.Flags().StringS("CApath", "CApath", "", "PEM-format directory of CA's")
	pkcs12Cmd.Flags().StringS("CAstore", "CAstore", "", "URI to store of CA's")
	pkcs12Cmd.Flags().StringS("CSP", "CSP", "", "Microsoft CSP name")
	pkcs12Cmd.Flags().BoolS("LMK", "LMK", false, "Add local machine keyset attribute to private key")
	pkcs12Cmd.Flags().BoolS("cacerts", "cacerts", false, "Only output CA certificates")
	pkcs12Cmd.Flags().StringS("caname", "caname", "", "Use name as CA friendly name (can be repeated)")
	pkcs12Cmd.Flags().StringS("certfile", "certfile", "", "Extra certificates for PKCS12 output")
	pkcs12Cmd.Flags().StringS("certpbe", "certpbe", "", "Certificate PBE algorithm (default PBES2 with PBKDF2 and AES-256 CBC)")
	pkcs12Cmd.Flags().BoolS("chain", "chain", false, "Build and add certificate chain for EE cert, which is the 1st cert from -in matching the private key (if given)")
	pkcs12Cmd.Flags().BoolS("clcerts", "clcerts", false, "Only output client certificates")
	pkcs12Cmd.Flags().BoolS("descert", "descert", false, "Encrypt output with 3DES (default PBES2 with PBKDF2 and AES-256 CBC)")
	pkcs12Cmd.Flags().StringS("engine", "engine", "", "Use engine, possibly a hardware device")
	pkcs12Cmd.Flags().BoolS("export", "export", false, "Create PKCS12 file")
	pkcs12Cmd.Flags().StringS("in", "in", "", "Input file")
	pkcs12Cmd.Flags().BoolS("info", "info", false, "Print info about PKCS#12 structure")
	pkcs12Cmd.Flags().StringS("inkey", "inkey", "", "Private key, else read from -in input file")
	pkcs12Cmd.Flags().StringS("iter", "iter", "", "Specify the iteration count for encryption and MAC")
	pkcs12Cmd.Flags().StringS("jdktrust", "jdktrust", "", "Mark certificate in PKCS#12 store as trusted for JDK compatibility")
	pkcs12Cmd.Flags().BoolS("keyex", "keyex", false, "Set key type to MS key exchange")
	pkcs12Cmd.Flags().StringS("keypbe", "keypbe", "", "Private key PBE algorithm (default AES-256 CBC)")
	pkcs12Cmd.Flags().BoolS("keysig", "keysig", false, "Set key type to MS key signature")
	pkcs12Cmd.Flags().BoolS("legacy", "legacy", false, "Use legacy encryption: 3DES_CBC for keys, RC2_CBC for certs")
	pkcs12Cmd.Flags().StringS("macalg", "macalg", "", "Digest algorithm to use in MAC (default SHA256)")
	pkcs12Cmd.Flags().BoolS("maciter", "maciter", false, "Unused, kept for backwards compatibility")
	pkcs12Cmd.Flags().StringS("macsaltlen", "macsaltlen", "", "Specify the salt len for MAC")
	pkcs12Cmd.Flags().StringS("name", "name", "", "Use name as friendly name")
	pkcs12Cmd.Flags().BoolS("no-CAfile", "no-CAfile", false, "Do not load the default certificates file")
	pkcs12Cmd.Flags().BoolS("no-CApath", "no-CApath", false, "Do not load certificates from the default certificates directory")
	pkcs12Cmd.Flags().BoolS("no-CAstore", "no-CAstore", false, "Do not load certificates from the default certificates store")
	pkcs12Cmd.Flags().BoolS("nocerts", "nocerts", false, "Don't output certificates")
	pkcs12Cmd.Flags().BoolS("nodes", "nodes", false, "Don't encrypt private keys; deprecated")
	pkcs12Cmd.Flags().BoolS("noenc", "noenc", false, "Don't encrypt private keys")
	pkcs12Cmd.Flags().BoolS("noiter", "noiter", false, "Don't use encryption iteration")
	pkcs12Cmd.Flags().BoolS("nokeys", "nokeys", false, "Don't output private keys")
	pkcs12Cmd.Flags().BoolS("nomac", "nomac", false, "Don't generate MAC")
	pkcs12Cmd.Flags().BoolS("nomaciter", "nomaciter", false, "Don't use MAC iteration)")
	pkcs12Cmd.Flags().BoolS("nomacver", "nomacver", false, "Don't verify integrity MAC")
	pkcs12Cmd.Flags().BoolS("noout", "noout", false, "Don't output anything, just verify PKCS#12 input")
	pkcs12Cmd.Flags().StringS("out", "out", "", "Output file")
	pkcs12Cmd.Flags().StringS("passcerts", "passcerts", "", "Certificate file pass phrase source")
	pkcs12Cmd.Flags().StringS("passin", "passin", "", "Input file pass phrase source")
	pkcs12Cmd.Flags().StringS("passout", "passout", "", "Output file pass phrase source")
	pkcs12Cmd.Flags().StringS("password", "password", "", "Set PKCS#12 import/export password source")
	pkcs12Cmd.Flags().BoolS("pbmac1_pbkdf2", "pbmac1_pbkdf2", false, "Use PBMAC1 with PBKDF2 instead of MAC")
	pkcs12Cmd.Flags().StringS("pbmac1_pbkdf2_md", "pbmac1_pbkdf2_md", "", "Digest to use for PBMAC1 KDF (default SHA256)")
	pkcs12Cmd.Flags().BoolS("twopass", "twopass", false, "Separate MAC, encryption passwords")
	pkcs12Cmd.Flags().StringS("untrusted", "untrusted", "", "Untrusted certificates for chain building")
	common.AddProviderFlags(pkcs12Cmd)
	common.AddRandomStateFlags(pkcs12Cmd)
	rootCmd.AddCommand(pkcs12Cmd)

	pkcs12Cmd.MarkFlagsMutuallyExclusive("keyex", "keysig")

	carapace.Gen(pkcs12Cmd).FlagCompletion(carapace.ActionMap{
		"CAfile":    carapace.ActionFiles(),
		"CApath":    carapace.ActionDirectories(),
		"certfile":  carapace.ActionFiles(),
		"in":        carapace.ActionFiles(),
		"inkey":     carapace.ActionFiles(),
		"out":       carapace.ActionFiles(),
		"untrusted": carapace.ActionFiles(),
	})
}
