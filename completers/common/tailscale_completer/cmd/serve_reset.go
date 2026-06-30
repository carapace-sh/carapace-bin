package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serve_resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset current serve config",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serve_resetCmd).Standalone()

	serveCmd.AddCommand(serve_resetCmd)
}
