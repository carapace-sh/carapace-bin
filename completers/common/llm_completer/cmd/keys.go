package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keysCmd = &cobra.Command{
	Use:   "keys",
	Short: "Manage stored API keys for different models",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keysCmd).Standalone()

	rootCmd.AddCommand(keysCmd)
}
