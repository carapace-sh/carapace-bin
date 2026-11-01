package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ipcs",
	Short: "report System V interprocess communication facilities status",
	Long:  "https://man.freebsd.org/cgi/man.cgi?ipcs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("M", "M", false, "System information for shared memory")
	rootCmd.Flags().BoolS("P", "P", false, "System information for semaphores")
	rootCmd.Flags().BoolS("Q", "Q", false, "System information for message queues")
	rootCmd.Flags().BoolS("T", "T", false, "System information for all")
	rootCmd.Flags().BoolS("a", "a", false, "Use all print options")
	rootCmd.Flags().BoolS("b", "b", false, "Maximum allowable size")
	rootCmd.Flags().BoolS("c", "c", false, "Creator's login name and group name")
	rootCmd.Flags().BoolS("o", "o", false, "Outstanding usage")
	rootCmd.Flags().BoolS("p", "p", false, "Process number information")
	rootCmd.Flags().BoolS("t", "t", false, "Time information")
	rootCmd.Flags().BoolS("u", "u", false, "Utilization information")
}
