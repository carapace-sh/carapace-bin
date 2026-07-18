package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/npm"
	"github.com/spf13/cobra"
)

var approveScriptsCmd = &cobra.Command{
	Use:   "approve-scripts",
	Short: "Approve install scripts for specific dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(approveScriptsCmd).Standalone()
	approveScriptsCmd.Flags().BoolP("all", "a", false, "act on all packages with pending install scripts")
	approveScriptsCmd.Flags().Bool("allow-scripts-pending", false, "list packages with install scripts not yet covered by the allowScripts policy")
	approveScriptsCmd.Flags().Bool("allow-scripts-pin", false, "write pinned entries when approving install scripts")
	approveScriptsCmd.Flags().Bool("json", false, "output as json")

	rootCmd.AddCommand(approveScriptsCmd)

	carapace.Gen(approveScriptsCmd).PositionalAnyCompletion(
		npm.ActionDependencies(),
	)
}
