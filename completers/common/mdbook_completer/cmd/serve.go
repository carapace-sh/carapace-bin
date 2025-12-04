package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serves a book at http://localhost:3000, and rebuilds it on changes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serveCmd).Standalone()

	serveCmd.Flags().StringP("dest-dir", "d", "", "Output directory for the book")
	serveCmd.Flags().BoolP("help", "h", false, "Print help")
	serveCmd.Flags().StringP("hostname", "n", "", "Hostname to listen on for HTTP connections")
	serveCmd.Flags().BoolP("open", "o", false, "Opens the compiled book in a web browser")
	serveCmd.Flags().StringP("port", "p", "", "Port to use for HTTP connections")
	serveCmd.Flags().BoolP("version", "V", false, "Print version")
	serveCmd.Flags().String("watcher", "", "The filesystem watching technique")
	rootCmd.AddCommand(serveCmd)

	carapace.Gen(serveCmd).FlagCompletion(carapace.ActionMap{
		"dest-dir": carapace.ActionDirectories(),
		"hostname": carapace.ActionValues("localhost", "0.0.0.0"),
		"port":     net.ActionPorts(),
	})

	carapace.Gen(serveCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
