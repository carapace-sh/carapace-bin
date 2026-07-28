package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_forcePreferDerpCmd = &cobra.Command{
	Use:   "force-prefer-derp",
	Short: "Prefer the given region ID if reachable (until restart, or 0 to clear)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_forcePreferDerpCmd).Standalone()

	debug_forcePreferDerpCmd.Flags().Int("region", 0, "DERP region ID to prefer")
	debugCmd.AddCommand(debug_forcePreferDerpCmd)
}
