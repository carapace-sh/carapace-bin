package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var template_pull_stackCmd = &cobra.Command{
	Use:   "stack",
	Short: "Downloads templates specified in the function definition yaml file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(template_pull_stackCmd).Standalone()
	template_pull_stackCmd.Flags().Bool("debug", false, "Enable debug output")
	template_pull_stackCmd.Flags().Bool("overwrite", false, "Overwrite existing templates?")
	template_pullCmd.AddCommand(template_pull_stackCmd)
}
