package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rereadCmd = &cobra.Command{
	Use:   "reread",
	Short: "Reload the daemon’s configuration files, without add/remove (no restarts)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rereadCmd).Standalone()

	rootCmd.AddCommand(rereadCmd)
}
