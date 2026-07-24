package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serve_advertiseCmd = &cobra.Command{
	Use:   "advertise",
	Short: "Advertise this node as a service proxy to the tailnet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serve_advertiseCmd).Standalone()

	serveCmd.AddCommand(serve_advertiseCmd)
}
