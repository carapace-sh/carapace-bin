package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var horizonStatusCmd = &cobra.Command{
	Use:   "horizon:status",
	Short: "Get the current status of Horizon",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(horizonStatusCmd).Standalone()

	rootCmd.AddCommand(horizonStatusCmd)
}
