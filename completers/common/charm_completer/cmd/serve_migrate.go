package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serve_migrateCmd = &cobra.Command{
	Use:     "migrate",
	Short:   "Run the server migration tool.",
	Aliases: []string{"migration"},
	Hidden:  true,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serve_migrateCmd).Standalone()

	serveCmd.AddCommand(serve_migrateCmd)
}
