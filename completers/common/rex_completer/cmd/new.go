package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/common"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new [LABEL] [-- COMMAND...]",
	Short: "Create a session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(newCmd).Standalone()

	newCmd.Flags().SetInterspersed(false)

	common.AddSessionCreateFlags(newCmd)

	rootCmd.AddCommand(newCmd)

	carapace.Gen(newCmd).DashAnyCompletion(
		bridge.ActionCarapaceBin(),
	)
}
