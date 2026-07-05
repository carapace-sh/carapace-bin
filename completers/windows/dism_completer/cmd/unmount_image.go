package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var UnmountImageCmd = &cobra.Command{
	Use:   "Unmount-Image",
	Short: "unmount a mounted image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(UnmountImageCmd).Standalone()
	rootCmd.AddCommand(UnmountImageCmd)
}
