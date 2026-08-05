package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var config_alias_dryRunCmd = &cobra.Command{
	Use:   "dry-run",
	Short: "Preview an alias invocation with template expansion",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_alias_dryRunCmd).Standalone()

	config_alias_dryRunCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_aliasCmd.AddCommand(config_alias_dryRunCmd)

	carapace.Gen(config_alias_dryRunCmd).PositionalCompletion(
		carapace.ActionValues(),
	)

	carapace.Gen(config_alias_dryRunCmd).DashAnyCompletion(
		bridge.ActionCarapaceBin(),
	)
}
