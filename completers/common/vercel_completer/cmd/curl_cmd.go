package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var curlCmd = &cobra.Command{
	Use:   "curl",
	Short: "Make curl requests to Vercel deployments with automatic protection bypass",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(curlCmd).Standalone()

	curlCmd.Flags().String("deployment", "", "The deployment ID or URL to target")
	curlCmd.Flags().Bool("json", false, "With --trace, emit JSON")
	curlCmd.Flags().String("protection-bypass", "", "Protection bypass secret")
	curlCmd.Flags().String("trace", "", "Capture a session trace for the request")
	curlCmd.Flags().BoolP("yes", "y", false, "Skip confirmation when linking is required")

	rootCmd.AddCommand(curlCmd)

	carapace.Gen(curlCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
