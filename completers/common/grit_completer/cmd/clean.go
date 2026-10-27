package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove untracked files from the working tree",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanCmd).Standalone()

	cleanCmd.Flags().Bool("dry-run", false, "Explicit dry run (default when `-f` is not passed)")
	cleanCmd.Flags().StringSliceP("exclude", "e", nil, "Additional exclude patterns; matching paths are not removed")
	cleanCmd.Flags().Bool("files-only", false, "Remove only files, not untracked directories")
	cleanCmd.Flags().BoolP("force", "f", false, "Actually delete files. Without this flag, only prints what would be removed")
	cleanCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	cleanCmd.Flags().Bool("ignored", false, "Remove ignored files as well (like `git clean -x`)")
	cleanCmd.Flags().Bool("only-ignored", false, "Remove only ignored files (like `git clean -X`)")
	cleanCmd.Flags().Bool("pathspec-file-nul", false, "`--pathspec-from-file` entries are NUL-separated")
	cleanCmd.Flags().String("pathspec-from-file", "", "Read pathspecs from a file (`-` for stdin) instead of or in addition to argv paths")
	rootCmd.AddCommand(cleanCmd)

	carapace.Gen(cleanCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
