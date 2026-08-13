package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/vercel_completer/cmd/action"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"list"},
	Short:   "List deployments",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lsCmd).Standalone()

	lsCmd.Flags().BoolP("all", "a", false, "List resources across all projects")
	lsCmd.Flags().BoolP("confirm", "c", false, "(deprecated)")
	lsCmd.Flags().String("environment", "", "Filter by target environment")
	lsCmd.Flags().StringP("format", "F", "", "Output format")
	lsCmd.Flags().Bool("json", false, "Output as JSON")
	lsCmd.Flags().String("limit", "", "Results per page")
	lsCmd.Flags().StringSliceP("meta", "m", nil, "Filter deployments by metadata")
	lsCmd.Flags().StringP("next", "N", "", "Show next page of results")
	lsCmd.Flags().StringP("policy", "p", "", "Filter by Deployment Retention policies")
	lsCmd.Flags().Bool("prod", false, "Only show production deployments")
	lsCmd.Flags().StringP("status", "s", "", "Filter by status (comma-separated)")
	lsCmd.Flags().BoolP("yes", "y", false, "Skip confirmation")

	rootCmd.AddCommand(lsCmd)

	carapace.Gen(lsCmd).PositionalCompletion(
		action.ActionProjects(lsCmd),
	)
}
