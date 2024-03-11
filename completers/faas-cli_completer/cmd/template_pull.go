package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var template_pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Downloads templates from the specified git repo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(template_pullCmd).Standalone()
	template_pullCmd.Flags().Bool("debug", false, "Enable debug output")
	template_pullCmd.Flags().Bool("overwrite", false, "Overwrite existing templates?")
	templateCmd.AddCommand(template_pullCmd)
}
