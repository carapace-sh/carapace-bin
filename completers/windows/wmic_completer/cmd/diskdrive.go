package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var diskdriveCmd = &cobra.Command{
	Use:   "diskdrive",
	Short: "disk drive management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diskdriveCmd).Standalone()
	rootCmd.AddCommand(diskdriveCmd)
}
