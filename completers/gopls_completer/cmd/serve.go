package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "run a server for Go code using the Language Server Protocol",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serveCmd).Standalone()
	addServerFlags(serveCmd)
	rootCmd.AddCommand(serveCmd)

	carapace.Gen(serveCmd).PreRun(func(cmd *cobra.Command, args []string) {
		cmd.Flags().VisitAll(func(f *pflag.Flag) {
			if flag := rootCmd.Flag(f.Name); flag != nil && flag.Changed {
				f.Hidden = true
			}
		})
	})
}

func addServerFlags(cmd *cobra.Command) {
	cmd.Flags().StringS("debug", "debug", "", "serve debug information on the supplied address")
	cmd.Flags().StringS("listen", "listen", "", "address on which to listen for remote connections. If prefixed by 'unix;', the subsequent address is assumed to be a unix domain socket. Otherwise, TCP is used.")
	cmd.Flags().StringS("listen.timeout", "listen.timeout", "", "when used with -listen, shut down the server when there are no connected clients for this duration")
	cmd.Flags().StringS("logfile", "logfile", "", "filename to log to. if value is \"auto\", then logging to a default output file is enabled")
	cmd.Flags().StringS("mode", "mode", "", "no effect")
	cmd.Flags().StringS("port", "port", "", "port on which to run gopls for debugging purposes")
	cmd.Flags().StringS("remote.debug", "remote.debug", "", "when used with -remote=auto, the -debug value used to start the daemon")
	cmd.Flags().StringS("remote.listen.timeout", "remote.listen.timeout", "", "when used with -remote=auto, the -listen.timeout value used to start the daemon (default 1m0s)")
	cmd.Flags().StringS("remote.logfile", "remote.logfile", "", "when used with -remote=auto, the -logfile value used to start the daemon")
	cmd.Flags().BoolS("rpc.trace", "rpc.trace", false, "print the full rpc trace in lsp inspector format")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"logfile":        carapace.ActionFiles(),
		"port":           net.ActionPorts(),
		"remote.logfile": carapace.ActionFiles(),
	})
}
