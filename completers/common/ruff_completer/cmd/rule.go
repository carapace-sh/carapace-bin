package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ruleCmd = &cobra.Command{
	Use:   "rule",
	Short: "Explain a rule (or all rules)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ruleCmd).Standalone()

	ruleCmd.Flags().Bool("all", false, "Explain all rules")
	ruleCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	ruleCmd.Flags().String("output-format", "text", "Output format")
	rootCmd.AddCommand(ruleCmd)
}
