package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install previously gotten apps from the App Store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()
	installCmd.Flags().Bool("bundle", false, "Process all app IDs as bundle IDs")
	installCmd.Flags().Bool("force", false, "Force reinstall")
	rootCmd.AddCommand(installCmd)
}
