package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var config_ui_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset (remove) a UI configuration value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_ui_unsetCmd).Standalone()

	config_ui_unsetCmd.Flags().BoolP("global", "g", false, "Unset the global configuration instead of local")
	config_ui_unsetCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_uiCmd.AddCommand(config_ui_unsetCmd)

	carapace.Gen(config_ui_unsetCmd).PositionalCompletion(
		but.ActionUIConfigNames(),
	)
}
