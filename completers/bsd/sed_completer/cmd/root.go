package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sed",
	Short: "stream editor",
	Long:  "https://manpage.me/?sed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("E", "E", false, "Interpret regular expressions as extended regular expressions")
	rootCmd.Flags().BoolS("H", "H", false, "Enable enhanced features in the regular expression syntax")
	rootCmd.Flags().StringS("I", "I", "", "Edit files in-place, saving backups with the specified extension")
	rootCmd.Flags().BoolS("a", "a", false, "Create or truncate files before any processing begins")
	rootCmd.Flags().StringS("e", "e", "", "Append the editing commands specified by the command argument to the list of commands")
	rootCmd.Flags().StringS("f", "f", "", "Append the editing commands found in the file command_file to the list of commands")
	rootCmd.Flags().StringS("i", "i", "", "Edit files in-place similarly to -I, but treat each file independently from other files")
	rootCmd.Flags().BoolS("l", "l", false, "Make output line buffered")
	rootCmd.Flags().BoolS("n", "n", false, "Suppress output")
	rootCmd.Flags().BoolS("r", "r", false, "Same as -E for compatibility with GNU sed")
	rootCmd.Flags().BoolS("u", "u", false, "Make output unbuffered")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"f": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flag("e").Changed || rootCmd.Flag("f").Changed {
				return carapace.ActionFiles()
			} else {
				return carapace.ActionValues()
			}
		}),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
