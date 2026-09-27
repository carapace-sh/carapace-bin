package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-pnpm/pkg/actions/tools/pnpm"
	"github.com/spf13/cobra"
)

var rbCmd = &cobra.Command{
	Use:   "rb",
	Short: "Rebuild a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rbCmd).Standalone()

	rbCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rbCmd.Flags().Bool("no-pending", false, "Rebuild all matching packages, including those without pending builds")
	rbCmd.Flags().Bool("pending", false, "Rebuild packages that were not built during installation, such as under `--ignore-scripts`")
	rbCmd.Flag("no-pending").Hidden = true
	rootCmd.AddCommand(rbCmd)

	carapace.Gen(rbCmd).PositionalAnyCompletion(
		pnpm.ActionDependencyNames(),
	)
}
