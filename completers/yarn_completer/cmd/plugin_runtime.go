package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_runtimeCmd = &cobra.Command{
	Use:   "runtime",
	Short: "List the active plugins",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_runtimeCmd).Standalone()

	plugin_runtimeCmd.Flags().Bool("json", false, "Format the output as an NDJSON stream")
	pluginCmd.AddCommand(plugin_runtimeCmd)
}
