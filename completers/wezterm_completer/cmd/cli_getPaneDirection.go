package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/wezterm"
	"github.com/spf13/cobra"
)

var cli_getPaneDirectionCmd = &cobra.Command{
	Use:   "get-pane-direction [OPTIONS] <DIRECTION>",
	Short: "Determine the adjacent pane in the specified direction",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cli_getPaneDirectionCmd).Standalone()

	cli_getPaneDirectionCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	cli_getPaneDirectionCmd.Flags().String("pane-id", "", "Specify the current pane")
	cliCmd.AddCommand(cli_getPaneDirectionCmd)

	carapace.Gen(cli_getPaneDirectionCmd).FlagCompletion(carapace.ActionMap{
		"pane-id": wezterm.ActionPanes(),
	})

	carapace.Gen(cli_getPaneDirectionCmd).PositionalCompletion(
		wezterm.ActionPaneDirections(),
	)
}
