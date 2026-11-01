package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:     "rm",
	Short:   "Remove tracked paths from the working tree and index",
	GroupID: "change",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rmCmd).Standalone()

	rmCmd.Flags().Bool("cached", false, "Remove from the index only; keep files in the working tree")
	rmCmd.Flags().Bool("dry-run", false, "List paths that would be removed without changing the index or work tree")
	rmCmd.Flags().BoolP("force", "f", false, "Remove even when paths have local modifications")
	rmCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rmCmd.Flags().Bool("ignore-unmatch", false, "Treat non-matching pathspecs as success with an empty result")
	rmCmd.Flags().Bool("pathspec-file-nul", false, "`--pathspec-from-file` entries are NUL-separated")
	rmCmd.Flags().String("pathspec-from-file", "", "Read pathspecs from a file (`-` for stdin) instead of or in addition to argv paths")
	rootCmd.AddCommand(rmCmd)

	carapace.Gen(rmCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
