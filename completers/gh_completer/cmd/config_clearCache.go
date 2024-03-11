package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_clearCacheCmd = &cobra.Command{
	Use:   "clear-cache",
	Short: "Clear the cli cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_clearCacheCmd).Standalone()

	configCmd.AddCommand(config_clearCacheCmd)
}
