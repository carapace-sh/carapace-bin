package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var themeChooseCmd = &cobra.Command{
	Use:   "choose",
	Short: "Load a color theme",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(themeChooseCmd).Standalone()

	themeChooseCmd.Flags().String("color-theme", "", "override fish_terminal_color_theme")

	themeCmd.AddCommand(themeChooseCmd)

	carapace.Gen(themeChooseCmd).FlagCompletion(carapace.ActionMap{
		"color-theme": carapace.ActionValues("dark", "light"),
	})
}
