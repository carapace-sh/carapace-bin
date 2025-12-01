package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var explainCmd = &cobra.Command{
	Use:     "explain TYPE [--recursive=FALSE|TRUE] [--api-version=api-version-group] [--output=plaintext|plaintext-openapiv2]",
	Short:   "Get documentation for a resource",
	GroupID: "basic intermediate",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(explainCmd).Standalone()

	explainCmd.Flags().String("api-version", "", "Get different explanations for particular API version (API group/version)")
	explainCmd.Flags().StringP("output", "o", "", "Format in which to render the schema (plaintext, plaintext-openapiv2)")
	explainCmd.Flags().Bool("recursive", false, "Print the fields of fields (Currently only 1 level deep)")
	rootCmd.AddCommand(explainCmd)

	// TODO api-version
	carapace.Gen(explainCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("plaintext", "plaintext-openapiv2"),
	})

	carapace.Gen(explainCmd).PositionalCompletion(
		kubectl.ActionApiResources(),
	)
}
