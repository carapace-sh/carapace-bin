package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serve_drainCmd = &cobra.Command{
	Use:   "drain",
	Short: "Drain a service from the current node",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serve_drainCmd).Standalone()

	serveCmd.AddCommand(serve_drainCmd)
}
