package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install extension source and add to settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()
	installCmd.Flags().BoolP("approve", "a", false, "Trust project-local files for this command")
	installCmd.Flags().BoolP("local", "l", false, "Install project-locally (.pi/settings.json)")
	installCmd.Flags().BoolP("no-approve", "na", false, "Ignore project-local files for this command")
	rootCmd.AddCommand(installCmd)
}
