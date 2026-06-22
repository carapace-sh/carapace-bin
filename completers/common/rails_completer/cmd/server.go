package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rails_completer/cmd/common"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the Rails server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverCmd).Standalone()

	serverCmd.Flags().StringP("binding", "b", "localhost", "IP to bind to")
	serverCmd.Flags().StringP("config", "c", "config.ru", "Rackup config file")
	serverCmd.Flags().BoolP("daemon", "d", false, "Run as daemon")
	serverCmd.Flags().BoolP("dev-caching", "C", false, "Enable development caching")
	serverCmd.Flags().Bool("early-hints", false, "Enable HTTP/2 early hints")
	serverCmd.Flags().Bool("log-to-stdout", false, "Log to stdout")
	serverCmd.Flags().StringP("pid", "P", "tmp/pids/server.pid", "PID file path")
	serverCmd.Flags().IntP("port", "p", 3000, "Port to run on")
	serverCmd.Flags().StringP("using", "u", "", "Rack server handler")

	common.AddEnvironmentFlag(serverCmd, "development")

	carapace.Gen(serverCmd).FlagCompletion(carapace.ActionMap{
		"using": carapace.ActionValues("puma", "thin", "webrick"),
	})
}
