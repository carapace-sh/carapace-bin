package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "at",
	Short: "queue, examine or delete jobs for later execution",
	Long:  "https://man.freebsd.org/cgi/man.cgi?at",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("b", "b", false, "Alias for batch")
	rootCmd.Flags().BoolS("c", "c", false, "Cat jobs to stdout")
	rootCmd.Flags().BoolS("d", "d", false, "Alias for atrm (deprecated)")
	rootCmd.Flags().StringS("f", "f", "", "Read job from file")
	rootCmd.Flags().BoolS("l", "l", false, "List jobs")
	rootCmd.Flags().BoolS("m", "m", false, "Send mail on completion")
	rootCmd.Flags().StringS("q", "q", "", "Use specified queue")
	rootCmd.Flags().BoolS("r", "r", false, "Remove specified jobs")
	rootCmd.Flags().StringS("t", "t", "", "POSIX time format")
	rootCmd.Flags().BoolS("v", "v", false, "Show completed jobs")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"f": carapace.ActionFiles(),
		"q": carapace.ActionValues("a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"),
	})
}
