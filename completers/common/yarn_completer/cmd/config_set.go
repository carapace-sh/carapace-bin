package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/yarn"
	"github.com/spf13/cobra"
)

var config_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Change a configuration settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_setCmd).Standalone()

	config_setCmd.Flags().BoolP("home", "H", false, "Update the home configuration instead of the project configuration")
	config_setCmd.Flags().Bool("json", false, "Set complex configuration settings to JSON values")
	configCmd.AddCommand(config_setCmd)

	carapace.Gen(config_setCmd).PositionalCompletion(
		yarn.ActionConfigKeys(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return yarn.ActionConfigValues(c.Args[0])
		}),
	)
}
