package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var indexPackCmd = &cobra.Command{
	Use:     "index-pack",
	Short:   "Build pack index file for an existing packed archive",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_manipulator].ID,
}

func init() {
	carapace.Gen(indexPackCmd).Standalone()

	indexPackCmd.Flags().Bool("check-self-contained-and-connected", false, "die if the pack contains broken links")
	indexPackCmd.Flags().Bool("fix-thin", false, "fix a \"thin\" pack produced by git pack-objects --thin")
	indexPackCmd.Flags().String("fsck-objects", "", "die if the pack contains broken objects")
	indexPackCmd.Flags().String("index-version", "", "force version for the generated pack index")
	indexPackCmd.Flags().Bool("keep", false, "create an empty .keep file before moving the index")
	indexPackCmd.Flags().String("max-input-size", "", "die, if the pack is larger than <size>")
	indexPackCmd.Flags().Bool("no-rev-index", false, "do not generate a reverse index")
	indexPackCmd.Flags().StringS("o", "o", "", "write the generated pack index into the specified file")
	indexPackCmd.Flags().String("object-format", "", "specify the given object format (hash algorithm) for the pack")
	indexPackCmd.Flags().Bool("progress-title", false, "set the title of the progress bar")
	indexPackCmd.Flags().String("promisor", "", "create a .promisor file for this pack")
	indexPackCmd.Flags().Bool("rev-index", false, "generate a reverse index")
	indexPackCmd.Flags().Bool("stdin", false, "read from stdin instead and write a copy to <pack-file>")
	indexPackCmd.Flags().String("strict", "", "die, if the pack contains broken objects or links")
	indexPackCmd.Flags().String("threads", "", "specifies the number of threads to spawn when resolving deltas")
	indexPackCmd.Flags().BoolS("v", "v", false, "be verbose about what is going on")
	rootCmd.AddCommand(indexPackCmd)

	indexPackCmd.Flag("fsck-objects").NoOptDefVal = " "
	indexPackCmd.Flag("promisor").NoOptDefVal = " "
	indexPackCmd.Flag("strict").NoOptDefVal = " "

	carapace.Gen(indexPackCmd).FlagCompletion(carapace.ActionMap{
		"o":             carapace.ActionFiles(),
		"object-format": carapace.ActionValues("sha1", "sha256"),
	})

	carapace.Gen(indexPackCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if indexPackCmd.Flag("stdin").Changed {
				return carapace.ActionValues()
			}
			return carapace.ActionFiles()
		}),
	)
}
