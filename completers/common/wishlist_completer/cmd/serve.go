package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:     "serve",
	Short:   "Serve the TUI over SSH.",
	Aliases: []string{"server", "s"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serveCmd).Standalone()
	rootCmd.AddCommand(serveCmd)
}
