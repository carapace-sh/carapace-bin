package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var experimentalCmd = &cobra.Command{
	Use:   "experimental",
	Short: "Experimental features (may change or be removed in the future)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(experimentalCmd).Standalone()

	rootCmd.AddCommand(experimentalCmd)
}
