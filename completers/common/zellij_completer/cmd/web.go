package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Run a web server to serve terminal sessions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(webCmd).Standalone()

	webCmd.Flags().String("cert", "", "The path to the SSL certificate (required if not listening on 127.0.0.1)")
	webCmd.Flags().Bool("create-read-only-token", false, "Create a read-only login token (can only attach to existing sessions as watcher)")
	webCmd.Flags().Bool("create-token", false, "Create a login token for the web interface, will only be displayed once and cannot later be retrieved. Returns the token name and the token")
	webCmd.Flags().BoolP("daemonize", "d", false, "Run the server in the background")
	webCmd.Flags().BoolP("help", "h", false, "Print help")
	webCmd.Flags().String("ip", "", "The ip address to listen on locally for connections (defaults to 127.0.0.1)")
	webCmd.Flags().String("key", "", "The path to the SSL key (required if not listening on 127.0.0.1)")
	webCmd.Flags().Bool("list-tokens", false, "List token names and their creation dates (cannot show actual tokens)")
	webCmd.Flags().String("port", "", "The port to listen on locally for connections (defaults to 8082)")
	webCmd.Flags().Bool("revoke-all-tokens", false, "Revoke all login tokens")
	webCmd.Flags().String("revoke-token", "", "Revoke a login token by its name")
	webCmd.Flags().String("server-startup-timeout", "", "Timeout in seconds waiting for the server to start (default: 10). Only used on Windows where the daemonized server is polled via TCP. On Unix, startup signaling uses pipes and this option is ignored")
	webCmd.Flags().Bool("start", false, "Start the server (default unless other arguments are specified)")
	webCmd.Flags().Bool("status", false, "Get the server status")
	webCmd.Flags().Bool("stop", false, "Stop the server")
	webCmd.Flags().String("timeout", "", "Timeout in seconds for the status check (default: 30)")
	webCmd.Flags().String("token-name", "", "Optional name for the token")
	rootCmd.AddCommand(webCmd)

	carapace.Gen(webCmd).FlagCompletion(carapace.ActionMap{
		"cert": carapace.ActionFiles(),
		"key":  carapace.ActionFiles(),
	})
}
