package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ulimit",
	Short: "Set resource limits",
	Long:  "https://fishshell.com/docs/current/cmds/ulimit.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("H", "H", false, "hard limit")
	rootCmd.Flags().BoolS("K", "K", false, "kernel queues")
	rootCmd.Flags().BoolS("P", "P", false, "pseudo-terminals")
	rootCmd.Flags().BoolS("S", "S", false, "soft limit")
	rootCmd.Flags().BoolS("T", "T", false, "threads")
	rootCmd.Flags().BoolS("a", "a", false, "all limits")
	rootCmd.Flags().Bool("all", false, "all limits")
	rootCmd.Flags().BoolS("b", "b", false, "socket buffer size")
	rootCmd.Flags().BoolS("c", "c", false, "core file size")
	rootCmd.Flags().Bool("core-size", false, "core file size")
	rootCmd.Flags().Bool("cpu-time", false, "CPU time")
	rootCmd.Flags().BoolS("d", "d", false, "data segment size")
	rootCmd.Flags().Bool("data-size", false, "data segment size")
	rootCmd.Flags().BoolS("e", "e", false, "nice value")
	rootCmd.Flags().BoolS("f", "f", false, "file size")
	rootCmd.Flags().Bool("file-descriptor-count", false, "open file descriptors")
	rootCmd.Flags().Bool("file-size", false, "file size")
	rootCmd.Flags().BoolS("h", "h", false, "display help")
	rootCmd.Flags().Bool("hard", false, "hard limit")
	rootCmd.Flags().BoolS("i", "i", false, "queued signals")
	rootCmd.Flags().Bool("kernel-queues", false, "kernel queues")
	rootCmd.Flags().BoolS("l", "l", false, "locked memory")
	rootCmd.Flags().Bool("lock-size", false, "locked memory")
	rootCmd.Flags().BoolS("m", "m", false, "resident set size")
	rootCmd.Flags().BoolS("n", "n", false, "open file descriptors")
	rootCmd.Flags().Bool("nice", false, "nice value")
	rootCmd.Flags().Bool("pending-signals", false, "queued signals")
	rootCmd.Flags().Bool("process-count", false, "user processes")
	rootCmd.Flags().Bool("ptys", false, "pseudo-terminals")
	rootCmd.Flags().BoolS("q", "q", false, "message queue size")
	rootCmd.Flags().Bool("queue-size", false, "message queue size")
	rootCmd.Flags().BoolS("r", "r", false, "realtime priority")
	rootCmd.Flags().Bool("realtime-maxtime", false, "realtime CPU time")
	rootCmd.Flags().Bool("realtime-priority", false, "realtime priority")
	rootCmd.Flags().Bool("resident-set-size", false, "resident set size")
	rootCmd.Flags().BoolS("s", "s", false, "stack size")
	rootCmd.Flags().Bool("socket-buffers", false, "socket buffer size")
	rootCmd.Flags().Bool("soft", false, "soft limit")
	rootCmd.Flags().Bool("stack-size", false, "stack size")
	rootCmd.Flags().Bool("swap-size", false, "swap space")
	rootCmd.Flags().BoolS("t", "t", false, "CPU time")
	rootCmd.Flags().Bool("threads", false, "threads")
	rootCmd.Flags().BoolS("u", "u", false, "user processes")
	rootCmd.Flags().BoolS("v", "v", false, "virtual memory")
	rootCmd.Flags().Bool("virtual-memory-size", false, "virtual memory")
	rootCmd.Flags().BoolS("w", "w", false, "swap space")
	rootCmd.Flags().BoolS("y", "y", false, "realtime CPU time")
}
