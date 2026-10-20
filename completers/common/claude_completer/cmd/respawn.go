package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var respawnCmd = &cobra.Command{
	Use:   "respawn",
	Short: "Restart a background session (or all of them) so it picks up the current Claude binary",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(respawnCmd).Standalone()

	rootCmd.AddCommand(respawnCmd)
}
