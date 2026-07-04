package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "top",
	Short: "display and update sorted information about processes",
	Long:  "https://keith.github.io/xcode-manpages/top.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("User", "U", "", "Display only processes for specified user")
	rootCmd.Flags().BoolP("all", "a", false, "Equivalent to -c a")
	rootCmd.Flags().StringP("command", "c", "", "Set command line/program name display mode: a, d, or e")
	rootCmd.Flags().BoolP("cpu", "u", false, "Alias for -o cpu -O time")
	rootCmd.Flags().StringP("delay", "s", "", "Set delay between updates in seconds")
	rootCmd.Flags().BoolP("delta", "d", false, "Equivalent to -c d")
	rootCmd.Flags().BoolP("extended", "e", false, "Equivalent to -c e")
	rootCmd.Flags().BoolP("frameworks", "f", false, "Calculate statistics on shared libraries")
	rootCmd.Flags().BoolP("help", "h", false, "Print command line usage and exit")
	rootCmd.Flags().StringP("interval", "i", "", "Set interval between updates")
	rootCmd.Flags().String("ncols", "", "Set number of columns to display")
	rootCmd.Flags().BoolP("no-frameworks", "F", false, "Do not calculate statistics on shared libraries")
	rootCmd.Flags().BoolP("no-recurse", "R", false, "Do not traverse memory object map for each process")
	rootCmd.Flags().StringP("nprocs", "n", "", "Set number of processes to display")
	rootCmd.Flags().StringP("order", "o", "", "Set primary sort key (descending)")
	rootCmd.Flags().String("pid", "", "Display only specified process ID")
	rootCmd.Flags().BoolP("recurse", "r", false, "Traverse and report memory object map for each process")
	rootCmd.Flags().StringP("samples", "l", "", "Set number of samples before exiting")
	rootCmd.Flags().StringP("secondaryKey", "O", "", "Set secondary sort key")
	rootCmd.Flags().String("stats", "", "Set custom statistics keys (comma-separated)")
	rootCmd.Flags().BoolP("swap", "S", false, "Display global statistics for swap and purgeable memory")
	rootCmd.Flags().String("user", "", "Display only processes for specified user")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"User":         os.ActionUsers(),
		"command":      carapace.ActionValues("a", "d", "e"),
		"order":        carapace.ActionValues("pid", "command", "cpu", "cpu_me", "cpu_others", "csw", "time", "threads", "ports", "mregion", "mem", "rprvt", "purg", "vsize", "vprvt", "kprvt", "kshrd", "pgrp", "ppid", "state", "uid", "wq", "faults", "cow", "user", "msgsent", "msgrecv", "sysbsd", "sysmach", "pageins", "boosts", "instrs", "cycles"),
		"secondaryKey": carapace.ActionValues("pid", "command", "cpu", "cpu_me", "cpu_others", "csw", "time", "threads", "ports", "mregion", "mem", "rprvt", "purg", "vsize", "vprvt", "kprvt", "kshrd", "pgrp", "ppid", "state", "uid", "wq", "faults", "cow", "user", "msgsent", "msgrecv", "sysbsd", "sysmach", "pageins", "boosts", "instrs", "cycles"),
		"stats":        carapace.ActionValues("pid", "command", "cpu", "cpu_me", "cpu_others", "csw", "time", "threads", "ports", "mregion", "mem", "rprvt", "purg", "vsize", "vprvt", "kprvt", "kshrd", "pgrp", "ppid", "state", "uid", "wq", "faults", "cow", "user", "msgsent", "msgrecv", "sysbsd", "sysmach", "pageins", "boosts", "instrs", "cycles"),
		"user":         os.ActionUsers(),
	})
}
