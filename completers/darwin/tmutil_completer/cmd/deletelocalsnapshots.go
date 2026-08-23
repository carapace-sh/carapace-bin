package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deletelocalsnapshotsCmd = &cobra.Command{
	Use:   "deletelocalsnapshots",
	Short: "delete local Time Machine snapshots",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deletelocalsnapshotsCmd).Standalone()
	rootCmd.AddCommand(deletelocalsnapshotsCmd)
}