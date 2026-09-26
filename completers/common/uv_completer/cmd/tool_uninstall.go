package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/uv"
	"github.com/spf13/cobra"
)

var tool_uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall a tool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tool_uninstallCmd).Standalone()

	tool_uninstallCmd.Flags().Bool("all", false, "Uninstall all tools")
	toolCmd.AddCommand(tool_uninstallCmd)
	carapace.Gen(tool_uninstallCmd).PositionalAnyCompletion(uv.ActionTools())
}
