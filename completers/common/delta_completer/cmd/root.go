package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/delta"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "delta [OPTIONS] <MINUS_FILE> <PLUS_FILE>",
	Short: "A viewer for git and diff output",
	Long:  "https://github.com/dandavison/delta",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("blame-code-style", "", "Style string for the code section of a git blame line")
	rootCmd.Flags().String("blame-format", "", "Format string for git blame commit metadata")
	rootCmd.Flags().String("blame-palette", "", "Background colors used for git blame lines (space-separated string)")
	rootCmd.Flags().String("blame-separator-format", "", "Separator between the blame format and the code section of a git blame line")
	rootCmd.Flags().String("blame-separator-style", "", "Style string for the blame-separator-format")
	rootCmd.Flags().String("blame-timestamp-format", "", "Format of `git blame` timestamp in raw git output received by delta")
	rootCmd.Flags().String("blame-timestamp-output-format", "", "Format string for git blame timestamp output")
	rootCmd.Flags().Bool("color-only", false, "Do not alter the input structurally in any way")
	rootCmd.Flags().String("commit-decoration-style", "", "Style string for the commit hash decoration")
	rootCmd.Flags().String("commit-regex", "", "Regular expression used to identify the commit line when parsing git output")
	rootCmd.Flags().String("commit-style", "", "Style string for the commit hash line")
	rootCmd.Flags().Bool("dark", false, "Use default colors appropriate for a dark terminal background")
	rootCmd.Flags().String("default-language", "", "Default language used for syntax highlighting")
	rootCmd.Flags().Bool("diff-highlight", false, "Emulate diff-highlight")
	rootCmd.Flags().Bool("diff-so-fancy", false, "Emulate diff-so-fancy")
	rootCmd.Flags().String("diff-stat-align-width", "", "Width allocated for file paths in a diff stat section")
	rootCmd.Flags().String("features", "", "Names of delta features to activate (space-separated)")
	rootCmd.Flags().String("file-added-label", "", "Text to display before an added file path")
	rootCmd.Flags().String("file-copied-label", "", "Text to display before a copied file path")
	rootCmd.Flags().String("file-decoration-style", "", "Style string for the file decoration")
	rootCmd.Flags().String("file-modified-label", "", "Text to display before a modified file path")
	rootCmd.Flags().String("file-removed-label", "", "Text to display before a removed file path")
	rootCmd.Flags().String("file-renamed-label", "", "Text to display before a renamed file path")
	rootCmd.Flags().String("file-style", "", "Style string for the file section")
	rootCmd.Flags().String("file-transformation", "", "Sed-style command transforming file paths for display")
	rootCmd.Flags().String("grep-context-line-style", "", "Style string for non-matching lines of grep output")
	rootCmd.Flags().String("grep-file-style", "", "Style string for file paths in grep output")
	rootCmd.Flags().String("grep-line-number-style", "", "Style string for line numbers in grep output")
	rootCmd.Flags().String("grep-match-line-style", "", "Style string for matching lines of grep output")
	rootCmd.Flags().String("grep-match-word-style", "", "Style string for the matching substrings within a matching line of grep output")
	rootCmd.Flags().String("grep-separator-symbol", "", "Separator symbol printed after the file path and line number in grep output")
	rootCmd.Flags().BoolP("help", "h", false, "Print help information")
	rootCmd.Flags().String("hunk-header-decoration-style", "", "Style string for the hunk-header decoration")
	rootCmd.Flags().String("hunk-header-file-style", "", "Style string for the file path part of the hunk-header")
	rootCmd.Flags().String("hunk-header-line-number-style", "", "Style string for the line number part of the hunk-header")
	rootCmd.Flags().String("hunk-header-style", "", "Style string for the hunk-header")
	rootCmd.Flags().String("hunk-label", "", "Text to display before a hunk header")
	rootCmd.Flags().Bool("hyperlinks", false, "Render commit hashes, file names, and line numbers as hyperlinks")
	rootCmd.Flags().String("hyperlinks-commit-link-format", "", "Format string for commit hyperlinks (requires --hyperlinks)")
	rootCmd.Flags().String("hyperlinks-file-link-format", "", "Format string for file hyperlinks (requires --hyperlinks)")
	rootCmd.Flags().String("inline-hint-style", "", "Style string for short inline hint text")
	rootCmd.Flags().String("inspect-raw-lines", "", "Kill-switch for --color-moved support")
	rootCmd.Flags().Bool("keep-plus-minus-markers", false, "Prefix added/removed lines with a +/- character, as git does")
	rootCmd.Flags().Bool("light", false, "Use default colors appropriate for a light terminal background")
	rootCmd.Flags().String("line-buffer-size", "", "Size of internal line buffer")
	rootCmd.Flags().String("line-fill-method", "", "Line-fill method in side-by-side mode")
	rootCmd.Flags().BoolP("line-numbers", "n", false, "Display line numbers next to the diff")
	rootCmd.Flags().String("line-numbers-left-format", "", "Format string for the left column of line numbers")
	rootCmd.Flags().String("line-numbers-left-style", "", "Style string for the left column of line numbers")
	rootCmd.Flags().String("line-numbers-minus-style", "", "Style string for line numbers in the old (minus) version of the file")
	rootCmd.Flags().String("line-numbers-plus-style", "", "Style string for line numbers in the new (plus) version of the file")
	rootCmd.Flags().String("line-numbers-right-format", "", "Format string for the right column of line numbers")
	rootCmd.Flags().String("line-numbers-right-style", "", "Style string for the right column of line numbers")
	rootCmd.Flags().String("line-numbers-zero-style", "", "Style string for line numbers in unchanged (zero) lines")
	rootCmd.Flags().Bool("list-languages", false, "List supported languages and associated file extensions")
	rootCmd.Flags().Bool("list-syntax-themes", false, "List available syntax-highlighting color themes")
	rootCmd.Flags().String("map-styles", "", "Map styles encountered in raw input to desired output styles")
	rootCmd.Flags().String("max-line-distance", "", "Maximum line pair distance parameter in within-line diff algorithm")
	rootCmd.Flags().String("max-line-length", "", "Truncate lines longer than this")
	rootCmd.Flags().String("merge-conflict-begin-symbol", "", "String marking the beginning of a merge conflict region")
	rootCmd.Flags().String("merge-conflict-end-symbol", "", "String marking the end of a merge conflict region")
	rootCmd.Flags().String("merge-conflict-ours-diff-header-decoration-style", "", "Style string for the decoration of the header above the 'ours' merge conflict diff")
	rootCmd.Flags().String("merge-conflict-ours-diff-header-style", "", "Style string for the header above the 'ours' branch merge conflict diff")
	rootCmd.Flags().String("merge-conflict-theirs-diff-header-decoration-style", "", "Style string for the decoration of the header above the 'theirs' merge conflict diff")
	rootCmd.Flags().String("merge-conflict-theirs-diff-header-style", "", "Style string for the header above the 'theirs' branch merge conflict diff")
	rootCmd.Flags().String("minus-emph-style", "", "Style string for emphasized sections of removed lines")
	rootCmd.Flags().String("minus-empty-line-marker-style", "", "Style string for removed empty line marker")
	rootCmd.Flags().String("minus-non-emph-style", "", "Style string for non-emphasized sections of removed lines that have an emphasized section")
	rootCmd.Flags().String("minus-style", "", "Style string for removed lines")
	rootCmd.Flags().Bool("navigate", false, "Activate diff navigation")
	rootCmd.Flags().String("navigate-regex", "", "Regular expression defining navigation stop points")
	rootCmd.Flags().Bool("no-gitconfig", false, "Do not read any settings from git config")
	rootCmd.Flags().String("pager", "", "Which pager to use")
	rootCmd.Flags().String("paging", "", "Whether to use a pager when displaying output")
	rootCmd.Flags().Bool("parse-ansi", false, "Display ANSI color escape sequences in human-readable form")
	rootCmd.Flags().String("plus-emph-style", "", "Style string for emphasized sections of added lines")
	rootCmd.Flags().String("plus-empty-line-marker-style", "", "Style string for added empty line marker")
	rootCmd.Flags().String("plus-non-emph-style", "", "Style string for non-emphasized sections of added lines that have an emphasized section")
	rootCmd.Flags().String("plus-style", "", "Style string for added lines")
	rootCmd.Flags().Bool("raw", false, "Do not alter the input in any way")
	rootCmd.Flags().Bool("relative-paths", false, "Output all file paths relative to the current directory")
	rootCmd.Flags().String("right-arrow", "", "Text to display with a changed file path")
	rootCmd.Flags().Bool("show-colors", false, "Show available named colors")
	rootCmd.Flags().Bool("show-config", false, "Display the active values for all Delta options")
	rootCmd.Flags().Bool("show-syntax-themes", false, "Show example diff for available syntax-highlighting themes")
	rootCmd.Flags().Bool("show-themes", false, "Show example diff for available delta themes")
	rootCmd.Flags().BoolP("side-by-side", "s", false, "Display diffs in side-by-side layout")
	rootCmd.Flags().String("syntax-theme", "", "The syntax-highlighting theme to use")
	rootCmd.Flags().String("tabs", "", "The number of spaces to replace tab characters with")
	rootCmd.Flags().String("true-color", "", "Whether to emit 24-bit (\"true color\") RGB color codes")
	rootCmd.Flags().BoolP("version", "V", false, "Print version information")
	rootCmd.Flags().String("whitespace-error-style", "", "Style string for whitespace errors")
	rootCmd.Flags().StringP("width", "w", "", "The width of underline/overline decorations")
	rootCmd.Flags().String("word-diff-regex", "", "Regular expression defining a 'word' in within-line diff algorithm")
	rootCmd.Flags().String("wrap-left-symbol", "", "End-of-line wrapped content symbol (left-aligned)")
	rootCmd.Flags().String("wrap-max-lines", "", "How often a line should be wrapped if it does not fit")
	rootCmd.Flags().String("wrap-right-percent", "", "Threshold for right-aligning wrapped content")
	rootCmd.Flags().String("wrap-right-prefix-symbol", "", "Pre-wrapped content symbol (right-aligned)")
	rootCmd.Flags().String("wrap-right-symbol", "", "End-of-line wrapped content symbol (right-aligned)")
	rootCmd.Flags().String("zero-style", "", "Style string for unchanged lines")

	// TODO complete styles
	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"inspect-raw-lines": carapace.ActionValues("true", "false").StyleF(style.ForKeyword),
		"pager": carapace.Batch(
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
		"paging":       carapace.ActionValues("auto", "always", "never").StyleF(style.ForKeyword),
		"syntax-theme": delta.ActionSyntaxThemes(),
		"true-color":   carapace.ActionValues("auto", "always", "never").StyleF(style.ForKeyword),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
