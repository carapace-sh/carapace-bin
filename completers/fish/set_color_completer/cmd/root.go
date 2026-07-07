package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/fish"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "set_color",
	Short: "Set terminal colors",
	Long:  "https://fishshell.com/docs/current/cmds/set_color.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("b", "b", "", "background color")
	rootCmd.Flags().String("background", "", "background color")
	rootCmd.Flags().Bool("bold", false, "bold mode")
	rootCmd.Flags().BoolS("c", "c", false, "print named colors")
	rootCmd.Flags().BoolS("d", "d", false, "dim mode")
	rootCmd.Flags().Bool("dim", false, "dim mode")
	rootCmd.Flags().StringS("f", "f", "", "foreground color")
	rootCmd.Flags().String("foreground", "", "foreground color")
	rootCmd.Flags().BoolS("h", "h", false, "display help")
	rootCmd.Flags().StringS("i", "i", "", "italics")
	rootCmd.Flags().BoolS("o", "o", false, "bold mode")
	rootCmd.Flags().Bool("print-colors", false, "print named colors")
	rootCmd.Flags().StringS("r", "r", "", "reverse mode")
	rootCmd.Flags().Bool("reset", false, "reset formatting")
	rootCmd.Flags().StringS("s", "s", "", "strikethrough")
	rootCmd.Flags().String("theme", "", "theme name")
	rootCmd.Flags().StringS("u", "u", "", "underline style")
	rootCmd.Flags().String("underline-color", "", "underline color")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"b":               fish.ActionColorNames(),
		"background":      fish.ActionColorNames(),
		"f":               fish.ActionColorNames(),
		"foreground":      fish.ActionColorNames(),
		"u":               carapace.ActionValues("single", "double", "curly", "dotted", "dashed", "off"),
		"underline-color": fish.ActionColorNames(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		fish.ActionColorNames(),
	)
}
