package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var themesCmd = &cobra.Command{
	Use:   "themes",
	Short: "Lists available themes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(themesCmd).Standalone()

	rootCmd.AddCommand(themesCmd)
}
