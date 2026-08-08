package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_state_cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Regenerable caches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_state_cacheCmd).Standalone()

	config_state_cacheCmd.PersistentFlags().String("format", "text", "Output format (text, json) [default: text]")
	config_state_cacheCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_stateCmd.AddCommand(config_state_cacheCmd)

	carapace.Gen(config_state_cacheCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("text", "json"),
	})
}
