package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/systemd-analyze_completer/cmd/action"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var unitShellCmd = &cobra.Command{
	Use:   "unit-shell",
	Short: "Run command on the namespace of the service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unitShellCmd).Standalone()
	unitShellCmd.Flags().SetInterspersed(false)

	rootCmd.AddCommand(unitShellCmd)

	carapace.Gen(unitShellCmd).PositionalCompletion(
		action.ActionServices(unitShellCmd),
	)

	carapace.Gen(unitShellCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin().Shift(1),
	)
}
