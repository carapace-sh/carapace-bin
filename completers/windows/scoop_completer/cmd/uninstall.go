package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "uninstall an app",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uninstallCmd).Standalone()
	uninstallCmd.Flags().BoolP("global", "g", false, "uninstall a globally installed app")
	uninstallCmd.Flags().BoolP("purge", "p", false, "remove all persistent data")
	rootCmd.AddCommand(uninstallCmd)
}
