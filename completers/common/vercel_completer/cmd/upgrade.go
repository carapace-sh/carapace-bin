package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrades the Vercel CLI to the latest version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgradeCmd).Standalone()

	upgradeCmd.Flags().Bool("disable-auto", false, "Disable automatic CLI updates")
	upgradeCmd.Flags().Bool("disable-binary", false, "Disable binary")
	upgradeCmd.Flags().Bool("dry-run", false, "Show the upgrade command without executing it")
	upgradeCmd.Flags().Bool("enable-auto", false, "Enable automatic CLI updates")
	upgradeCmd.Flags().StringP("format", "F", "", "Output format")
	upgradeCmd.Flags().Bool("json", false, "Output as JSON")

	rootCmd.AddCommand(upgradeCmd)

	carapace.Gen(upgradeCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("plain", "json"),
	})
}
