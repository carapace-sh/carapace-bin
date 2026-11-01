package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_forceNetmapUpdateCmd = &cobra.Command{
	Use:   "force-netmap-update",
	Short: "Force a full no-op netmap update (for load testing)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_forceNetmapUpdateCmd).Standalone()

	debugCmd.AddCommand(debug_forceNetmapUpdateCmd)
}
