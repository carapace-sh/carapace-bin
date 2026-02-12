package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync with the configured server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(syncCmd).Standalone()

	syncCmd.Flags().BoolP("force", "f", false, "Force re-download everything")
	syncCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(syncCmd)
}
