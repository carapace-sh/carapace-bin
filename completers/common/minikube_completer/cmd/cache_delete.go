package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/minikube_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cache_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an image from the local cache.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cache_deleteCmd).Standalone()
	cacheCmd.AddCommand(cache_deleteCmd)

	carapace.Gen(cache_deleteCmd).PositionalCompletion(
		action.ActionCachedImages(),
	)
}
