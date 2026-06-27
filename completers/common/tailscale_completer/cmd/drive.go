package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var driveCmd = &cobra.Command{
	Use:   "drive",
	Short: "Share a directory with your tailnet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(driveCmd).Standalone()

	rootCmd.AddCommand(driveCmd)
}
