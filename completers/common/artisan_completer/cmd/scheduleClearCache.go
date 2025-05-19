package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scheduleClearCacheCmd = &cobra.Command{
	Use:   "schedule:clear-cache",
	Short: "Delete the cached mutex files created by scheduler",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scheduleClearCacheCmd).Standalone()

	rootCmd.AddCommand(scheduleClearCacheCmd)
}
