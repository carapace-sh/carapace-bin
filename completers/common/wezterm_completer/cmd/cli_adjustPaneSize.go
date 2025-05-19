package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/wezterm"
	"github.com/spf13/cobra"
)

var cli_adjustPaneSizeCmd = &cobra.Command{
	Use:   "adjust-pane-size [OPTIONS] <DIRECTION>",
	Short: "Adjust the size of a pane directionally",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cli_adjustPaneSizeCmd).Standalone()

	cli_adjustPaneSizeCmd.Flags().String("amount", "", "Specify the number of cells to resize by, defaults to 1")
	cli_adjustPaneSizeCmd.Flags().BoolP("help", "h", false, "Print help")
	cli_adjustPaneSizeCmd.Flags().String("pane-id", "", "Specify the target pane")
	cliCmd.AddCommand(cli_adjustPaneSizeCmd)

	carapace.Gen(cli_adjustPaneSizeCmd).FlagCompletion(carapace.ActionMap{
		"pane-id": wezterm.ActionPanes(),
	})

	carapace.Gen(cli_adjustPaneSizeCmd).PositionalCompletion(
		wezterm.ActionPaneDirections(),
	)
}
