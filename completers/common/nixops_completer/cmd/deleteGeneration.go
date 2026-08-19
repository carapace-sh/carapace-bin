package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deleteGenerationCmd = &cobra.Command{
	Use:   "delete-generation",
	Short: "delete a previous configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deleteGenerationCmd).Standalone()
	rootCmd.AddCommand(deleteGenerationCmd)
}
