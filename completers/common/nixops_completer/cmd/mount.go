package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var MountCmd = &cobra.Command{
	Use:   "mount",
	Short: "Mount",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(MountCmd).Standalone()
	rootCmd.AddCommand(MountCmd)
}
