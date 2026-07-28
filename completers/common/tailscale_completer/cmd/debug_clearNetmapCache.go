package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_clearNetmapCacheCmd = &cobra.Command{
	Use:   "clear-netmap-cache",
	Short: "Remove and discard cached network maps (if any)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_clearNetmapCacheCmd).Standalone()

	debugCmd.AddCommand(debug_clearNetmapCacheCmd)
}
