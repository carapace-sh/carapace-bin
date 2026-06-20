package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var explainConfigCmd = &cobra.Command{
	Use:   "+explain-config",
	Short: "Explain a config key or action",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(explainConfigCmd).Standalone()

	rootCmd.AddCommand(explainConfigCmd)
}
