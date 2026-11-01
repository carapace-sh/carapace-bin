package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Dump LaunchServices logging information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logCmd).Standalone()
	logCmd.Flags().Bool("B", false, "Log option")
	logCmd.Flags().Bool("a", false, "Log option")
	logCmd.Flags().Bool("c", false, "Log option")
	logCmd.Flags().Bool("d", false, "Log option")
	logCmd.Flags().Bool("e", false, "Log option")
	logCmd.Flags().Bool("i", false, "Log option")
	logCmd.Flags().Bool("n", false, "Log option")
	logCmd.Flags().String("sender", "", "Sender process name")
	logCmd.Flags().Bool("w", false, "Log option")
	rootCmd.AddCommand(logCmd)
}
