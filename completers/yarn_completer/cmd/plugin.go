package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pluginCmd = &cobra.Command{
	Use:     "plugin",
	Short:   "manage plugins",
	GroupID: "plugin-related",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pluginCmd).Standalone()

	rootCmd.AddCommand(pluginCmd)
}
