package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quiltimportCmd = &cobra.Command{
	Use:     "quiltimport",
	Short:   "Applies a quilt patchset onto the current branch",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_interaction].ID,
}

func init() {
	carapace.Gen(quiltimportCmd).Standalone()

	quiltimportCmd.Flags().String("author", "", "author name and email address for patches without any")
	quiltimportCmd.Flags().BoolP("dry-run", "n", false, "dry run")
	quiltimportCmd.Flags().Bool("keep-non-patch", false, "Pass -b to git mailinfo")
	quiltimportCmd.Flags().String("patches", "", "path to the quilt patches")
	quiltimportCmd.Flags().String("series", "", "path to the quilt series file")
	rootCmd.AddCommand(quiltimportCmd)

	carapace.Gen(quiltimportCmd).FlagCompletion(carapace.ActionMap{
		"patches": carapace.ActionDirectories(),
		"series":  carapace.ActionDirectories(),
	})
}
