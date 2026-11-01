package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/eopkg"
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

	listInstalledCmd.Flags().BoolP("automatic", "a", false, "show automatically installed packages and the parent dependency")
	listInstalledCmd.Flags().StringP("component", "c", "", "list installed packages under given component")
	listInstalledCmd.Flags().BoolP("explicit", "e", false, "show installed packages that were installed by a user")
	listInstalledCmd.Flags().BoolP("install-info", "i", false, "show detailed install info")
	listInstalledCmd.Flags().BoolP("long", "l", false, "show in long format")
	listInstalledCmd.Flags().StringP("with-build-host", "b", "", "only list the installed packages built by the given host")

	rootCmd.AddCommand(listInstalledCmd)

	carapace.Gen(listInstalledCmd).FlagCompletion(carapace.ActionMap{
		"component": eopkg.ActionComponents(),
	})
}
