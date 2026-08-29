package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pluginListCmd = &cobra.Command{
	Use:   "list",
	Short: "List imported plugins",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pluginListCmd).Standalone()
	pluginCmd.AddCommand(pluginListCmd)
}
