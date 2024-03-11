package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var context_exportCmd = &cobra.Command{
	Use:   "export [OPTIONS] CONTEXT [FILE|-]",
	Short: "Export a context to a tar archive FILE or a tar stream on STDOUT.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_exportCmd).Standalone()

	contextCmd.AddCommand(context_exportCmd)

	carapace.Gen(context_exportCmd).PositionalCompletion(
		docker.ActionContexts(),
		carapace.ActionFiles(),
	)
}
