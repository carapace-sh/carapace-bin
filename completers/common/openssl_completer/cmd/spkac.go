package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var spkacCmd = &cobra.Command{
	Use:     "spkac",
	Short:   "SPKAC printing and generating command",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(spkacCmd).Standalone()

	spkacCmd.Flags().StringS("challenge", "challenge", "", "Challenge string")
	spkacCmd.Flags().StringS("digest", "digest", "", "Sign new SPKAC with the specified digest (default: MD5)")
	spkacCmd.Flags().StringS("in", "in", "", "Input file")
	spkacCmd.Flags().StringS("key", "key", "", "Create SPKAC using private key")
	spkacCmd.Flags().StringS("keyform", "keyform", "", "Private key file format (DER/PEM)")
	spkacCmd.Flags().BoolS("noout", "noout", false, "Don't print SPKAC")
	spkacCmd.Flags().StringS("out", "out", "", "Output file")
	spkacCmd.Flags().StringS("passin", "passin", "", "Input file pass phrase source")
	spkacCmd.Flags().BoolS("pubkey", "pubkey", false, "Output public key")
	spkacCmd.Flags().StringS("spkac", "spkac", "", "Alternative SPKAC name")
	spkacCmd.Flags().StringS("spksect", "spksect", "", "Specify the name of an SPKAC-dedicated section of configuration")
	spkacCmd.Flags().BoolS("verify", "verify", false, "Verify SPKAC signature")
	rootCmd.AddCommand(spkacCmd)

	carapace.Gen(spkacCmd).FlagCompletion(carapace.ActionMap{
		"in":      carapace.ActionFiles(),
		"key":     carapace.ActionFiles(),
		"keyform": carapace.ActionValues("DER", "PEM", "P12"),
		"out":     carapace.ActionFiles(),
	})
}
