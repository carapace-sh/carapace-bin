package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var pkcs7Cmd = &cobra.Command{
	Use:     "pkcs7",
	Short:   "PKCS#7 Data Management",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pkcs7Cmd).Standalone()

	pkcs7Cmd.Flags().StringS("in", "in", "", "Input file")
	pkcs7Cmd.Flags().StringS("inform", "inform", "", "Input format - DER or PEM")
	pkcs7Cmd.Flags().BoolS("noout", "noout", false, "Don't output encoded data")
	pkcs7Cmd.Flags().StringS("out", "out", "", "Output file")
	pkcs7Cmd.Flags().StringS("outform", "outform", "", "Output format - DER or PEM")
	pkcs7Cmd.Flags().BoolS("print", "print", false, "Print out all fields of the PKCS7 structure")
	pkcs7Cmd.Flags().BoolS("print_certs", "print_certs", false, "Print_certs  print any certs or crl in the input")
	pkcs7Cmd.Flags().BoolS("quiet", "quiet", false, "When used with -print_certs, it produces a cleaner output")
	pkcs7Cmd.Flags().BoolS("text", "text", false, "Print full details of certificates")
	common.AddProviderFlags(pkcs7Cmd)
	rootCmd.AddCommand(pkcs7Cmd)

	carapace.Gen(pkcs7Cmd).FlagCompletion(carapace.ActionMap{
		"in":      carapace.ActionFiles(),
		"inform":  carapace.ActionValues("DER", "PEM"),
		"out":     carapace.ActionFiles(),
		"outform": carapace.ActionValues("DER", "PEM"),
	})
}
