package cmd

import (
	"github.com/spf13/cobra"
)

var blameCmd = &cobra.Command{
	Use:   "blame",
	Short: "Show what revision and author last modified each line of a file",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	blameCmd.Flags().String("abbrev", "", "use <n> digits to display SHA-1s")
	blameCmd.Flags().BoolP("b", "b", false, "Show blank SHA-1 for boundary commits (Default: off)")
	blameCmd.Flags().Bool("color-by-age", false, "color lines by age")
	blameCmd.Flags().Bool("color-lines", false, "color redundant metadata from previous line differently")
	blameCmd.Flags().String("contents", "", "Use <file>'s contents as the final image")
	blameCmd.Flags().StringP("C", "C", "", "Find line copies within and across files")
	blameCmd.Flags().BoolP("c", "c", false, "Use the same output mode as git-annotate (Default: off)")
	blameCmd.Flags().BoolP("show-email", "e", false, "Show author email instead of name (Default: off)")
	blameCmd.Flags().BoolP("show-name", "f", false, "Show original filename (Default: auto)")
	blameCmd.Flags().String("ignore-rev", "", "Ignore <rev> when blaming")
	blameCmd.Flags().String("ignore-revs-file", "", "Ignore revisions from <file>")
	blameCmd.Flags().Bool("incremental", false, "Show blame entries as we find them, incrementally")
	blameCmd.Flags().Bool("line-porcelain", false, "Show porcelain format with per-line commit information")
	blameCmd.Flags().StringP("L", "L", "", "Process only line range n,m, counting from 1")
	blameCmd.Flags().BoolP("l", "l", false, "Show long commit SHA1 (Default: off)")
	blameCmd.Flags().Bool("minimal", false, "Spend extra cycles to find better match")
	blameCmd.Flags().StringP("M", "M", "", "Find line movements within and across files")
	blameCmd.Flags().BoolP("show-number", "n", false, "Show original linenumber (Default: off)")
	blameCmd.Flags().BoolP("porcelain", "p", false, "Show in a format designed for machine consumption")
	blameCmd.Flags().Bool("progress", false, "Force progress reporting")
	blameCmd.Flags().Bool("root", false, "Do not treat root commits as boundaries (Default: off)")
	blameCmd.Flags().Bool("score-debug", false, "Show output score for blame entries")
	blameCmd.Flags().StringP("S", "S", "", "Use revisions from <file> instead of calling git-rev-list")
	blameCmd.Flags().Bool("show-stats", false, "Show work cost statistics")
	blameCmd.Flags().BoolP("s", "s", false, "Suppress author name and timestamp (Default: off)")
	blameCmd.Flags().BoolP("t", "t", false, "Show raw timestamp (Default: off)")
	blameCmd.Flags().BoolP("w", "w", false, "Ignore whitespace differences")
	rootCmd.AddCommand(blameCmd)
}
