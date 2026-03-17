package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daemon_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the daemon's current status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daemon_statusCmd).Standalone()

	daemon_statusCmd.Flags().BoolP("help", "h", false, "Print help")
	daemonCmd.AddCommand(daemon_statusCmd)
}
