package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_pane_openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open a plugin pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_pane_openCmd).Standalone()

	plugin_pane_openCmd.Flags().String("cwd", "", "")
	plugin_pane_openCmd.Flags().String("direction", "", "")
	plugin_pane_openCmd.Flags().String("entrypoint", "", "")
	plugin_pane_openCmd.Flags().StringSlice("env", nil, "Set an environment variable for the launched process")
	plugin_pane_openCmd.Flags().Bool("focus", false, "")
	plugin_pane_openCmd.Flags().Bool("no-focus", false, "")
	plugin_pane_openCmd.Flags().String("placement", "", "")
	plugin_pane_openCmd.Flags().String("plugin", "", "")
	plugin_pane_openCmd.Flags().String("target-pane", "", "")
	plugin_pane_openCmd.Flags().String("workspace", "", "")
	plugin_paneCmd.AddCommand(plugin_pane_openCmd)

	carapace.Gen(plugin_pane_openCmd).FlagCompletion(carapace.ActionMap{
		"cwd":       carapace.ActionFiles(),
		"direction": carapace.ActionValues("right", "down"),
		"placement": carapace.ActionValues("overlay", "split", "tab", "zoomed"),
	})
}
