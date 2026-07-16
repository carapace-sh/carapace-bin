package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
	"github.com/spf13/cobra"
)

var plugin_unlinkCmd = &cobra.Command{
	Use:   "unlink <plugin_id>",
	Short: "Unlink a local plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_unlinkCmd).Standalone()

	pluginCmd.AddCommand(plugin_unlinkCmd)

	carapace.Gen(plugin_unlinkCmd).PositionalCompletion(herdr.ActionPlugins())
}
