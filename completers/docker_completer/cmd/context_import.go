package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker_completer/cmd/action"
	"github.com/spf13/cobra"
)

var context_importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import a context from a tar or zip file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_importCmd).Standalone()

	contextCmd.AddCommand(context_importCmd)

	carapace.Gen(context_importCmd).PositionalCompletion(
		action.ActionContexts(),
		carapace.ActionFiles(),
	)
}
