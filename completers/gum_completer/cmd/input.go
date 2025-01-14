package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gum"
	"github.com/spf13/cobra"
)

var inputCmd = &cobra.Command{
	Use:   "input",
	Short: "Prompt for some input",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inputCmd).Standalone()

	inputCmd.Flags().String("char-limit", "", "Maximum value length (0 for no limit)")
	inputCmd.Flags().String("cursor.align", "", "Text Alignment")
	inputCmd.Flags().String("cursor.background", "", "Background Color")
	inputCmd.Flags().Bool("cursor.bold", false, "Bold text")
	inputCmd.Flags().String("cursor.border", "", "Border Style")
	inputCmd.Flags().String("cursor.border-background", "", "Border Background Color")
	inputCmd.Flags().String("cursor.border-foreground", "", "Border Foreground Color")
	inputCmd.Flags().Bool("cursor.faint", false, "Faint text")
	inputCmd.Flags().String("cursor.foreground", "", "Foreground Color")
	inputCmd.Flags().String("cursor.height", "", "Text height")
	inputCmd.Flags().Bool("cursor.italic", false, "Italicize text")
	inputCmd.Flags().String("cursor.margin", "", "Text margin")
	inputCmd.Flags().String("cursor.mode", "", "Cursor mode")
	inputCmd.Flags().String("cursor.padding", "", "Text padding")
	inputCmd.Flags().Bool("cursor.strikethrough", false, "Strikethrough text")
	inputCmd.Flags().Bool("cursor.underline", false, "Underline text")
	inputCmd.Flags().String("cursor.width", "", "Text width")
	inputCmd.Flags().String("header", "", "Header value")
	inputCmd.Flags().String("header.align", "", "Text Alignment")
	inputCmd.Flags().String("header.background", "", "Background Color")
	inputCmd.Flags().Bool("header.bold", false, "Bold text")
	inputCmd.Flags().String("header.border", "", "Border Style")
	inputCmd.Flags().String("header.border-background", "", "Border Background Color")
	inputCmd.Flags().String("header.border-foreground", "", "Border Foreground Color")
	inputCmd.Flags().Bool("header.faint", false, "Faint text")
	inputCmd.Flags().String("header.foreground", "", "Foreground Color")
	inputCmd.Flags().String("header.height", "", "Text height")
	inputCmd.Flags().Bool("header.italic", false, "Italicize text")
	inputCmd.Flags().String("header.margin", "", "Text margin")
	inputCmd.Flags().String("header.padding", "", "Text padding")
	inputCmd.Flags().Bool("header.strikethrough", false, "Strikethrough text")
	inputCmd.Flags().Bool("header.underline", false, "Underline text")
	inputCmd.Flags().String("header.width", "", "Text width")
	inputCmd.Flags().Bool("password", false, "Mask input characters")
	inputCmd.Flags().String("placeholder", "", "Placeholder value")
	inputCmd.Flags().String("placeholder.align", "", "Text Alignment")
	inputCmd.Flags().String("placeholder.background", "", "Background Color")
	inputCmd.Flags().Bool("placeholder.bold", false, "Bold text")
	inputCmd.Flags().String("placeholder.border", "", "Border Style")
	inputCmd.Flags().String("placeholder.border-background", "", "Border Background Color")
	inputCmd.Flags().String("placeholder.border-foreground", "", "Border Foreground Color")
	inputCmd.Flags().Bool("placeholder.faint", false, "Faint text")
	inputCmd.Flags().String("placeholder.foreground", "", "Foreground Color")
	inputCmd.Flags().String("placeholder.height", "", "Text height")
	inputCmd.Flags().Bool("placeholder.italic", false, "Italicize text")
	inputCmd.Flags().String("placeholder.margin", "", "Text margin")
	inputCmd.Flags().String("placeholder.padding", "", "Text padding")
	inputCmd.Flags().Bool("placeholder.strikethrough", false, "Strikethrough text")
	inputCmd.Flags().Bool("placeholder.underline", false, "Underline text")
	inputCmd.Flags().String("placeholder.width", "", "Text width")
	inputCmd.Flags().String("prompt", "", "Prompt to display")
	inputCmd.Flags().String("prompt.align", "", "Text Alignment")
	inputCmd.Flags().String("prompt.background", "", "Background Color")
	inputCmd.Flags().Bool("prompt.bold", false, "Bold text")
	inputCmd.Flags().String("prompt.border", "", "Border Style")
	inputCmd.Flags().String("prompt.border-background", "", "Border Background Color")
	inputCmd.Flags().String("prompt.border-foreground", "", "Border Foreground Color")
	inputCmd.Flags().Bool("prompt.faint", false, "Faint text")
	inputCmd.Flags().String("prompt.foreground", "", "Foreground Color")
	inputCmd.Flags().String("prompt.height", "", "Text height")
	inputCmd.Flags().Bool("prompt.italic", false, "Italicize text")
	inputCmd.Flags().String("prompt.margin", "", "Text margin")
	inputCmd.Flags().String("prompt.padding", "", "Text padding")
	inputCmd.Flags().Bool("prompt.strikethrough", false, "Strikethrough text")
	inputCmd.Flags().Bool("prompt.underline", false, "Underline text")
	inputCmd.Flags().String("prompt.width", "", "Text width")
	inputCmd.Flags().Bool("show-help", false, "Show help keybinds")
	inputCmd.Flags().Bool("strip-ansi", false, "Strip ANSI sequences when reading from STDIN")
	inputCmd.Flags().String("timeout", "", "Timeout until input aborts")
	inputCmd.Flags().String("value", "", "Initial value (can also be passed via stdin)")
	inputCmd.Flags().String("width", "", "Input width (0 for terminal width)")
	rootCmd.AddCommand(inputCmd)

	carapace.Gen(inputCmd).FlagCompletion(carapace.ActionMap{
		"cursor.align":                  gum.ActionAlignments(),
		"cursor.background":             gum.ActionColors(),
		"cursor.border":                 gum.ActionBorders(),
		"cursor.border-background":      gum.ActionColors(),
		"cursor.border-foreground":      gum.ActionColors(),
		"cursor.foreground":             gum.ActionColors(),
		"cursor.mode":                   gum.ActionCursorModes(),
		"header.align":                  gum.ActionAlignments(),
		"header.background":             gum.ActionColors(),
		"header.border":                 gum.ActionBorders(),
		"header.border-background":      gum.ActionColors(),
		"header.border-foreground":      gum.ActionColors(),
		"header.foreground":             gum.ActionColors(),
		"placeholder.align":             gum.ActionAlignments(),
		"placeholder.background":        gum.ActionColors(),
		"placeholder.border":            gum.ActionBorders(),
		"placeholder.border-background": gum.ActionColors(),
		"placeholder.border-foreground": gum.ActionColors(),
		"placeholder.foreground":        gum.ActionColors(),
		"prompt.align":                  gum.ActionAlignments(),
		"prompt.background":             gum.ActionColors(),
		"prompt.border":                 gum.ActionBorders(),
		"prompt.border-background":      gum.ActionColors(),
		"prompt.border-foreground":      gum.ActionColors(),
		"prompt.foreground":             gum.ActionColors(),
	})
}
