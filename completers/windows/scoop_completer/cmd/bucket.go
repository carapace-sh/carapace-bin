package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bucketCmd = &cobra.Command{
	Use:   "bucket",
	Short: "manage buckets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bucketCmd).Standalone()
	rootCmd.AddCommand(bucketCmd)
}
