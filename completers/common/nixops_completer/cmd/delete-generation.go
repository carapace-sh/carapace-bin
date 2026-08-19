package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var DeleteGenerationCmd = &cobra.Command{
	Use:   "delete-generation",
	Short: "Delete Generation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(DeleteGenerationCmd).Standalone()
	rootCmd.AddCommand(DeleteGenerationCmd)
}
