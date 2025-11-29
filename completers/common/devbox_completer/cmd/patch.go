package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var patchCmd = &cobra.Command{
	Use:    "patch <store-path>",
	Short:  "Apply Devbox patches to a package to fix common linker errors",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(patchCmd).Standalone()

	patchCmd.Flags().String("gcc", "", "patch binaries to use a different gcc")
	patchCmd.Flags().String("glibc", "", "patch binaries to use a different glibc")
	patchCmd.Flags().Bool("restore-refs", false, "restore references to removed store paths")
	rootCmd.AddCommand(patchCmd)

	carapace.Gen(patchCmd).FlagCompletion(carapace.ActionMap{
		"gcc": carapace.Batch(
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
		"glibc": carapace.ActionFiles(),
	})

	carapace.Gen(patchCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
