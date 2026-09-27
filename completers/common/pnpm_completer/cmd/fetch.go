package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/npm"
	"github.com/spf13/cobra"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch packages from the lockfile into the virtual store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fetchCmd).Standalone()

	fetchCmd.Flags().BoolP("dev", "D", false, "")
	fetchCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	fetchCmd.Flags().Bool("ignore-pnpmfile", false, "Disable pnpm hooks defined in `.pnpmfile.cjs`, including the pnpmfiles of config dependencies")
	fetchCmd.Flags().BoolP("prod", "P", false, "")
	fetchCmd.Flags().Bool("production", false, "")
	rootCmd.AddCommand(fetchCmd)

	carapace.Gen(fetchCmd).PositionalAnyCompletion(
		npm.ActionPackageSearch(""),
	)
}
