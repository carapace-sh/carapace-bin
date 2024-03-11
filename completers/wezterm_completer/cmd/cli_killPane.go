package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/wezterm"
	"github.com/spf13/cobra"
)

var cli_killPaneCmd = &cobra.Command{
	Use:   "kill-pane [OPTIONS]",
	Short: "Kill a pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cli_killPaneCmd).Standalone()

	cli_killPaneCmd.Flags().BoolP("help", "h", false, "Print help")
	cli_killPaneCmd.Flags().String("pane-id", "", "Specify the target pane")
	cliCmd.AddCommand(cli_killPaneCmd)

	carapace.Gen(cli_killPaneCmd).FlagCompletion(carapace.ActionMap{
		"pane-id": wezterm.ActionPanes(),
	})
}
