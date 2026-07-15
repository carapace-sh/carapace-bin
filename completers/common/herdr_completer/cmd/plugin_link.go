package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Link a local plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_linkCmd).Standalone()

	plugin_linkCmd.Flags().Bool("disabled", false, "")
	plugin_linkCmd.Flags().Bool("enabled", false, "")
	pluginCmd.AddCommand(plugin_linkCmd)

	carapace.Gen(plugin_linkCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
