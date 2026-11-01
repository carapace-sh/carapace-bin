package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Run a web server for controlling Tailscale",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(webCmd).Standalone()

	webCmd.Flags().Bool("cgi", false, "run as CGI script")
	webCmd.Flags().String("listen", "", "listen address; use port 0 for automatic")
	webCmd.Flags().String("origin", "", "origin at which the web UI is served (if behind a reverse proxy or used with cgi)")
	webCmd.Flags().String("prefix", "", "URL prefix added to requests (for cgi or reverse proxies)")
	webCmd.Flags().Bool("readonly", false, "run web UI in read-only mode")
	rootCmd.AddCommand(webCmd)
}
