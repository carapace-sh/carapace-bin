package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dtruss",
	Short: "Trace system calls made by a process",
	Long:  "https://keith.github.io/xcode-manpages/dtruss.1m.html",
	Run:   func(*cobra.Command, []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all", "a", false, "print all details")
	rootCmd.Flags().StringP("buffer", "b", "", "dynamic variable buffer size")
	rootCmd.Flags().BoolP("counts", "c", false, "print syscall counts")
	rootCmd.Flags().BoolP("delta", "d", false, "print relative times (us)")
	rootCmd.Flags().BoolP("elapsed", "e", false, "print elapsed times (us)")
	rootCmd.Flags().BoolP("follow", "f", false, "follow children")
	rootCmd.Flags().BoolP("help", "h", false, "print usage")
	rootCmd.Flags().BoolP("lwpid", "l", false, "force printing pid/lwpid")
	rootCmd.Flags().StringP("name", "n", "", "examine this process name")
	rootCmd.Flags().BoolP("nolwpid", "L", false, "don't print pid/lwpid")
	rootCmd.Flags().BoolP("oncpu", "o", false, "print on cpu times")
	rootCmd.Flags().StringP("pid", "p", "", "examine this PID")
	rootCmd.Flags().BoolP("stack", "s", false, "print stack backtraces")
	rootCmd.Flags().StringP("syscall", "t", "", "examine this syscall only")
	rootCmd.Flags().StringP("wait", "W", "", "wait for a process matching this name")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"syscall": carapace.ActionValues(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
