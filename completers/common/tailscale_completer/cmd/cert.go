package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var certCmd = &cobra.Command{
	Use:   "cert",
	Short: "Get TLS certs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(certCmd).Standalone()

	certCmd.Flags().String("cert-file", "", "output cert file or \"-\" for stdout")
	certCmd.Flags().String("key-file", "", "output key file or \"-\" for stdout")
	certCmd.Flags().String("min-validity", "", "ensure the certificate is valid for at least this duration")
	certCmd.Flags().Bool("serve-demo", false, "serve on port :443 using the cert as a demo instead of writing files to disk")
	rootCmd.AddCommand(certCmd)
}
