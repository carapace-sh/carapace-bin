package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var checkoutIndexCmd = &cobra.Command{
	Use:     "checkout-index",
	Short:   "Copy files from the index to the working tree",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_manipulator].ID,
}

func init() {
	carapace.Gen(checkoutIndexCmd).Standalone()

	checkoutIndexCmd.Flags().BoolP("all", "a", false, "check out all files in the index")
	checkoutIndexCmd.Flags().BoolP("force", "f", false, "force overwrite of existing files")
	checkoutIndexCmd.Flags().Bool("ignore-skip-worktree-bits", false, "do not skip files with skip-worktree set")
	checkoutIndexCmd.Flags().BoolP("index", "u", false, "update stat information in the index file")
	checkoutIndexCmd.Flags().BoolP("no-create", "n", false, "don't checkout new files")
	checkoutIndexCmd.Flags().String("prefix", "", "when creating files, prepend <string>")
	checkoutIndexCmd.Flags().BoolP("quiet", "q", false, "no warning for existing files and files not in index")
	checkoutIndexCmd.Flags().String("stage", "", "copy out the files from named stage")
	checkoutIndexCmd.Flags().Bool("stdin", false, "read list of paths from the standard input")
	checkoutIndexCmd.Flags().Bool("temp", false, "write the content to temporary files")
	checkoutIndexCmd.Flags().BoolS("z", "z", false, "paths are separated with NUL character")
	rootCmd.AddCommand(checkoutIndexCmd)

	carapace.Gen(checkoutIndexCmd).FlagCompletion(carapace.ActionMap{
		"stage": carapace.ActionValues("1", "2", "3", "all"),
	})

	carapace.Gen(checkoutIndexCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if checkoutIndexCmd.Flag("stdin").Changed {
				return carapace.ActionValues()
			}
			return carapace.ActionFiles()
		}),
	)

	carapace.Gen(checkoutIndexCmd).DashAnyCompletion(
		carapace.ActionPositional(checkoutIndexCmd),
	)
}
