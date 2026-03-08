package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf_completer/cmd/action"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove [options] <package-spec>|@<group-spec>|@<environment-spec>...",
	Aliases: []string{"rm"},
	Short:   "remove packages",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeCmd).Standalone()

	removeCmd.Flags().String("limit", "", "override installonly_limit when using --oldinstallonly")
	removeCmd.Flags().Bool("no-autoremove", false, "disable removal of dependencies")
	removeCmd.Flags().Bool("oldinstallonly", false, "remove old installonly packages")

	rootCmd.AddCommand(removeCmd)

	carapace.Gen(removeCmd).PositionalAnyCompletion(
		action.ActionInstalledPackages(removeCmd),
	)
}
