package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_state_cache_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Drop all caches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_state_cache_clearCmd).Standalone()

	config_state_cache_clearCmd.Flags().BoolP("help", "h", false, "Print help")
	config_state_cacheCmd.AddCommand(config_state_cache_clearCmd)

	carapace.Gen(config_state_cache_clearCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("text", "json"),
	})
}
