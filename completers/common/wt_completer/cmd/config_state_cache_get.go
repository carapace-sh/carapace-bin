package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_state_cache_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Show cache contents",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_state_cache_getCmd).Standalone()

	config_state_cache_getCmd.Flags().BoolP("help", "h", false, "Print help")
	config_state_cacheCmd.AddCommand(config_state_cache_getCmd)

	carapace.Gen(config_state_cache_getCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("text", "json"),
	})
}
