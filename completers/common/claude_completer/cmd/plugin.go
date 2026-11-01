package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pluginCmd = &cobra.Command{
	Use:     "plugin",
	Short:   "Manage Claude Code plugins",
	Aliases: []string{"plugins"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pluginCmd).Standalone()

	pluginCmd.Flags().BoolP("help", "h", false, "Display help for command")
	rootCmd.AddCommand(pluginCmd)
}
