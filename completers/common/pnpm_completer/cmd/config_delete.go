package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-pnpm/pkg/actions/tools/pnpm"
	"github.com/spf13/cobra"
)

var config_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove the config key from the config file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_deleteCmd).Standalone()

	config_deleteCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	configCmd.AddCommand(config_deleteCmd)

	carapace.Gen(config_deleteCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			global, _ := config_deleteCmd.Flags().GetBool("global")
			if location, _ := config_deleteCmd.Flags().GetString("location"); location == "global" {
				global = true
			}
			if global {
				return pnpm.ActionGlobalConfigKeys()
			}
			return pnpm.ActionLocalConfigKeys()
		}),
	)
}
