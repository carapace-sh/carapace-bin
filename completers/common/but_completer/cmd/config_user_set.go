package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var config_user_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a user configuration value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_user_setCmd).Standalone()

	config_user_setCmd.Flags().BoolP("global", "g", false, "Set the configuration globally instead of locally")
	config_user_setCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_userCmd.AddCommand(config_user_setCmd)

	carapace.Gen(config_user_setCmd).PositionalCompletion(
		but.ActionConfigNames(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return but.ActionConfigValues(c.Args[0])
		}),
	)
}
