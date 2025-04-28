package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/lsirq_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lsirq",
	Short: "utility to display kernel interrupt information",
	Long:  "https://man7.org/linux/man-pages/man1/lsirq.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("cpu-list", "C", "", "only show counters for these CPUs")
	rootCmd.Flags().BoolP("help", "h", false, "display this help")
	rootCmd.Flags().BoolP("json", "J", false, "use JSON output format")
	rootCmd.Flags().BoolP("noheadings", "n", false, "don't print headings")
	rootCmd.Flags().StringP("output", "o", "", "define which output columns to use")
	rootCmd.Flags().BoolP("pairs", "P", false, "use key=\"value\" output format")
	rootCmd.Flags().BoolP("softirq", "S", false, "show softirqs instead of interrupts")
	rootCmd.Flags().StringP("sort", "s", "", "specify sort column")
	rootCmd.Flags().StringP("threshold", "t", "", "only IRQs with counters above <N>")
	rootCmd.Flags().BoolP("version", "V", false, "display version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"output": action.ActionColumns().UniqueList(","),
		"sort":   action.ActionColumns(),
	})
}
