package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sync_nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Sync node installed packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sync_nodeCmd).Standalone()

	syncCmd.AddCommand(sync_nodeCmd)
}
