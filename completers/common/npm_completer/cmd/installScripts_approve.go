package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/npm"
	"github.com/spf13/cobra"
)

var installScripts_approveCmd = &cobra.Command{
	Use:   "approve",
	Short: "Approve install scripts for specific dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installScripts_approveCmd).Standalone()

	installScriptsCmd.AddCommand(installScripts_approveCmd)

	carapace.Gen(installScripts_approveCmd).PositionalAnyCompletion(
		npm.ActionDependencies(),
	)
}
