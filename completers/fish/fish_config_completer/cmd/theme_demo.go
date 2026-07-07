package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var themeDemoCmd = &cobra.Command{
	Use:   "demo",
	Short: "Display sample text with current theme",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(themeDemoCmd).Standalone()

	themeCmd.AddCommand(themeDemoCmd)
}
