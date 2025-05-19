package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/yarn"
	"github.com/spf13/cobra"
)

var config_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset a configuration setting",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_unsetCmd).Standalone()

	config_unsetCmd.Flags().BoolP("home", "H", false, "Update the home configuration instead of the project configuration")
	configCmd.AddCommand(config_unsetCmd)

	carapace.Gen(config_unsetCmd).PositionalCompletion(
		yarn.ActionConfigKeys(),
	)
}
