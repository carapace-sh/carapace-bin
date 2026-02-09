package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var config_ui_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a UI configuration value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_ui_setCmd).Standalone()

	config_ui_setCmd.Flags().BoolP("global", "g", false, "Set the configuration globally instead of locally")
	config_ui_setCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_uiCmd.AddCommand(config_ui_setCmd)

	carapace.Gen(config_ui_setCmd).PositionalCompletion(
		but.ActionUIConfigNames(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return but.ActionUIConfigValues(c.Args[0])
		}),
	)
}
