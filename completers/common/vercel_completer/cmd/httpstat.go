package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var httpstatCmd = &cobra.Command{
	Use:   "httpstat",
	Short: "Execute httpstat with automatic deployment URL and protection bypass",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(httpstatCmd).Standalone()

	httpstatCmd.Flags().String("deployment", "", "The deployment ID or URL to target")
	httpstatCmd.Flags().String("protection-bypass", "", "Protection bypass secret")
	httpstatCmd.Flags().BoolP("yes", "y", false, "Skip confirmation when linking is required")

	rootCmd.AddCommand(httpstatCmd)

	carapace.Gen(httpstatCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
