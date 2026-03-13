package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listAvailableCmd = &cobra.Command{
	Use:     "list-available",
	Short:   "List available packages in the repositories",
	Aliases: []string{"la"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listAvailableCmd).Standalone()

	listAvailableCmd.Flags().BoolP("long", "l", false, "Show in long format")
	listAvailableCmd.Flags().StringP("component", "c", "", "List available packages under given component")
	listAvailableCmd.Flags().BoolP("uninstalled", "U", false, "Show only uninstalled packages")

	rootCmd.AddCommand(listAvailableCmd)
}
