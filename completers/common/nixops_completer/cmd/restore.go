package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "restore machines from snapshots",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(restoreCmd).Standalone()

	restoreCmd.Flags().String("backup-id", "", "Use specified backup instead of latest")
	restoreCmd.Flags().StringSlice("devices", nil, "Only restore the specified devices")
	restoreCmd.Flags().StringSlice("exclude", nil, "Do not perform backup actions on the specified machines")
	restoreCmd.Flags().StringSlice("include", nil, "Perform backup actions on the specified machines only")

	rootCmd.AddCommand(restoreCmd)
}
