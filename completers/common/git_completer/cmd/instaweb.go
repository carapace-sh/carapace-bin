package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var instawebCmd = &cobra.Command{
	Use:     "instaweb",
	Short:   "Instantly browse your working repository in gitweb",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_interrogator].ID,
}

func init() {
	carapace.Gen(instawebCmd).Standalone()

	instawebCmd.Flags().StringP("browser", "b", "", "the browser to launch")
	instawebCmd.Flags().StringP("httpd", "d", "", "the command to launch")
	instawebCmd.Flags().BoolP("local", "l", false, "only bind on 127.0.0.1")
	instawebCmd.Flags().StringP("port", "p", "", "the port to bind to")
	instawebCmd.Flags().Bool("restart", false, "restart the web server")
	instawebCmd.Flags().Bool("start", false, "start the web server")
	instawebCmd.Flags().Bool("stop", false, "stop the web server")
	rootCmd.AddCommand(instawebCmd)

	carapace.Gen(instawebCmd).FlagCompletion(carapace.ActionMap{
		"browser": bridge.ActionCarapaceBin().Split(),
		"httpd":   bridge.ActionCarapaceBin().Split(),
		"port":    net.ActionPorts(),
	})
}
