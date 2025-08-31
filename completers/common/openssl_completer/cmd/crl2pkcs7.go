package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var crl2pkcs7Cmd = &cobra.Command{
	Use:     "crl2pkcs7",
	Short:   "CRL to PKCS#7 Conversion",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(crl2pkcs7Cmd).Standalone()

	crl2pkcs7Cmd.Flags().StringS("certfile", "certfile", "", "File of chain of certs to a trusted CA; can be repeated")
	crl2pkcs7Cmd.Flags().StringS("in", "in", "", "Input file")
	crl2pkcs7Cmd.Flags().StringS("inform", "inform", "", "Input format - DER or PEM")
	crl2pkcs7Cmd.Flags().BoolS("nocrl", "nocrl", false, "No crl to load, just certs from '-certfile'")
	crl2pkcs7Cmd.Flags().StringS("out", "out", "", "Output file")
	crl2pkcs7Cmd.Flags().StringS("outform", "outform", "", "Output format - DER or PEM")
	common.AddProviderFlags(crl2pkcs7Cmd)
	rootCmd.AddCommand(crl2pkcs7Cmd)

	carapace.Gen(crl2pkcs7Cmd).FlagCompletion(carapace.ActionMap{
		"certfile": carapace.ActionFiles(),
		"in":       carapace.ActionFiles(),
		"inform":   carapace.ActionValues("DER", "PEM"),
		"out":      carapace.ActionFiles(),
		"outform":  carapace.ActionValues("DER", "PEM"),
	})
}
