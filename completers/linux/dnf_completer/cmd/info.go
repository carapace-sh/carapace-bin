package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf_completer/cmd/action"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:     "info [options] [<package-spec>...]",
	Aliases: []string{"if"},
	Short:   "provide detailed information about installed or available packages",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()

	infoCmd.Flags().Bool("showduplicates", false, "show all versions of packages")
	infoCmd.Flags().Bool("installed", false, "list only installed packages")
	infoCmd.Flags().Bool("available", false, "list only available packages")
	infoCmd.Flags().Bool("extras", false, "list only extras")
	infoCmd.Flags().Bool("obsoletes", false, "list only obsoleted packages")
	infoCmd.Flags().Bool("recent", false, "list only recently added packages")
	infoCmd.Flags().Bool("upgrades", false, "list only available upgrades")
	infoCmd.Flags().Bool("autoremove", false, "list packages that will be removed by autoremove")

	rootCmd.AddCommand(infoCmd)

	carapace.Gen(infoCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(infoCmd),
	)
}
