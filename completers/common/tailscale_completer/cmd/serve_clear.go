package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serve_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Remove all config for a service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serve_clearCmd).Standalone()

	serveCmd.AddCommand(serve_clearCmd)
}
