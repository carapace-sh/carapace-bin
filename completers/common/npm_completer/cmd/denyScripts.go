package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/npm"
	"github.com/spf13/cobra"
)

var denyScriptsCmd = &cobra.Command{
	Use:   "deny-scripts",
	Short: "Deny install scripts for specific dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(denyScriptsCmd).Standalone()
	denyScriptsCmd.Flags().BoolP("all", "a", false, "act on all packages with pending install scripts")
	denyScriptsCmd.Flags().Bool("allow-scripts-pin", false, "write pinned entries when denying install scripts")
	denyScriptsCmd.Flags().Bool("json", false, "output as json")

	rootCmd.AddCommand(denyScriptsCmd)

	carapace.Gen(denyScriptsCmd).PositionalAnyCompletion(
		npm.ActionDependencies(),
	)
}
