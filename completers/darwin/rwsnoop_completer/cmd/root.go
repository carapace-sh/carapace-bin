package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rwsnoop",
	Short: "snoop read/write events",
	Long:  "https://man.freebsd.org/cgi/man.cgi?rwsnoop",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("P", "P", false, "Print parent process ID")
	rootCmd.Flags().BoolS("Z", "Z", false, "Print zone ID")
	rootCmd.Flags().BoolS("j", "j", false, "Print project ID")
	rootCmd.Flags().StringS("n", "n", "", "Process name to track")
	rootCmd.Flags().StringS("p", "p", "", "PID to track")
	rootCmd.Flags().BoolS("t", "t", false, "Print timestamp")
	rootCmd.Flags().BoolS("v", "v", false, "Print time string")
}
