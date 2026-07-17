package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "Debug commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugCmd).Standalone()

	debugCmd.Flags().String("cpu-profile", "", "write a CPU profile to the specified file")
	debugCmd.Flags().String("file", "", "file name to log output to")
	debugCmd.Flags().String("mem-profile", "", "write a memory profile to the specified file")
	debugCmd.Flags().Int("profile-seconds", 0, "duration to profile for")
	rootCmd.AddCommand(debugCmd)
}
