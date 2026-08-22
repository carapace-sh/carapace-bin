package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/windows/scoop_completer/cmd/action"
	"github.com/spf13/cobra"
)

var bucket_addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a bucket",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bucket_addCmd).Standalone()
	bucketCmd.AddCommand(bucket_addCmd)

	carapace.Gen(bucket_addCmd).PositionalCompletion(
		action.ActionKnownBuckets(),
	)
}
