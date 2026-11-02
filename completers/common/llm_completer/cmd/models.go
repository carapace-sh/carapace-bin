package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var modelsCmd = &cobra.Command{
	Use:   "models",
	Short: "Manage available models",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(modelsCmd).Standalone()

	rootCmd.AddCommand(modelsCmd)
}
