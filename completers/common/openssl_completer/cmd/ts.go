package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tsCmd = &cobra.Command{
	Use:     "ts",
	Short:   "Time Stamping Authority command",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tsCmd).Standalone()

	tsCmd.Flags().StringS("CAfile", "CAfile", "", "File with trusted CA certs")
	tsCmd.Flags().StringS("CApath", "CApath", "", "Path to trusted CA files")
	tsCmd.Flags().StringS("CAstore", "CAstore", "", "URI to trusted CA store")
	tsCmd.Flags().BoolS("cert", "cert", false, "Put cert request into query")
	tsCmd.Flags().StringS("chain", "chain", "", "File with signer CA chain")
	tsCmd.Flags().StringS("config", "config", "", "Configuration file")
	tsCmd.Flags().StringS("data", "data", "", "File to hash")
	tsCmd.Flags().StringS("digest", "digest", "", "Digest (as a hex string)")
	tsCmd.Flags().StringS("in", "in", "", "Input file")
	tsCmd.Flags().StringS("inkey", "inkey", "", "File with private key for reply")
	tsCmd.Flags().BoolS("no_nonce", "no_nonce", false, "Do not include a nonce")
	tsCmd.Flags().StringS("out", "out", "", "Output file")
	tsCmd.Flags().StringS("passin", "passin", "", "Input file pass phrase source")
	tsCmd.Flags().BoolS("query", "query", false, "Generate a TS query")
	tsCmd.Flags().StringS("queryfile", "queryfile", "", "File containing a TS query")
	tsCmd.Flags().BoolS("reply", "reply", false, "Generate a TS reply")
	tsCmd.Flags().StringS("section", "section", "", "Section to use within config file")
	tsCmd.Flags().StringS("signer", "signer", "", "Signer certificate file")
	tsCmd.Flags().BoolS("text", "text", false, "Output text (not DER)")
	tsCmd.Flags().BoolS("token_in", "token_in", false, "Input is a PKCS#7 file")
	tsCmd.Flags().BoolS("token_out", "token_out", false, "Output is a PKCS#7 file")
	tsCmd.Flags().StringS("tspolicy", "tspolicy", "", "Policy OID to use")
	tsCmd.Flags().StringS("untrusted", "untrusted", "", "Extra untrusted certs")
	tsCmd.Flags().BoolS("verify", "verify", false, "Verify a TS response")
	rootCmd.AddCommand(tsCmd)

	carapace.Gen(tsCmd).FlagCompletion(carapace.ActionMap{
		"CAfile":    carapace.ActionFiles(),
		"CApath":    carapace.ActionDirectories(),
		"chain":     carapace.ActionFiles(),
		"config":    carapace.ActionFiles(),
		"data":      carapace.ActionFiles(),
		"in":        carapace.ActionFiles(),
		"inkey":     carapace.ActionFiles(),
		"out":       carapace.ActionFiles(),
		"queryfile": carapace.ActionFiles(),
		"signer":    carapace.ActionFiles(),
	})
}
