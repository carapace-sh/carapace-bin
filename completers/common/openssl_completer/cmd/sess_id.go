package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sess_idCmd = &cobra.Command{
	Use:     "sess_id",
	Short:   "SSL/TLS session handling command",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sess_idCmd).Standalone()

	sess_idCmd.Flags().BoolS("cert", "cert", false, "Output certificate")
	sess_idCmd.Flags().StringS("context", "context", "", "Set the session ID context")
	sess_idCmd.Flags().StringS("in", "in", "", "Input file - default stdin")
	sess_idCmd.Flags().StringS("inform", "inform", "", "Input format - default PEM (DER or PEM)")
	sess_idCmd.Flags().BoolS("noout", "noout", false, "Don't output the encoded session info")
	sess_idCmd.Flags().StringS("out", "out", "", "Output file - default stdout")
	sess_idCmd.Flags().StringS("outform", "outform", "", "Output format - default PEM (PEM, DER or NSS)")
	sess_idCmd.Flags().BoolS("text", "text", false, "Print ssl session id details")
	rootCmd.AddCommand(sess_idCmd)

	carapace.Gen(sess_idCmd).FlagCompletion(carapace.ActionMap{
		"in":      carapace.ActionFiles(),
		"inform":  carapace.ActionValues("DER", "PEM"),
		"out":     carapace.ActionFiles(),
		"outform": carapace.ActionValues("DER", "PEM", "NSS"),
	})
}
