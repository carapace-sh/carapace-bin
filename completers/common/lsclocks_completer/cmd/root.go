package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/lsclocks_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lsclocks",
	Short: "display system clocks",
	Long:  "https://man7.org/linux/man-pages/man1/lsclocks.1.html",
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

	rootCmd.Flags().StringP("cpu-clock", "c", "", "also display CPU clock of specified process")
	rootCmd.Flags().StringP("dynamic-clock", "d", "", "also display specified dynamic clock")
	rootCmd.Flags().BoolP("help", "h", false, "display this help")
	rootCmd.Flags().BoolP("json", "J", false, "use JSON output format")
	rootCmd.Flags().Bool("no-discover-dynamic", false, "do not try to discover dynamic clocks")
	rootCmd.Flags().BoolP("noheadings", "n", false, "don't print headings")
	rootCmd.Flags().StringP("output", "o", "", "output columns")
	rootCmd.Flags().Bool("output-all", false, "output all columns")
	rootCmd.Flags().BoolP("raw", "r", false, "use raw output format")
	rootCmd.Flags().StringP("time", "t", "", "show current time of single clock")
	rootCmd.Flags().BoolP("version", "V", false, "display version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cpu-clock":     ps.ActionProcessIds(),
		"dynamic-clock": carapace.ActionFiles(),
		"output":        action.ActionColumns().UniqueList(","),
		"time":          action.ActionClocks(),
	})
}
