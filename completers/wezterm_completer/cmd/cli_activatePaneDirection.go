package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/wezterm"
	"github.com/spf13/cobra"
)

var cli_activatePaneDirectionCmd = &cobra.Command{
	Use:   "activate-pane-direction [OPTIONS] <DIRECTION>",
	Short: "Activate an adjacent pane in the specified direction",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cli_activatePaneDirectionCmd).Standalone()

	cli_activatePaneDirectionCmd.Flags().BoolP("help", "h", false, "Print help")
	cli_activatePaneDirectionCmd.Flags().String("pane-id", "", "Specify the current pane")
	cliCmd.AddCommand(cli_activatePaneDirectionCmd)

	carapace.Gen(cli_activatePaneDirectionCmd).FlagCompletion(carapace.ActionMap{
		"pane-id": wezterm.ActionPanes(),
	})

	carapace.Gen(cli_activatePaneDirectionCmd).PositionalCompletion(
		wezterm.ActionPaneDirections(),
	)
}
