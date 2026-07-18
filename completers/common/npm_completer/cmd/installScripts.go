package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var installScriptsCmd = &cobra.Command{
	Use:   "install-scripts",
	Short: "Manage install-script approvals for dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installScriptsCmd).Standalone()
	installScriptsCmd.PersistentFlags().BoolP("all", "a", false, "act on all packages with pending install scripts")
	installScriptsCmd.PersistentFlags().Bool("allow-scripts-pin", false, "write pinned entries when approving install scripts")
	installScriptsCmd.PersistentFlags().Bool("dry-run", false, "only report what would be done")
	installScriptsCmd.PersistentFlags().Bool("json", false, "output as json")

	rootCmd.AddCommand(installScriptsCmd)
}
