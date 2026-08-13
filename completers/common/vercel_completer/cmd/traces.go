package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tracesCmd = &cobra.Command{
	Use:   "traces",
	Short: "Fetch traces captured for a Vercel project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tracesCmd).Standalone()

	tracesCmd.Flags().Bool("json", false, "Output as JSON")
	tracesCmd.Flags().Bool("open", false, "Open the trace in the browser")
	tracesCmd.Flags().String("project", "", "Project name or ID")
	tracesCmd.Flags().Bool("view", false, "View the trace")

	rootCmd.AddCommand(tracesCmd)

	carapace.Gen(tracesCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
