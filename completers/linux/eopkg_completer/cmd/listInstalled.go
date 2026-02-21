package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listInstalledCmd = &cobra.Command{
	Use:     "list-installed",
	Short:   "List installed packages",
	Aliases: []string{"li"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listInstalledCmd).Standalone()

	listInstalledCmd.Flags().Bool("without-buildno", false, "Show packages without build number")
	listInstalledCmd.Flags().BoolP("long", "l", false, "Show in long format")
	listInstalledCmd.Flags().StringP("component", "c", "", "List installed packages under given component")
	listInstalledCmd.Flags().BoolP("install-info", "i", false, "Show detailed install info")

	rootCmd.AddCommand(listInstalledCmd)
}
