package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "http.server",
	Short: "simple http server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("bind", "b", "", "Specify alternate bind address")
	rootCmd.Flags().Bool("cgi", false, "Run as CGI Server")
	rootCmd.Flags().StringP("directory", "d", "", "Specify alternative directory")
	rootCmd.Flags().BoolP("help", "h", false, "show this help message and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"directory": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		net.ActionPorts(),
	)
}
