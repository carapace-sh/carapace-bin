package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var context_useCmd = &cobra.Command{
	Use:   "use CONTEXT",
	Short: "Set the current docker context",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_useCmd).Standalone()

	contextCmd.AddCommand(context_useCmd)

	carapace.Gen(context_useCmd).PositionalCompletion(
		docker.ActionContexts(),
	)
}
