package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var linterCmd = &cobra.Command{
	Use:   "linter",
	Short: "List all supported upstream linters",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(linterCmd).Standalone()

	linterCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	linterCmd.Flags().String("output-format", "text", "Output format")
	rootCmd.AddCommand(linterCmd)
}
