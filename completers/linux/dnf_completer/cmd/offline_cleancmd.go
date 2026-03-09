package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var offline_cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "remove stored offline transaction and delete cached package files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(offline_cleanCmd).Standalone()

	offlineCmd.AddCommand(offline_cleanCmd)
}
