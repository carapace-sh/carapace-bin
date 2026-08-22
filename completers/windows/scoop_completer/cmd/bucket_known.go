package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bucket_knownCmd = &cobra.Command{
	Use:   "known",
	Short: "list all known buckets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bucket_knownCmd).Standalone()
	bucketCmd.AddCommand(bucket_knownCmd)
}
