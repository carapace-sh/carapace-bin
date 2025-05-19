package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:                "exec",
	Short:              "Executes a command after loading the first .envrc or .env found in DIR",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(execCmd).Standalone()

	rootCmd.AddCommand(execCmd)

	carapace.Gen(execCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)

	carapace.Gen(execCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin().Shift(1),
	)

}
