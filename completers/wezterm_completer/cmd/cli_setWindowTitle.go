package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/wezterm"
	"github.com/spf13/cobra"
)

var cli_setWindowTitleCmd = &cobra.Command{
	Use:   "set-window-title [OPTIONS] <TITLE>",
	Short: "Change the title of a window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cli_setWindowTitleCmd).Standalone()

	cli_setWindowTitleCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	cli_setWindowTitleCmd.Flags().String("pane-id", "", "Specify the current pane")
	cli_setWindowTitleCmd.Flags().String("window-id", "", "Specify the target window by its id")
	cliCmd.AddCommand(cli_setWindowTitleCmd)

	carapace.Gen(cli_setWindowTitleCmd).FlagCompletion(carapace.ActionMap{
		"pane-id":   wezterm.ActionPanes(),
		"window-id": wezterm.ActionWindows(),
	})
}
