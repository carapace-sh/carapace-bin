package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install and activate a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()
	installCmd.Flags().Bool("no-rev-upgrade", false, "Do not run rev-upgrade after install")
	installCmd.Flags().Bool("unrequested", false, "Mark as unrequested (like a dependency)")
	rootCmd.AddCommand(installCmd)
}
