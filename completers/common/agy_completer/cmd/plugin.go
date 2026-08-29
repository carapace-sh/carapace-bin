package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pluginCmd = &cobra.Command{
	Use:     "plugin",
	GroupID: "integration",
	Short:   "Manage plugins (install, uninstall, list, enable, disable)",
	Run:     func(cmd *cobra.Command, args []string) {},
	Aliases: []string{"plugins"},
}

func init() {
	carapace.Gen(pluginCmd).Standalone()
	rootCmd.AddCommand(pluginCmd)
}
