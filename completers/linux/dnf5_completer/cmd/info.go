package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf5_completer/cmd/action"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info [options] [<package-spec>...]",
	Short: "list packages with additional details",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()

	infoCmd.Flags().Bool("autoremove", false, "List packages which will be removed by 'dnf autoremove'")
	infoCmd.Flags().Bool("available", false, "List available packages")
	infoCmd.Flags().Bool("extras", false, "List extras (installed but not in any repo)")
	infoCmd.Flags().Bool("installed", false, "List installed packages")
	infoCmd.Flags().String("installed-from-repo", "", "Filter installed packages by repository ID")
	infoCmd.Flags().Bool("json", false, "Request json output format")
	infoCmd.Flags().Bool("obsoletes", false, "List packages obsoleted by packages in repos")
	infoCmd.Flags().Bool("recent", false, "List packages recently added into the repositories")
	infoCmd.Flags().Bool("showduplicates", false, "Show all versions of the packages")
	infoCmd.Flags().Bool("upgrades", false, "List upgrades available for installed packages")

	rootCmd.AddCommand(infoCmd)

	carapace.Gen(infoCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(infoCmd),
	)
}
