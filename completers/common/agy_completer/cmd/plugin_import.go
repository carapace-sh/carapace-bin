package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pluginImportCmd = &cobra.Command{
	Use:   "import",
	Short: "Import plugins from gemini or claude",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pluginImportCmd).Standalone()
	pluginCmd.AddCommand(pluginImportCmd)

	carapace.Gen(pluginImportCmd).PositionalCompletion(
		carapace.ActionValues("gemini", "claude"),
	)
}
