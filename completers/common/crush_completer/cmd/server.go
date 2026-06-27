package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the Crush server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverCmd).Standalone()

	serverCmd.Flags().StringP("host", "H", "", "Server host (TCP or Unix socket)")
	rootCmd.AddCommand(serverCmd)

	carapace.Gen(serverCmd).FlagCompletion(carapace.ActionMap{
		"host": net.ActionHosts(),
	})
}
