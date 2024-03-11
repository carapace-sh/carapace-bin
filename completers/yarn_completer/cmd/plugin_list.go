package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the available official plugins",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_listCmd).Standalone()

	plugin_listCmd.Flags().Bool("json", false, "Format the output as an NDJSON stream")
	pluginCmd.AddCommand(plugin_listCmd)
}
