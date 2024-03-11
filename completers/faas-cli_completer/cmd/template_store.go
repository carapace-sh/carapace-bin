package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var template_storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Command for pulling and listing templates from store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(template_storeCmd).Standalone()
	template_storeCmd.PersistentFlags().Bool("debug", false, "Enable debug output")
	template_storeCmd.PersistentFlags().Bool("overwrite", false, "Overwrite existing templates?")
	templateCmd.AddCommand(template_storeCmd)
}
