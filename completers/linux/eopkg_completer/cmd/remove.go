package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/eopkg"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove <package1> <package2> ...",
	Aliases: []string{"rm"},
	Short:   "remove packages from the system",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeCmd).Standalone()

	removeCmd.Flags().StringP("component", "c", "", "remove any packages under the given component, and any child component")
	removeCmd.Flags().BoolP("dry-run", "n", false, "only show what would happen, do not actually perform changes")
	removeCmd.Flags().Bool("ignore-comar", false, "bypass system configuration")
	removeCmd.Flags().Bool("ignore-dependency", false, "do not take dependency information into account")
	removeCmd.Flags().Bool("ignore-safety", false, "ignore safety switch on system.base component")
	removeCmd.Flags().Bool("purge", false, "removes everything including changed config files of the package")

	rootCmd.AddCommand(removeCmd)

	carapace.Gen(removeCmd).FlagCompletion(carapace.ActionMap{
		"component": eopkg.ActionComponents(),
	})

	carapace.Gen(removeCmd).PositionalAnyCompletion(
		eopkg.ActionPackages().FilterArgs(),
	)
}
