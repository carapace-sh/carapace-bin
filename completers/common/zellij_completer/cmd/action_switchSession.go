package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zellij"
	"github.com/spf13/cobra"
)

var action_switchSessionCmd = &cobra.Command{
	Use:   "switch-session",
	Short: "Switch to a different session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_switchSessionCmd).Standalone()

	action_switchSessionCmd.Flags().StringP("cwd", "c", "", "Change the working directory when switching")
	action_switchSessionCmd.Flags().BoolP("help", "h", false, "Print help")
	action_switchSessionCmd.Flags().StringP("layout", "l", "", "Layout to apply when switching to the session (relative paths start at layout-dir)")
	action_switchSessionCmd.Flags().String("layout-dir", "", "Default folder to look for layouts")
	action_switchSessionCmd.Flags().String("layout-string", "", "Raw KDL layout string to use directly")
	action_switchSessionCmd.Flags().String("pane-id", "", "Optional pane ID to focus (eg. \"terminal_1\" for terminal pane with id 1, or \"plugin_2\" for plugin pane with id 2)")
	action_switchSessionCmd.Flags().String("tab-position", "", "Optional tab position to focus")
	actionCmd.AddCommand(action_switchSessionCmd)

	carapace.Gen(action_switchSessionCmd).FlagCompletion(carapace.ActionMap{
		"cwd":        carapace.ActionDirectories(),
		"layout":     carapace.ActionFiles(),
		"layout-dir": carapace.ActionDirectories(),
		"pane-id":    zellij.ActionPanes(),
	})

	carapace.Gen(action_switchSessionCmd).PositionalCompletion(
		zellij.ActionSessions(),
	)
}
