package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/common"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [flags] -- COMMAND...",
	Short: "Run a command in a new terminal block",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().SetInterspersed(false)

	common.AddCreateFlags(runCmd, "")

	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).DashAnyCompletion(
		bridge.ActionCarapaceBin(),
	)
}
