package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var partitionTableCmd = &cobra.Command{
	Use:   "partition-table",
	Short: "Build only partition table",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(partitionTableCmd).Standalone()

	rootCmd.AddCommand(partitionTableCmd)
}
