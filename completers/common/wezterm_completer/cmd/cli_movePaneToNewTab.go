package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/wezterm"
	"github.com/spf13/cobra"
)

var cli_movePaneToNewTabCmd = &cobra.Command{
	Use:   "move-pane-to-new-tab [OPTIONS]",
	Short: "Move a pane into a new tab",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cli_movePaneToNewTabCmd).Standalone()

	cli_movePaneToNewTabCmd.Flags().BoolP("help", "h", false, "Print help")
	cli_movePaneToNewTabCmd.Flags().Bool("new-window", false, "Create tab in a new window")
	cli_movePaneToNewTabCmd.Flags().String("pane-id", "", "Specify the pane that should be moved")
	cli_movePaneToNewTabCmd.Flags().String("window-id", "", "Specify the window into which the new tab will be created")
	cli_movePaneToNewTabCmd.Flags().String("workspace", "", "Override the default workspace name with the provided name")
	cliCmd.AddCommand(cli_movePaneToNewTabCmd)

	carapace.Gen(cli_movePaneToNewTabCmd).FlagCompletion(carapace.ActionMap{
		"pane-id":   wezterm.ActionPanes(),
		"window-id": wezterm.ActionWindows(),
	})
}
