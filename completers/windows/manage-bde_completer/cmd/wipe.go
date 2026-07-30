package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wipeCmd = &cobra.Command{
	Use:   "wipe",
	Short: "wipe the free space on a drive",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wipeCmd).Standalone()
	rootCmd.AddCommand(wipeCmd)
}
