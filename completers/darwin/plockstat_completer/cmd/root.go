package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "plockstat",
	Short: "print statistics about POSIX mutexes and read/write locks",
	Long:  "https://man.freebsd.org/cgi/man.cgi?plockstat",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("A", "A", false, "Trace contention and hold events")
	rootCmd.Flags().BoolS("C", "C", false, "Trace contention events for mutexes and rwlocks")
	rootCmd.Flags().BoolS("H", "H", false, "Trace hold events for mutexes and rwlocks")
	rootCmd.Flags().BoolS("V", "V", false, "Print the dtrace script to run")
	rootCmd.Flags().StringS("e", "e", "", "Exit after specified seconds")
	rootCmd.Flags().StringS("n", "n", "", "Display only count entries for each event type")
	rootCmd.Flags().StringS("p", "p", "", "Attach and trace the specified process id")
	rootCmd.Flags().StringS("s", "s", "", "Show stack trace up to depth entries")
	rootCmd.Flags().BoolS("v", "v", false, "Print a message when tracing starts")
	rootCmd.Flags().StringS("x", "x", "", "Enable a DTrace runtime option")
}
