package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf_completer/cmd/action"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list [options] [<package-spec>...]",
	Aliases: []string{"ls"},
	Short:   "list installed or available packages",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().Bool("autoremove", false, "list packages that will be removed by autoremove")
	listCmd.Flags().Bool("available", false, "list only available packages")
	listCmd.Flags().Bool("extras", false, "list only extras")
	listCmd.Flags().Bool("installed", false, "list only installed packages")
	listCmd.Flags().Bool("json", false, "request JSON output format")
	listCmd.Flags().Bool("obsoletes", false, "list only obsoleted packages")
	listCmd.Flags().Bool("recent", false, "list only recently added packages")
	listCmd.Flags().Bool("showduplicates", false, "show all versions of packages")
	listCmd.Flags().Bool("upgrades", false, "list only available upgrades")

	rootCmd.AddCommand(listCmd)

	carapace.Gen(listCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(listCmd),
	)
}
