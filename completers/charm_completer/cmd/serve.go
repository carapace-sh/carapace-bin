package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:     "serve",
	Short:   "Start a self-hosted Charm Cloud server.",
	Aliases: []string{"server"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serveCmd).Standalone()

	serveCmd.Flags().String("data-dir", "", "Directory to store SQLite db, SSH keys and file data")
	serveCmd.Flags().String("health-port", "", "Health port to listen on")
	serveCmd.Flags().String("http-port", "", "HTTP port to listen on")
	serveCmd.Flags().String("ssh-port", "", "SSH port to listen on")
	serveCmd.Flags().String("stats-port", "", "Stats port to listen on")
	rootCmd.AddCommand(serveCmd)

	carapace.Gen(serveCmd).FlagCompletion(carapace.ActionMap{
		"data-dir":    carapace.ActionDirectories(),
		"health-port": net.ActionPorts(),
		"http-port":   net.ActionPorts(),
		"ssh-port":    net.ActionPorts(),
		"stats-port":  net.ActionPorts(),
	})
}
