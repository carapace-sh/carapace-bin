package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "taskpolicy",
	Short: "execute a program with altered I/O or scheduling policy",
	Long:  "https://keith.github.io/xcode-manpages/taskpolicy.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("B", "B", false, "Background")
	rootCmd.Flags().StringS("S", "S", "", "Shims")
	rootCmd.Flags().BoolS("a", "a", false, "All")
	rootCmd.Flags().BoolS("b", "b", false, "Background")
	rootCmd.Flags().StringS("c", "c", "", "QoS clamp")
	rootCmd.Flags().StringS("d", "d", "", "Disk I/O policy")
	rootCmd.Flags().StringS("g", "g", "", "Darwin BG policy")
	rootCmd.Flags().StringS("j", "j", "", "Jetsam priority")
	rootCmd.Flags().StringS("l", "l", "", "Latency tier")
	rootCmd.Flags().StringS("m", "m", "", "Memory limit (MiB)")
	rootCmd.Flags().StringS("p", "p", "", "PID")
	rootCmd.Flags().BoolS("s", "s", false, "Set")
	rootCmd.Flags().StringS("t", "t", "", "Throughput tier")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"c": carapace.ActionValues("utility", "background", "maintenance"),
		"d": carapace.ActionValues("default", "throttle"),
		"g": carapace.ActionValues("default", "throttle"),
		"p": ps.ActionProcessIds(),
	})
}
