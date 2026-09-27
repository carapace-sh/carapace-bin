package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-pnpm/pkg/actions/tools/pnpm"
	"github.com/spf13/cobra"
)

var config_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Print the config value for the provided key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getCmd).Standalone()

	config_getCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	configCmd.AddCommand(config_getCmd)

	carapace.Gen(config_getCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			global, _ := config_getCmd.Flags().GetBool("global")
			if location, _ := config_getCmd.Flags().GetString("location"); location == "global" {
				global = true
			}
			if global {
				return pnpm.ActionGlobalConfigKeys()
			}
			return pnpm.ActionLocalConfigKeys()
		}),
	)
}
