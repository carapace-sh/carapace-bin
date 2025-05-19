package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/wezterm"
	"github.com/spf13/cobra"
)

var cli_setTabTitleCmd = &cobra.Command{
	Use:   "set-tab-title [OPTIONS] <TITLE>",
	Short: "Change the title of a tab",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cli_setTabTitleCmd).Standalone()

	cli_setTabTitleCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	cli_setTabTitleCmd.Flags().String("pane-id", "", "Specify the current pane")
	cli_setTabTitleCmd.Flags().String("tab-id", "", "Specify the target tab by its id")
	cliCmd.AddCommand(cli_setTabTitleCmd)

	carapace.Gen(cli_setTabTitleCmd).FlagCompletion(carapace.ActionMap{
		"pane-id": wezterm.ActionPanes(),
		"tab-id":  wezterm.ActionTabs(),
	})
}
