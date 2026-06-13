package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ulimit",
	Short: "Modify shell resource limits",
	Long:  "https://www.gnu.org/software/bash/manual/html_node/Bash-Builtins.html#index-ulimit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("H", "H", false, "use the hard resource limit")
	rootCmd.Flags().BoolS("P", "P", false, "the maximum number of pseudoterminals")
	rootCmd.Flags().BoolS("R", "R", false, "the maximum time a real-time process can run before blocking")
	rootCmd.Flags().BoolS("S", "S", false, "use the soft resource limit")
	rootCmd.Flags().BoolS("T", "T", false, "the maximum number of threads")
	rootCmd.Flags().BoolS("a", "a", false, "all current limits are reported")
	rootCmd.Flags().BoolS("b", "b", false, "the socket buffer size")
	rootCmd.Flags().BoolS("c", "c", false, "the maximum size of core files created")
	rootCmd.Flags().BoolS("d", "d", false, "the maximum size of a process's data segment")
	rootCmd.Flags().BoolS("e", "e", false, "the maximum scheduling priority")
	rootCmd.Flags().BoolS("f", "f", false, "the maximum size of files written by the shell")
	rootCmd.Flags().BoolS("i", "i", false, "the maximum number of pending signals")
	rootCmd.Flags().BoolS("k", "k", false, "the maximum number of kqueues allocated for this process")
	rootCmd.Flags().BoolS("l", "l", false, "the maximum size a process may lock into memory")
	rootCmd.Flags().BoolS("m", "m", false, "the maximum resident set size")
	rootCmd.Flags().BoolS("n", "n", false, "the maximum number of open file descriptors")
	rootCmd.Flags().BoolS("p", "p", false, "the pipe buffer size")
	rootCmd.Flags().BoolS("q", "q", false, "the maximum number of bytes in POSIX message queues")
	rootCmd.Flags().BoolS("r", "r", false, "the maximum real-time scheduling priority")
	rootCmd.Flags().BoolS("s", "s", false, "the maximum stack size")
	rootCmd.Flags().BoolS("t", "t", false, "the maximum amount of cpu time in seconds")
	rootCmd.Flags().BoolS("u", "u", false, "the maximum number of user processes")
	rootCmd.Flags().BoolS("v", "v", false, "the size of virtual memory")
	rootCmd.Flags().BoolS("x", "x", false, "the maximum number of file locks")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"hard", "current hard limit",
			"soft", "current soft limit",
			"unlimited", "no limit",
		),
	)
}
