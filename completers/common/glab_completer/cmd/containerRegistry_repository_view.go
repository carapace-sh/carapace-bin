package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var containerRegistry_repository_viewCmd = &cobra.Command{
	Use:     "view <repository-id> [flags]",
	Short:   "View a container registry repository.",
	Aliases: []string{"show"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(containerRegistry_repository_viewCmd).Standalone()

	containerRegistry_repository_viewCmd.Flags().Bool("include-tags", false, "Include tags in the response.")
	containerRegistry_repository_viewCmd.Flags().Bool("include-tags-count", false, "Include the number of tags in the response.")
	containerRegistry_repository_viewCmd.Flags().String("jq", "", "Filter JSON output with a jq expression.")
	containerRegistry_repository_viewCmd.Flags().StringP("output", "F", "", "Format output as: text, json.")
	containerRegistry_repositoryCmd.AddCommand(containerRegistry_repository_viewCmd)

	carapace.Gen(containerRegistry_repository_viewCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("text", "json"),
	})
}
