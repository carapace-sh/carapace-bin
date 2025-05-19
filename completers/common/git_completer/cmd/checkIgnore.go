package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var checkIgnoreCmd = &cobra.Command{
	Use:     "check-ignore",
	Short:   "Debug gitignore / exclude files",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_helper].ID,
}

func init() {
	carapace.Gen(checkIgnoreCmd).Standalone()

	checkIgnoreCmd.Flags().Bool("no-index", false, "ignore index when checking")
	checkIgnoreCmd.Flags().BoolP("non-matching", "n", false, "show non-matching input paths")
	checkIgnoreCmd.Flags().BoolP("quiet", "q", false, "suppress progress reporting")
	checkIgnoreCmd.Flags().Bool("stdin", false, "read file names from stdin")
	checkIgnoreCmd.Flags().BoolP("verbose", "v", false, "be verbose")
	checkIgnoreCmd.Flags().BoolS("z", "z", false, "terminate input and output records by a NUL character")
	rootCmd.AddCommand(checkIgnoreCmd)

	carapace.Gen(checkIgnoreCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if checkIgnoreCmd.Flag("stdin").Changed {
				return carapace.ActionValues()
			}
			return carapace.ActionFiles()
		}),
	)

	carapace.Gen(checkIgnoreCmd).DashAnyCompletion(
		carapace.ActionPositional(checkIgnoreCmd),
	)
}
