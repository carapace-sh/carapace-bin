package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bridge/pkg/actions/bridge"
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
		carapace.Batch(
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
	)

	carapace.Gen(execCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return bridge.ActionCarapaceBin(c.Args[1]).Shift(2)
		}),
	)

}
