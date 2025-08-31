package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openssl_completer/cmd/common"
	"github.com/spf13/cobra"
)

var s_timeCmd = &cobra.Command{
	Use:     "s_time",
	Short:   "SSL/TLS performance timing program",
	GroupID: "standard",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s_timeCmd).Standalone()

	s_timeCmd.Flags().StringS("CAfile", "CAfile", "", "PEM format file of CA's")
	s_timeCmd.Flags().StringS("CApath", "CApath", "", "PEM format directory of CA's")
	s_timeCmd.Flags().StringS("CAstore", "CAstore", "", "URI to store of CA's")
	s_timeCmd.Flags().BoolS("bugs", "bugs", false, "Turn on SSL bug compatibility")
	s_timeCmd.Flags().StringS("cafile", "cafile", "", "PEM format file of CA's")
	s_timeCmd.Flags().StringS("cert", "cert", "", "Cert file to use, PEM format assumed")
	s_timeCmd.Flags().StringS("cipher", "cipher", "", "TLSv1.2 and below cipher list to be used")
	s_timeCmd.Flags().StringS("ciphersuites", "ciphersuites", "", "Specify TLSv1.3 ciphersuites to be used")
	s_timeCmd.Flags().StringS("connect", "connect", "", "Where to connect as post:port (default is localhost:4433)")
	s_timeCmd.Flags().StringS("key", "key", "", "File with key, PEM; default is -cert file")
	s_timeCmd.Flags().StringSliceS("nameopt", "nameopt", nil, "Certificate subject/issuer name printing options")
	s_timeCmd.Flags().BoolS("new", "new", false, "Just time new connections")
	s_timeCmd.Flags().BoolS("no-CAfile", "no-CAfile", false, "Do not load the default certificates file")
	s_timeCmd.Flags().BoolS("no-CApath", "no-CApath", false, "Do not load certificates from the default certificates directory")
	s_timeCmd.Flags().BoolS("no-CAstore", "no-CAstore", false, "Do not load certificates from the default certificates store URI")
	s_timeCmd.Flags().BoolS("reuse", "reuse", false, "Just time connection reuse")
	s_timeCmd.Flags().BoolS("ssl3", "ssl3", false, "Just use SSLv3")
	s_timeCmd.Flags().StringS("time", "time", "", "Seconds to collect data, default 30")
	s_timeCmd.Flags().BoolS("tls1", "tls1", false, "Just use TLSv1.0")
	s_timeCmd.Flags().BoolS("tls1_1", "tls1_1", false, "Just use TLSv1.1")
	s_timeCmd.Flags().BoolS("tls1_2", "tls1_2", false, "Just use TLSv1.2")
	s_timeCmd.Flags().BoolS("tls1_3", "tls1_3", false, "Just use TLSv1.3")
	s_timeCmd.Flags().StringS("verify", "verify", "", "Turn on peer certificate verification, set depth")
	s_timeCmd.Flags().StringS("www", "www", "", "Fetch specified page from the site")
	common.AddProviderFlags(s_timeCmd)
	rootCmd.AddCommand(s_timeCmd)

	carapace.Gen(s_timeCmd).FlagCompletion(carapace.ActionMap{
		"CAfile": carapace.ActionFiles(),
		"CApath": carapace.ActionDirectories(),
		"cafile": carapace.ActionFiles(),
		"cert":   carapace.ActionFiles(),
		"key":    carapace.ActionFiles(),
	})
}
