package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Show commit logs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logCmd).Standalone()

	logCmd.Flags().StringArrayP("authors-file", "A", []string{}, "Authors file")
	logCmd.Flags().Bool("color", false, "Color output")
	logCmd.Flags().Bool("incremental", false, "Incremental output")
	logCmd.Flags().Int("limit", 0, "Limit number of log entries")
	logCmd.Flags().Bool("non-recursive", false, "Non-recursive")
	logCmd.Flags().Bool("oneline", false, "One line per commit")
	logCmd.Flags().String("pager", "", "Pager to use")
	logCmd.Flags().StringP("revision", "r", "", "Revision range")
	logCmd.Flags().Bool("show-commit", false, "Show the Git commit sha1")
	logCmd.Flags().BoolP("verbose", "v", false, "Verbose")
	rootCmd.AddCommand(logCmd)

	carapace.Gen(logCmd).FlagCompletion(carapace.ActionMap{
		"authors-file": carapace.ActionFiles(),
		"pager":        carapace.ActionValues("less", "more", "cat"),
		"revision":     carapace.ActionValues(),
	})
}
