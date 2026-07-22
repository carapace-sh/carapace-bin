package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/npm"
	"github.com/spf13/cobra"
)

var installScripts_denyCmd = &cobra.Command{
	Use:   "deny",
	Short: "Deny install scripts for specific dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installScripts_denyCmd).Standalone()

	installScriptsCmd.AddCommand(installScripts_denyCmd)

	carapace.Gen(installScripts_denyCmd).PositionalAnyCompletion(
		npm.ActionDependencies(),
	)
}
