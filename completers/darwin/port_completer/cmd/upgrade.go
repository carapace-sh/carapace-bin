package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade a port and its dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgradeCmd).Standalone()
	upgradeCmd.Flags().Bool("enforce-variants", false, "Upgrade even if not outdated to change variants")
	upgradeCmd.Flags().Bool("force", false, "Always consider ports outdated")
	upgradeCmd.Flags().Bool("no-replace", false, "Do not install replacement ports")
	rootCmd.AddCommand(upgradeCmd)
}
