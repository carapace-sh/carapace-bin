package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_reindexCmd = &cobra.Command{
	Use:   "reindex",
	Short: "Rebuild commit index",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_reindexCmd).Standalone()

	debug_reindexCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	debugCmd.AddCommand(debug_reindexCmd)
}
