package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var themeSaveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save theme to universal variables",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(themeSaveCmd).Standalone()

	themeCmd.AddCommand(themeSaveCmd)
}
