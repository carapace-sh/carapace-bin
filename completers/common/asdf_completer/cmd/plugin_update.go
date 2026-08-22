package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a plugin to latest commit on default branch or a particular git-ref",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_updateCmd).Standalone()

	plugin_updateCmd.Flags().Bool("all", false, "Update all installed plugins")
	pluginCmd.AddCommand(plugin_updateCmd)
}
