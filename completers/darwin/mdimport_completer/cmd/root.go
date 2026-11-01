package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mdimport",
	Short: "import files into Spotlight index",
	Long:  "https://keith.github.io/xcode-manpages/mdimport.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")
	rootCmd.Flags().BoolP("list", "A", false, "List all installed metadata importers")
	rootCmd.Flags().BoolP("list-plugins", "L", false, "List known plug-ins")
	rootCmd.Flags().BoolP("modern-importers", "m", false, "Import using modern importers")
	rootCmd.Flags().StringP("outfile", "o", "", "Save test-import attributes to file")
	rootCmd.Flags().BoolP("performance", "p", false, "Print performance information")
	rootCmd.Flags().BoolP("reimport", "r", false, "Reimport files for UTIs claimed by plug-in")
	rootCmd.Flags().BoolP("test", "t", false, "Test the metadata importer")
	rootCmd.Flags().StringP("type", "y", "", "File type for modern importers")
	rootCmd.Flags().StringP("urls", "u", "", "URLs for files (requires -m)")
	rootCmd.Flags().BoolP("verbose", "v", false, "Verbose mode")
	rootCmd.Flags().BoolP("xml", "X", false, "Print the schema file")

	rootCmd.Flags().StringP("debug-level", "d", "", "Debug level (1-3) for test-import output")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"debug-level": carapace.ActionValues("1", "2", "3"),
		"outfile":     carapace.ActionFiles(),
		"type":        carapace.ActionValues(),
		"urls":        carapace.ActionValues(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
