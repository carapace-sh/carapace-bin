package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/common"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var splitCmd = &cobra.Command{
	Use:   "split [flags] [-- COMMAND...]",
	Short: "Create a terminal block for a shell or command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(splitCmd).Standalone()

	splitCmd.Flags().SetInterspersed(false)

	common.AddCreateFlags(splitCmd, "right")

	rootCmd.AddCommand(splitCmd)

	carapace.Gen(splitCmd).DashAnyCompletion(
		bridge.ActionCarapaceBin(),
	)
}
