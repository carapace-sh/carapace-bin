package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:     "serve",
	Short:   "start stand-alone webserver",
	GroupID: groups[group_remote_repository_management].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serveCmd).Standalone()

	serveCmd.Flags().StringP("accesslog", "A", "", "name of access log file to write to")
	serveCmd.Flags().StringP("address", "a", "", "address to listen on (default: all interfaces)")
	serveCmd.Flags().String("certificate", "", "SSL certificate file")
	serveCmd.Flags().String("cmdserver", "", "for remote clients (ADVANCED)")
	serveCmd.Flags().BoolP("daemon", "d", false, "run server in background")
	serveCmd.Flags().String("daemon-postexec", "", "used internally by daemon mode")
	serveCmd.Flags().StringP("errorlog", "E", "", "name of error log file to write to")
	serveCmd.Flags().BoolP("ipv6", "6", false, "use IPv6 instead of IPv4")
	serveCmd.Flags().StringP("name", "n", "", "name to show in web pages (default: working directory)")
	serveCmd.Flags().String("pid-file", "", "name of file to write process ID to")
	serveCmd.Flags().StringP("port", "p", "", "port to listen on (default: 8000)")
	serveCmd.Flags().String("prefix", "", "prefix path to serve from (default: server root)")
	serveCmd.Flags().Bool("print-url", false, "start and print only the URL")
	serveCmd.Flags().Bool("stdio", false, "for remote clients (ADVANCED)")
	serveCmd.Flags().String("style", "", "template style to use")
	serveCmd.Flags().BoolP("subrepos", "S", false, "recurse into subrepositories")
	serveCmd.Flags().StringP("templates", "t", "", "web templates to use")
	serveCmd.Flags().String("web-conf", "", "name of the hgweb config file (see 'hg help hgweb')")
	serveCmd.Flags().String("webdir-conf", "", "name of the hgweb config file (DEPRECATED)")
	rootCmd.AddCommand(serveCmd)

	carapace.Gen(serveCmd).FlagCompletion(carapace.ActionMap{
		"accesslog":   carapace.ActionFiles(),
		"certificate": carapace.ActionFiles(),
		"errorlog":    carapace.ActionFiles(),
		"pid-file":    carapace.ActionFiles(),
		"web-conf":    carapace.ActionFiles(),
		"webdir-conf": carapace.ActionFiles(),
	})
}
