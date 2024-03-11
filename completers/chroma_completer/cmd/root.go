package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/chroma"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chroma",
	Short: "Chroma is a general purpose syntax highlighter",
	Long:  "https://github.com/alecthomas/chroma",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("check", false, "Do not format, check for tokenisation errors instead.")
	rootCmd.Flags().Bool("fail", false, "Exit silently with status 1 if no specific lexer was found.")
	rootCmd.Flags().String("filename", "", "Filename to use for selecting a lexer when reading from stdin.")
	rootCmd.Flags().StringP("formatter", "f", "", "Formatter to use.")
	rootCmd.Flags().BoolP("help", "h", false, "Show context-sensitive help.")
	rootCmd.Flags().Bool("html", false, "Convenience flag to use HTML formatter.")
	rootCmd.Flags().Bool("html-all-styles", false, "Output all HTML CSS styles, including redundant ones.")
	rootCmd.Flags().String("html-base-line", "", "Base line number.")
	rootCmd.Flags().String("html-highlight", "", "Highlight these lines.")
	rootCmd.Flags().String("html-highlight-style", "", "Style used for highlighting lines.")
	rootCmd.Flags().Bool("html-inline-styles", false, "Output HTML with inline styles (no classes).")
	rootCmd.Flags().Bool("html-lines", false, "Include line numbers in output.")
	rootCmd.Flags().String("html-lines-style", "", "Style for line numbers.")
	rootCmd.Flags().Bool("html-lines-table", false, "Split line numbers and code in a HTML table")
	rootCmd.Flags().Bool("html-linkable-lines", false, "Make the line numbers linkable and be a link to themselves.")
	rootCmd.Flags().Bool("html-only", false, "Output HTML fragment.")
	rootCmd.Flags().String("html-prefix", "", "HTML CSS class prefix.")
	rootCmd.Flags().Bool("html-prevent-surrounding-pre", false, "Prevent the surrounding pre tag.")
	rootCmd.Flags().Bool("html-styles", false, "Output HTML CSS styles.")
	rootCmd.Flags().String("html-tab-width", "", "Set the HTML tab width.")
	rootCmd.Flags().Bool("json", false, "Convenience flag to use JSON formatter.")
	rootCmd.Flags().StringP("lexer", "l", "", "Lexer to use when formatting or path to an XML file to load.")
	rootCmd.Flags().Bool("list", false, "List lexers, styles and formatters.")
	rootCmd.Flags().String("profile", "", "Enable profiling to file.")
	rootCmd.Flags().StringP("style", "s", "", "Style to use for formatting or path to an XML file to load.")
	rootCmd.Flags().Bool("svg", false, "Convenience flag to use SVG formatter.")
	rootCmd.Flags().Bool("trace", false, "Trace lexer states as they are traversed.")
	rootCmd.Flags().Bool("unbuffered", false, "Do not buffer output.")
	rootCmd.Flags().Bool("version", false, "Show version.")
	rootCmd.Flags().String("xml", "", "Generate XML lexer definitions.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"formatter": chroma.ActionFormatters(),
		"lexer":     chroma.ActionLexers(),
		"profile":   carapace.ActionFiles(),
		"style": carapace.Batch(
			chroma.ActionStyles(),
			carapace.ActionFiles(),
		).ToA(),
		"xml": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
