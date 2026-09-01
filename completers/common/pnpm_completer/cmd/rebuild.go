package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-pnpm/pkg/actions/tools/pnpm"
	"github.com/spf13/cobra"
)

var rebuildCmd = &cobra.Command{
	Use:     "rebuild",
	Short:   "Rebuild a package",
	Aliases: []string{"rb"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rebuildCmd).Standalone()

	rebuildCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rebuildCmd.Flags().Bool("no-pending", false, "Rebuild all matching packages, including those without pending builds")
	rebuildCmd.Flags().Bool("pending", false, "Rebuild packages that were not built during installation, such as under `--ignore-scripts`")
	rebuildCmd.Flag("no-pending").Hidden = true
	rootCmd.AddCommand(rebuildCmd)

	carapace.Gen(rebuildCmd).PositionalAnyCompletion(
		pnpm.ActionDependencyNames(),
	)
}
