package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var storeutlCmd = &cobra.Command{
	Use:     "storeutl",
	Short:   "Command to list and display certificates, keys, CRLs, etc",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storeutlCmd).Standalone()

	storeutlCmd.Flags().StringS("alias", "alias", "", "Search by alias")
	storeutlCmd.Flags().BoolS("certs", "certs", false, "Search for certificates only")
	storeutlCmd.Flags().BoolS("crls", "crls", false, "Search for CRLs only")
	storeutlCmd.Flags().StringS("engine", "engine", "", "Use engine, possibly a hardware device")
	storeutlCmd.Flags().StringS("fingerprint", "fingerprint", "", "Search by public key fingerprint, given in hex")
	storeutlCmd.Flags().StringS("issuer", "issuer", "", "Search by issuer and serial, issuer name")
	storeutlCmd.Flags().BoolS("keys", "keys", false, "Search for keys only")
	storeutlCmd.Flags().BoolS("noout", "noout", false, "No PEM output, just status")
	storeutlCmd.Flags().StringS("out", "out", "", "Output file - default stdout")
	storeutlCmd.Flags().StringS("passin", "passin", "", "Input file pass phrase source")
	storeutlCmd.Flags().BoolS("r", "r", false, "Recurse through names")
	storeutlCmd.Flags().StringS("serial", "serial", "", "Search by issuer and serial, serial number")
	storeutlCmd.Flags().StringS("subject", "subject", "", "Search by subject")
	storeutlCmd.Flags().BoolS("text", "text", false, "Print a text form of the objects")
	rootCmd.AddCommand(storeutlCmd)

	carapace.Gen(storeutlCmd).FlagCompletion(carapace.ActionMap{
		"engine": action.ActionEngines(),
		"out":    carapace.ActionFiles(),
	})
}
