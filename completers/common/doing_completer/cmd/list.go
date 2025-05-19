package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/doing"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list [OPTIONS]",
	Short: "List issues related to the project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().StringP("assignee", "a", "", "Filter by assignee")
	listCmd.Flags().StringP("author", "A", "", "Filter by author")
	listCmd.Flags().Bool("help", false, "Show this message and exit")
	listCmd.Flags().StringP("label", "l", "", "Filter by labels")
	listCmd.Flags().StringP("output_format", "o", "", "Output format")
	listCmd.Flags().StringP("state", "s", "", "Filter by state")
	listCmd.Flags().String("story_points", "", "Filter on number of story points")
	listCmd.Flags().StringP("type", "t", "", "Type of work item")
	rootCmd.AddCommand(listCmd)

	// TODO completion
	carapace.Gen(listCmd).FlagCompletion(carapace.ActionMap{
		"assignee":      carapace.ActionValues(),
		"author":        carapace.ActionValues(),
		"label":         carapace.ActionValues(),
		"output_format": carapace.ActionValues("table", "array"),
		"state":         carapace.ActionValues("open", "closed", "all").StyleF(style.ForKeyword), // TODO custom states
		"story_points":  carapace.ActionValues(),
		"type":          doing.ActionWorkItemTypes(),
	})
}
