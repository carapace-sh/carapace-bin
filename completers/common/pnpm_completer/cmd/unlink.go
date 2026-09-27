package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/npm"
	"github.com/spf13/cobra"
)

var unlinkCmd = &cobra.Command{
	Use:     "unlink",
	Short:   "Removes links to a local package and reinstalls it",
	Aliases: []string{"dislink"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unlinkCmd).Standalone()

	unlinkCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	unlinkCmd.Flags().Bool("ignore-pnpmfile", false, "Disable pnpm hooks defined in `.pnpmfile.cjs`, including the pnpmfiles of config dependencies")
	rootCmd.AddCommand(unlinkCmd)

	carapace.Gen(unlinkCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionDirectories(),
			npm.ActionPackageSearch(""),
		).ToA(),
	)
}
