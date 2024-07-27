package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var makeRuleCmd = &cobra.Command{
	Use:   "make:rule [-f|--force] [-i|--implicit] [--] <name>",
	Short: "Create a new validation rule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(makeRuleCmd).Standalone()

	makeRuleCmd.Flags().Bool("force", false, "Create the class even if the rule already exists")
	makeRuleCmd.Flags().Bool("implicit", false, "Generate an implicit rule")
	rootCmd.AddCommand(makeRuleCmd)
}
