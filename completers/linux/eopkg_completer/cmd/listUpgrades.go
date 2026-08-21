package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listUpgradesCmd = &cobra.Command{
	Use:     "list-upgrades",
	Aliases: []string{"lu"},
	Short:   "list all package upgrades that are currently available",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listUpgradesCmd).Standalone()

	listUpgradesCmd.Flags().StringP("component", "c", "", "only show upgrades from the given component")
	listUpgradesCmd.Flags().BoolP("install-info", "i", false, "show detailed installation information on each available upgrade")
	listUpgradesCmd.Flags().BoolP("long", "l", false, "show detailed information on each package to be updated")

	rootCmd.AddCommand(listUpgradesCmd)
}
