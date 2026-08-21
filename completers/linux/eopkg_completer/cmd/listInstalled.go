package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listInstalledCmd = &cobra.Command{
	Use:     "list-installed",
	Aliases: []string{"li"},
	Short:   "show a list of all installed packages",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listInstalledCmd).Standalone()

	listInstalledCmd.Flags().BoolP("automatic", "a", false, "show packages that have been automatically installed as a dependency")
	listInstalledCmd.Flags().StringP("build-host", "b", "", "only show packages that come from a particular build host")
	listInstalledCmd.Flags().StringP("component", "c", "", "only show installed packages from the specified component")
	listInstalledCmd.Flags().BoolP("install-info", "i", false, "show detailed installation information for each package")
	listInstalledCmd.Flags().BoolP("long", "l", false, "show full details of each package instead of one line summaries")

	rootCmd.AddCommand(listInstalledCmd)
}
