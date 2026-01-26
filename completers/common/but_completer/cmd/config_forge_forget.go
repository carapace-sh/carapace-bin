package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/but"
	"github.com/spf13/cobra"
)

var config_forge_forgetCmd = &cobra.Command{
	Use:   "forget",
	Short: "Forget a previously authenticated forge account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_forge_forgetCmd).Standalone()

	config_forge_forgetCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_forgeCmd.AddCommand(config_forge_forgetCmd)

	carapace.Gen(config_forge_forgetCmd).PositionalCompletion(
		but.ActionUsers(),
	)
}
