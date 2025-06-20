package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sd",
	Short: "Intuitive find & replace CLI",
	Long:  "https://github.com/chmln/sd",
	Run:   func(*cobra.Command, []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("preview", "p", false, "Display changes in a human reviewable format (the specifics of the format are likely to change in the future)")
	rootCmd.Flags().BoolP("fixed-strings", "F", false, "Treat FIND and REPLACE_WITH args as literal strings")
	rootCmd.Flags().IntP("max-replacements", "n", 0, "Limit the number of replacements that can occur per file. 0 indicates unlimited replacements")
	rootCmd.Flags().StringP("flags", "f", "", "Regex flags. May be combined")
	rootCmd.Flags().BoolP("help", "h", false, "Print help (see a summary with '-h')")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"flags": carapace.ActionValuesDescribed(
			"c", "case-sensitive",
			"e", "disable multi-line matching",
			"i", "case-insensitive",
			"m", "multi-line matching",
			"s", "make `.` match newlines",
			"w", "match full words only",
		).UniqueList(""),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionMessage("string / regex to find"),
		carapace.ActionMessage("string to replace with"),
		carapace.ActionFiles(),
	)
}
