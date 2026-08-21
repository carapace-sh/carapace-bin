package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/eopkg"
	"github.com/spf13/cobra"
)

var autoremoveCmd = &cobra.Command{
	Use:     "autoremove <package1> <package2> ...",
	Aliases: []string{"rmf"},
	Short:   "remove a package along with reverse dependencies and now-unneeded automatic dependencies",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoremoveCmd).Standalone()

	autoremoveCmd.Flags().BoolP("dry-run", "n", false, "only show what would happen, do not actually perform changes")
	autoremoveCmd.Flags().Bool("ignore-comar", false, "bypass system configuration")
	autoremoveCmd.Flags().Bool("ignore-dependency", false, "do not attempt the removal/validation of reverse dependencies")
	autoremoveCmd.Flags().Bool("ignore-safety", false, "ignore safety switch on system.base component")
	autoremoveCmd.Flags().BoolP("purge", "p", false, "remove files tagged as configuration files too")

	rootCmd.AddCommand(autoremoveCmd)

	carapace.Gen(autoremoveCmd).PositionalAnyCompletion(
		eopkg.ActionPackages().FilterArgs(),
	)
}
