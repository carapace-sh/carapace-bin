package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-pnpm/pkg/actions/tools/pnpm"
	"github.com/spf13/cobra"
)

var config_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the config key to the value provided",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_setCmd).Standalone()

	config_setCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	configCmd.AddCommand(config_setCmd)

	carapace.Gen(config_setCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			global, _ := config_setCmd.Flags().GetBool("global")
			if location, _ := config_setCmd.Flags().GetString("location"); location == "global" {
				global = true
			}
			if global {
				return pnpm.ActionGlobalConfigKeys()
			}
			return pnpm.ActionLocalConfigKeys()
		}),
	)
}
