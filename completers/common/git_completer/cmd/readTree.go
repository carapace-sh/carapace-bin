package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var readTreeCmd = &cobra.Command{
	Use:     "read-tree",
	Short:   "Reads tree information into the index",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_manipulator].ID,
}

func init() {
	carapace.Gen(readTreeCmd).Standalone()

	readTreeCmd.Flags().Bool("aggressive", false, "resolve more merge cases internally")
	readTreeCmd.Flags().BoolP("dry-run", "n", false, "check if the command would error out")
	readTreeCmd.Flags().Bool("empty", false, "instead of reading tree objects into the index, just empty it")
	readTreeCmd.Flags().BoolS("i", "i", false, "disable index check")
	readTreeCmd.Flags().String("index-output", "", "write the resulting index in the named file")
	readTreeCmd.Flags().BoolS("m", "m", false, "perform a merge, not just a read")
	readTreeCmd.Flags().Bool("no-recurse-submodules", false, "do not recurse submodules")
	readTreeCmd.Flags().Bool("no-sparse-checkout", false, "disable sparse checkout support even if core.sparseCheckout is true")
	readTreeCmd.Flags().String("prefix", "", "read the contents of the named tree-ish under the directory at <prefix")
	readTreeCmd.Flags().BoolP("quiet", "q", false, "suppress feedback messages")
	readTreeCmd.Flags().Bool("recurse-submodules", false, "recurse submodules")
	readTreeCmd.Flags().Bool("reset", false, "same as -m, except that unmerged entries are discarded instead of failing")
	readTreeCmd.Flags().Bool("trivial", false, "restrict three-way merge by git read-tree to happen only if there is no file-level merging required")
	readTreeCmd.Flags().BoolS("u", "u", false, "after a successful merge, update the files in the work tree with the result of the merge")
	readTreeCmd.Flags().BoolS("v", "v", false, "show the progress of checking files out")
	rootCmd.AddCommand(readTreeCmd)

	carapace.Gen(readTreeCmd).FlagCompletion(carapace.ActionMap{
		"index-output": carapace.ActionFiles(),
		"prefix":       carapace.ActionDirectories(),
	})

	carapace.Gen(readTreeCmd).PositionalCompletion(
		git.ActionRefs(git.RefOption{}.Default()),
		git.ActionRefs(git.RefOption{}.Default()),
		git.ActionRefs(git.RefOption{}.Default()),
	)
}
