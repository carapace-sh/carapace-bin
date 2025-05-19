package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/wezterm"
	"github.com/spf13/cobra"
)

var cli_activateTabCmd = &cobra.Command{
	Use:   "activate-tab [OPTIONS]",
	Short: "Activate a tab",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cli_activateTabCmd).Standalone()

	cli_activateTabCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	cli_activateTabCmd.Flags().Bool("no-wrap", false, "Prevents wrapping around")
	cli_activateTabCmd.Flags().String("pane-id", "", "Specify the current pane")
	cli_activateTabCmd.Flags().String("tab-id", "", "Specify the target tab by its id")
	cli_activateTabCmd.Flags().String("tab-index", "", "Specify the target tab by its index")
	cli_activateTabCmd.Flags().String("tab-relative", "", "Specify the target tab by its relative offset")
	cliCmd.AddCommand(cli_activateTabCmd)

	carapace.Gen(cli_activateTabCmd).FlagCompletion(carapace.ActionMap{
		"pane-id": wezterm.ActionPanes(),
		"tab-id":  wezterm.ActionTabs(),
		// TODO more flags
	})
}
