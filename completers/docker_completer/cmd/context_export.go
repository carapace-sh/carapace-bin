package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/docker_completer/cmd/action"
	"github.com/spf13/cobra"
)

var context_exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export a context to a tar or kubeconfig file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_exportCmd).Standalone()

	context_exportCmd.Flags().Bool("kubeconfig", false, "Export as a kubeconfig file")
	contextCmd.AddCommand(context_exportCmd)

	carapace.Gen(context_exportCmd).PositionalCompletion(
		action.ActionContexts(),
		carapace.ActionFiles(),
	)
}
