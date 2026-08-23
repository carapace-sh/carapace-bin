package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listlocalsnapshotsCmd = &cobra.Command{
	Use:   "listlocalsnapshots",
	Short: "list local Time Machine snapshots",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listlocalsnapshotsCmd).Standalone()
	rootCmd.AddCommand(listlocalsnapshotsCmd)
}
