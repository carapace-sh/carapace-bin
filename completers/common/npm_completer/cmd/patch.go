package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var patchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Apply local patches to installed registry dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(patchCmd).Standalone()
	patchCmd.Flags().Bool("allow-unused-patches", false, "allow patches that don't match any installed node")
	patchCmd.Flags().String("edit-dir", "", "custom directory path to use for editing a package")
	patchCmd.Flags().Bool("ignore-existing", false, "wipe the existing directory before extracting")
	patchCmd.Flags().Bool("ignore-patch-failures", false, "don't fail if a patch can't be applied")
	patchCmd.Flags().Bool("keep-edit-dir", false, "don't delete the edit directory after committing")
	patchCmd.Flags().String("patches-dir", "", "directory where patch files are stored")
	patchCmd.Flags().String("to", "", "target version for patch update")

	rootCmd.AddCommand(patchCmd)

	carapace.Gen(patchCmd).PositionalCompletion(
		carapace.ActionValues("add", "commit", "ls", "rm", "update"),
	)
}
