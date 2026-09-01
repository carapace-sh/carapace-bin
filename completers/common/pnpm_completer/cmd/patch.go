package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-pnpm/pkg/actions/tools/pnpm"
	"github.com/spf13/cobra"
)

var patchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Prepare a package for patching",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(patchCmd).Standalone()

	patchCmd.Flags().StringP("edit-dir", "d", "", "The package that needs to be modified will be extracted to this directory")
	patchCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	patchCmd.Flags().Bool("ignore-existing", false, "Ignore existing patch files when patching")

	carapace.Gen(patchCmd).FlagCompletion(carapace.ActionMap{
		"edit-dir": carapace.ActionDirectories(),
	})

	carapace.Gen(patchCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return pnpm.ActionPackageSearch("")
		}),
	)

	rootCmd.AddCommand(patchCmd)
}
