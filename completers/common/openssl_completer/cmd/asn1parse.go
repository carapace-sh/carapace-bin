package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var asn1parseCmd = &cobra.Command{
	Use:     "asn1parse",
	Short:   "Parse an ASN.1 sequence",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(asn1parseCmd).Standalone()

	asn1parseCmd.Flags().StringS("dlimit", "dlimit", "", "dump the first arg bytes of unknown data in hex form")
	asn1parseCmd.Flags().BoolS("dump", "dump", false, "unknown data in hex form")
	asn1parseCmd.Flags().StringS("genconf", "genconf", "", "file to generate ASN1 structure from")
	asn1parseCmd.Flags().StringS("genstr", "genstr", "", "string to generate ASN1 structure from")
	asn1parseCmd.Flags().BoolS("i", "i", false, "indents the output")
	asn1parseCmd.Flags().StringS("in", "in", "", "input file")
	asn1parseCmd.Flags().StringS("inform", "inform", "", "input format - one of DER PEM B64")
	asn1parseCmd.Flags().StringS("item", "item", "", "item to parse and print (-inform  will be ignored)")
	asn1parseCmd.Flags().StringS("length", "length", "", "length of section in file")
	asn1parseCmd.Flags().BoolS("noout", "noout", false, "do not produce any output")
	asn1parseCmd.Flags().StringS("offset", "offset", "", "offset into file")
	asn1parseCmd.Flags().StringS("oid", "oid", "", "file of extra oid definitions")
	asn1parseCmd.Flags().StringS("out", "out", "", "output file (output format is always DER)")
	asn1parseCmd.Flags().BoolS("strictpem", "strictpem", false, "equivalent to '-inform pem' (obsolete)")
	asn1parseCmd.Flags().StringSliceS("strparse", "strparse", nil, "offset; a series of these can be used to 'dig' into multiple ASN1 blob wrappings")
	rootCmd.AddCommand(asn1parseCmd)

	carapace.Gen(asn1parseCmd).FlagCompletion(carapace.ActionMap{
		"genconf": carapace.ActionFiles(),
		"in":      carapace.ActionFiles(),
		"inform":  carapace.ActionValues("B64", "DER", "PEM"),
		"oid":     carapace.ActionFiles(),
		"out":     carapace.ActionFiles(),
	})
}
