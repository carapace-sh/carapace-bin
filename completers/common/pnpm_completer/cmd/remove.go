package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-pnpm/pkg/actions/tools/pnpm"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Removes packages from `node_modules` and from the project's `package.json`",
	Aliases: []string{"uninstall", "rm", "un", "uni"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeCmd).Standalone()

	removeCmd.Flags().BoolP("global", "g", false, "Remove the package from the global packages directory and unlink its bins")
	removeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	removeCmd.Flags().String("lockfile-dir", "", "The directory in which `pnpm-lock.yaml` is created. Several projects may share a single lockfile")
	removeCmd.Flags().Bool("lockfile-only", false, "Dependencies are not removed from `node_modules`. Only the manifest and `pnpm-lock.yaml` are updated")
	removeCmd.Flags().BoolP("save-dev", "D", false, "Remove the dependency only from \"devDependencies\"")
	removeCmd.Flags().BoolP("save-optional", "O", false, "Remove the dependency only from \"optionalDependencies\"")
	removeCmd.Flags().BoolP("save-prod", "P", false, "Remove the dependency only from \"dependencies\"")
	rootCmd.AddCommand(removeCmd)

	carapace.Gen(removeCmd).PositionalAnyCompletion(
		pnpm.ActionDependencyNames(),
	)
}
