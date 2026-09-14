package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/nix"
	"github.com/spf13/cobra"
)

var config_showCmd = &cobra.Command{
	Use:   "show [name]",
	Short: "show the Nix configuration or the value of a specific setting",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_showCmd).Standalone()

	common.AddJSONFlags(config_showCmd)
	common.AddLoggingFlags(config_showCmd)

	configCmd.AddCommand(config_showCmd)

	carapace.Gen(config_showCmd).PositionalCompletion(
		nix.ActionConfigKeys(),
	)
}
