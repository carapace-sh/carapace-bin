package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf5_completer/cmd/action"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list [options] [<package-spec>...]",
	Short: "list packages depending on the packages' relation to the system",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().Bool("autoremove", false, "List packages which will be removed by 'dnf autoremove'")
	listCmd.Flags().Bool("available", false, "List available packages")
	listCmd.Flags().Bool("extras", false, "List extras (installed but not in any repo)")
	listCmd.Flags().Bool("installed", false, "List installed packages")
	listCmd.Flags().String("installed-from-repo", "", "Filter installed packages by repository ID")
	listCmd.Flags().Bool("json", false, "Request json output format")
	listCmd.Flags().Bool("obsoletes", false, "List packages obsoleted by packages in repos")
	listCmd.Flags().Bool("recent", false, "List packages recently added into the repositories")
	listCmd.Flags().Bool("showduplicates", false, "Show all versions of the packages")
	listCmd.Flags().Bool("upgrades", false, "List upgrades available for installed packages")

	rootCmd.AddCommand(listCmd)

	carapace.Gen(listCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(listCmd),
	)
}
