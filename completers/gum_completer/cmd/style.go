package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gum"
	"github.com/spf13/cobra"
)

var styleCmd = &cobra.Command{
	Use:   "style",
	Short: "Apply coloring, borders, spacing to text",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(styleCmd).Standalone()

	styleCmd.Flags().String("align", "", "Text Alignment")
	styleCmd.Flags().String("background", "", "Background Color")
	styleCmd.Flags().Bool("bold", false, "Bold text")
	styleCmd.Flags().String("border", "", "Border Style")
	styleCmd.Flags().String("border-background", "", "Border Background Color")
	styleCmd.Flags().String("border-foreground", "", "Border Foreground Color")
	styleCmd.Flags().Bool("faint", false, "Faint text")
	styleCmd.Flags().String("foreground", "", "Foreground Color")
	styleCmd.Flags().String("height", "", "Text height")
	styleCmd.Flags().Bool("italic", false, "Italicize text")
	styleCmd.Flags().String("margin", "", "Text margin")
	styleCmd.Flags().String("padding", "", "Text padding")
	styleCmd.Flags().Bool("strikethrough", false, "Strikethrough text")
	styleCmd.Flags().Bool("strip-ansi", false, "Strip ANSI sequences when reading from STDIN")
	styleCmd.Flags().Bool("trim", false, "Trim whitespaces on every input line")
	styleCmd.Flags().Bool("underline", false, "Underline text")
	styleCmd.Flags().String("width", "", "Text width")
	rootCmd.AddCommand(styleCmd)

	carapace.Gen(styleCmd).FlagCompletion(carapace.ActionMap{
		"align":             gum.ActionAlignments(),
		"background":        gum.ActionColors(),
		"border":            gum.ActionBorders(),
		"border-background": gum.ActionColors(),
		"border-foreground": gum.ActionColors(),
		"foreground":        gum.ActionColors(),
	})
}
