package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listUpgradesCmd = &cobra.Command{
	Use:     "list-upgrades",
	Short:   "List packages to be upgraded",
	Aliases: []string{"lu"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listUpgradesCmd).Standalone()

	listUpgradesCmd.Flags().Bool("ignore-build-no", false, "Do not take build no into account")
	listUpgradesCmd.Flags().BoolP("long", "l", false, "Show in long format")
	listUpgradesCmd.Flags().StringP("component", "c", "", "List upgrades under given component")
	listUpgradesCmd.Flags().BoolP("install-info", "i", false, "Show detailed install info")

	rootCmd.AddCommand(listUpgradesCmd)
}
