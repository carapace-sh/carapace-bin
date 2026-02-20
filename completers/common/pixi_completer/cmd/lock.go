package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lockCmd = &cobra.Command{
	Use:   "lock",
	Short: "Solve environment and update the lock file without installing the environments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lockCmd).Standalone()

	lockCmd.Flags().Bool("check", false, "Check if any changes have been made to the lock file. If yes, exit with a non-zero code")
	lockCmd.Flags().Bool("json", false, "Output the changes in JSON format")
	lockCmd.Flags().Bool("no-install", false, "Don't modify the environment, only modify the lock-file")
	rootCmd.AddCommand(lockCmd)
}
