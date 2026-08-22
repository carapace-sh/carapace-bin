package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bucket_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove a bucket",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bucket_rmCmd).Standalone()
	bucketCmd.AddCommand(bucket_rmCmd)
}
