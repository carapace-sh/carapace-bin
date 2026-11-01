package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var GetMountedImageInfoCmd = &cobra.Command{
	Use:   "Get-MountedImageInfo",
	Short: "list currently mounted images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(GetMountedImageInfoCmd).Standalone()
	rootCmd.AddCommand(GetMountedImageInfoCmd)
}
