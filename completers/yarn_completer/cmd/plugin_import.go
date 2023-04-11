package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/spf13/cobra"
)

var plugin_importCmd = &cobra.Command{
	Use:   "import",
	Short: "Download a plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_importCmd).Standalone()

	pluginCmd.AddCommand(plugin_importCmd)

	carapace.Gen(pluginCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if util.HasPathPrefix(c.Value) {
				return carapace.ActionFiles()
			}
			// TODO plugin completion
			return carapace.ActionValues()
		}),
	)
}
