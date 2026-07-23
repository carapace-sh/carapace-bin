package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_logCmd = &cobra.Command{
	Use:     "log",
	Short:   "Inspect plugin command logs",
	Aliases: []string{"logs"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_logCmd).Standalone()

	pluginCmd.AddCommand(plugin_logCmd)
}
