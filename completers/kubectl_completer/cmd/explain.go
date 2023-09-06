package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var explainCmd = &cobra.Command{
	Use:     "explain RESOURCE",
	Short:   "Get documentation for a resource",
	GroupID: "basic intermediate",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(explainCmd).Standalone()

	explainCmd.Flags().String("api-version", "", "Use given api-version (group/version) of the resource.")
	explainCmd.Flags().String("output", "", "Format in which to render the schema. Valid values are: (plaintext, plaintext-openapiv2).")
	explainCmd.Flags().Bool("recursive", false, "When true, print the name of all the fields recursively. Otherwise, print the available fields with their description.")
	rootCmd.AddCommand(explainCmd)

	// TODO api-version
	carapace.Gen(explainCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("plaintext", "plaintext-openapiv2"),
	})

	carapace.Gen(explainCmd).PositionalCompletion(
		kubectl.ActionApiResources(),
	)
}
