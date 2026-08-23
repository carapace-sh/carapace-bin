package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var thinlocalsnapshotsCmd = &cobra.Command{
	Use:   "thinlocalsnapshots",
	Short: "thin local Time Machine snapshots",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(thinlocalsnapshotsCmd).Standalone()
	rootCmd.AddCommand(thinlocalsnapshotsCmd)
}