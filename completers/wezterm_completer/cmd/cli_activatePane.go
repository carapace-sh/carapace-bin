package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/wezterm"
	"github.com/spf13/cobra"
)

var cli_activatePaneCmd = &cobra.Command{
	Use:   "activate-pane [OPTIONS]",
	Short: "Activate (focus) a pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cli_activatePaneCmd).Standalone()

	cli_activatePaneCmd.Flags().BoolP("help", "h", false, "Print help")
	cli_activatePaneCmd.Flags().String("pane-id", "", "Specify the target pane")
	cliCmd.AddCommand(cli_activatePaneCmd)

	carapace.Gen(cli_activatePaneCmd).FlagCompletion(carapace.ActionMap{
		"pane-id": wezterm.ActionPanes(),
	})
}
