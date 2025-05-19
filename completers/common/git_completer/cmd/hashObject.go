package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hashObjectCmd = &cobra.Command{
	Use:     "hash-object",
	Short:   "Compute object ID and optionally creates a blob from a file",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_low_level_manipulator].ID,
}

func init() {
	carapace.Gen(hashObjectCmd).Standalone()

	hashObjectCmd.Flags().Bool("literally", false, "Allow --stdin to hash any garbage into a loose object")
	hashObjectCmd.Flags().Bool("no-filters", false, "Hash the contents as is")
	hashObjectCmd.Flags().String("path", "", "Hash object as if it were located at the given path")
	hashObjectCmd.Flags().Bool("stdin", false, "Read the object from standard input instead of from a file")
	hashObjectCmd.Flags().Bool("stdin-paths", false, "Read file names from the standard input")
	hashObjectCmd.Flags().StringS("t", "t", "", "Specify the type of object to be created")
	hashObjectCmd.Flags().BoolS("w", "w", false, "Actually write the object into the object database")
	rootCmd.AddCommand(hashObjectCmd)

	carapace.Gen(hashObjectCmd).FlagCompletion(carapace.ActionMap{
		"path": carapace.ActionFiles(),
		"t":    carapace.ActionValues("commit", "tree", "blob", "tag"),
	})

	carapace.Gen(hashObjectCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if hashObjectCmd.Flag("stdin").Changed ||
				hashObjectCmd.Flag("stdin-paths").Changed {
				return carapace.ActionValues()
			}
			return carapace.ActionFiles()
		}),
	)

	carapace.Gen(hashObjectCmd).DashAnyCompletion(
		carapace.ActionPositional(hashObjectCmd),
	)
}
