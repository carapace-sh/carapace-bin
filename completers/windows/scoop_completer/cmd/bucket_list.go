package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bucket_listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all installed buckets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bucket_listCmd).Standalone()
	bucketCmd.AddCommand(bucket_listCmd)
}
