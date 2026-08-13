package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/vercel_completer/cmd/action"
	"github.com/spf13/cobra"
)

var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Show information about a deployment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspectCmd).Standalone()

	inspectCmd.Flags().StringP("format", "F", "", "Output format")
	inspectCmd.Flags().Bool("json", false, "Output as JSON")
	inspectCmd.Flags().BoolP("logs", "l", false, "Prints the build logs")
	inspectCmd.Flags().String("timeout", "", "Time to wait for deployment completion [3m]")
	inspectCmd.Flags().Bool("wait", false, "Blocks until deployment completes")

	rootCmd.AddCommand(inspectCmd)

	carapace.Gen(inspectCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("plain", "json"),
	})

	carapace.Gen(inspectCmd).PositionalCompletion(
		action.ActionDeployments(inspectCmd),
	)
}
