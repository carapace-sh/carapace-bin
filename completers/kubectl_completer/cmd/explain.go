package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var explainCmd = &cobra.Command{
	Use:   "explain",
	Short: "Documentation of resources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(explainCmd).Standalone()

	explainCmd.Flags().String("api-version", "", "get different explanations for particular API version")
	explainCmd.Flags().Bool("recursive", false, "print the fields of fields")
	rootCmd.AddCommand(explainCmd)

	carapace.Gen(explainCmd).PositionalCompletion(
		action.ActionApiResources(),
	)
}
