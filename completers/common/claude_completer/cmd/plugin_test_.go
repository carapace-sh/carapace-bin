package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run a mod's tests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_testCmd).Standalone()

	plugin_testCmd.Flags().BoolP("help", "h", false, "Display help for command")
	pluginCmd.AddCommand(plugin_testCmd)
}
