package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serves a book at http://localhost:3000, and rebuilds it on changes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_serveCmd).Standalone()

	helpCmd.AddCommand(help_serveCmd)
}
