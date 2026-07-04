package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Deactivate and uninstall a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uninstallCmd).Standalone()
	uninstallCmd.Flags().Bool("follow-dependencies", false, "Also uninstall unneeded dependency ports")
	uninstallCmd.Flags().Bool("follow-dependents", false, "Recursively uninstall dependent ports first")
	uninstallCmd.Flags().Bool("no-exec", false, "Do not run uninstall hooks")
	rootCmd.AddCommand(uninstallCmd)
}
