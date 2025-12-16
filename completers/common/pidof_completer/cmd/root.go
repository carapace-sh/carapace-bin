package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pidof",
	Short: "find the process ID of a running program",
	Long:  "https://man7.org/linux/man-pages/man1/pidof.1.html",
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

	rootCmd.Flags().BoolP("check-root", "c", false, "omit processes with different root")
	rootCmd.Flags().BoolP("help", "h", false, "display this help and exit")
	rootCmd.Flags().BoolP("lightweight", "t", false, "list threads too")
	rootCmd.Flags().StringP("omit-pid", "o", "", "omit processes with PID")
	rootCmd.Flags().BoolS("q", "q", false, "quiet mode, only set the exit code")
	rootCmd.Flags().StringP("separator", "S", "", "use SEP as separator put between PIDs")
	rootCmd.Flags().BoolP("single-shot", "s", false, "return one PID only")
	rootCmd.Flags().BoolP("version", "V", false, "output version information and exit")
	rootCmd.Flags().BoolP("with-workers", "w", false, "show kernel workers too")
	rootCmd.Flags().BoolS("x", "x", false, "also find shells running the named scripts")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"omit-pid": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.Batch(
				ps.ActionProcessIds(),
				carapace.ActionValuesDescribed("%PPID", "Parent process"),
			).ToA().UniqueList(",")
		}),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		ps.ActionProcessExecutables(),
	)
}
