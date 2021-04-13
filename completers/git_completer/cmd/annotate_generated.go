package cmd

import (
	"github.com/spf13/cobra"
)

var annotateCmd = &cobra.Command{
	Use:   "annotate",
	Short: "Annotate file lines with commit information",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	annotateCmd.Flags().StringS("C", "C", "", "Find line copies within and across files")
	annotateCmd.Flags().StringS("L", "L", "", "Process only line range n,m, counting from 1")
	annotateCmd.Flags().StringS("M", "M", "", "Find line movements within and across files")
	annotateCmd.Flags().StringS("S", "S", "", "Use revisions from <file> instead of calling git-rev-list")
	annotateCmd.Flags().String("abbrev", "", "use <n> digits to display SHA-1s")
	annotateCmd.Flags().BoolS("b", "b", false, "Show blank SHA-1 for boundary commits (Default: off)")
	annotateCmd.Flags().BoolS("c", "c", false, "Use the same output mode as git-annotate (Default: off)")
	annotateCmd.Flags().Bool("color-by-age", false, "color lines by age")
	annotateCmd.Flags().Bool("color-lines", false, "color redundant metadata from previous line differently")
	annotateCmd.Flags().String("contents", "", "Use <file>'s contents as the final image")
	annotateCmd.Flags().String("ignore-rev", "", "Ignore <rev> when blaming")
	annotateCmd.Flags().String("ignore-revs-file", "", "Ignore revisions from <file>")
	annotateCmd.Flags().Bool("incremental", false, "Show blame entries as we find them, incrementally")
	annotateCmd.Flags().BoolS("l", "l", false, "Show long commit SHA1 (Default: off)")
	annotateCmd.Flags().Bool("line-porcelain", false, "Show porcelain format with per-line commit information")
	annotateCmd.Flags().Bool("minimal", false, "Spend extra cycles to find better match")
	annotateCmd.Flags().BoolP("porcelain", "p", false, "Show in a format designed for machine consumption")
	annotateCmd.Flags().Bool("progress", false, "Force progress reporting")
	annotateCmd.Flags().Bool("root", false, "Do not treat root commits as boundaries (Default: off)")
	annotateCmd.Flags().BoolS("s", "s", false, "Suppress author name and timestamp (Default: off)")
	annotateCmd.Flags().Bool("score-debug", false, "Show output score for blame entries")
	annotateCmd.Flags().BoolP("show-email", "e", false, "Show author email instead of name (Default: off)")
	annotateCmd.Flags().BoolP("show-name", "f", false, "Show original filename (Default: auto)")
	annotateCmd.Flags().BoolP("show-number", "n", false, "Show original linenumber (Default: off)")
	annotateCmd.Flags().Bool("show-stats", false, "Show work cost statistics")
	annotateCmd.Flags().BoolS("t", "t", false, "Show raw timestamp (Default: off)")
	annotateCmd.Flags().BoolS("w", "w", false, "Ignore whitespace differences")
	rootCmd.AddCommand(annotateCmd)
}
