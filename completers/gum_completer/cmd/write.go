package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gum"
	"github.com/spf13/cobra"
)

var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Prompt for long-form text",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(writeCmd).Standalone()

	writeCmd.Flags().String("base.align", "", "Text Alignment")
	writeCmd.Flags().String("base.background", "", "Background Color")
	writeCmd.Flags().Bool("base.bold", false, "Bold text")
	writeCmd.Flags().String("base.border", "", "Border Style")
	writeCmd.Flags().String("base.border-background", "", "Border Background Color")
	writeCmd.Flags().String("base.border-foreground", "", "Border Foreground Color")
	writeCmd.Flags().Bool("base.faint", false, "Faint text")
	writeCmd.Flags().String("base.foreground", "", "Foreground Color")
	writeCmd.Flags().String("base.height", "", "Text height")
	writeCmd.Flags().Bool("base.italic", false, "Italicize text")
	writeCmd.Flags().String("base.margin", "", "Text margin")
	writeCmd.Flags().String("base.padding", "", "Text padding")
	writeCmd.Flags().Bool("base.strikethrough", false, "Strikethrough text")
	writeCmd.Flags().Bool("base.underline", false, "Underline text")
	writeCmd.Flags().String("base.width", "", "Text width")
	writeCmd.Flags().String("char-limit", "", "Maximum value length (0 for no limit)")
	writeCmd.Flags().String("cursor-line-number.align", "", "Text Alignment")
	writeCmd.Flags().String("cursor-line-number.background", "", "Background Color")
	writeCmd.Flags().Bool("cursor-line-number.bold", false, "Bold text")
	writeCmd.Flags().String("cursor-line-number.border", "", "Border Style")
	writeCmd.Flags().String("cursor-line-number.border-background", "", "Border Background Color")
	writeCmd.Flags().String("cursor-line-number.border-foreground", "", "Border Foreground Color")
	writeCmd.Flags().Bool("cursor-line-number.faint", false, "Faint text")
	writeCmd.Flags().String("cursor-line-number.foreground", "", "Foreground Color")
	writeCmd.Flags().String("cursor-line-number.height", "", "Text height")
	writeCmd.Flags().Bool("cursor-line-number.italic", false, "Italicize text")
	writeCmd.Flags().String("cursor-line-number.margin", "", "Text margin")
	writeCmd.Flags().String("cursor-line-number.padding", "", "Text padding")
	writeCmd.Flags().Bool("cursor-line-number.strikethrough", false, "Strikethrough text")
	writeCmd.Flags().Bool("cursor-line-number.underline", false, "Underline text")
	writeCmd.Flags().String("cursor-line-number.width", "", "Text width")
	writeCmd.Flags().String("cursor-line.align", "", "Text Alignment")
	writeCmd.Flags().String("cursor-line.background", "", "Background Color")
	writeCmd.Flags().Bool("cursor-line.bold", false, "Bold text")
	writeCmd.Flags().String("cursor-line.border", "", "Border Style")
	writeCmd.Flags().String("cursor-line.border-background", "", "Border Background Color")
	writeCmd.Flags().String("cursor-line.border-foreground", "", "Border Foreground Color")
	writeCmd.Flags().Bool("cursor-line.faint", false, "Faint text")
	writeCmd.Flags().String("cursor-line.foreground", "", "Foreground Color")
	writeCmd.Flags().String("cursor-line.height", "", "Text height")
	writeCmd.Flags().Bool("cursor-line.italic", false, "Italicize text")
	writeCmd.Flags().String("cursor-line.margin", "", "Text margin")
	writeCmd.Flags().String("cursor-line.padding", "", "Text padding")
	writeCmd.Flags().Bool("cursor-line.strikethrough", false, "Strikethrough text")
	writeCmd.Flags().Bool("cursor-line.underline", false, "Underline text")
	writeCmd.Flags().String("cursor-line.width", "", "Text width")
	writeCmd.Flags().String("cursor.align", "", "Text Alignment")
	writeCmd.Flags().String("cursor.background", "", "Background Color")
	writeCmd.Flags().Bool("cursor.bold", false, "Bold text")
	writeCmd.Flags().String("cursor.border", "", "Border Style")
	writeCmd.Flags().String("cursor.border-background", "", "Border Background Color")
	writeCmd.Flags().String("cursor.border-foreground", "", "Border Foreground Color")
	writeCmd.Flags().Bool("cursor.faint", false, "Faint text")
	writeCmd.Flags().String("cursor.foreground", "", "Foreground Color")
	writeCmd.Flags().String("cursor.height", "", "Text height")
	writeCmd.Flags().Bool("cursor.italic", false, "Italicize text")
	writeCmd.Flags().String("cursor.margin", "", "Text margin")
	writeCmd.Flags().String("cursor.mode", "", "Cursor mode")
	writeCmd.Flags().String("cursor.padding", "", "Text padding")
	writeCmd.Flags().Bool("cursor.strikethrough", false, "Strikethrough text")
	writeCmd.Flags().Bool("cursor.underline", false, "Underline text")
	writeCmd.Flags().String("cursor.width", "", "Text width")
	writeCmd.Flags().String("end-of-buffer.align", "", "Text Alignment")
	writeCmd.Flags().String("end-of-buffer.background", "", "Background Color")
	writeCmd.Flags().Bool("end-of-buffer.bold", false, "Bold text")
	writeCmd.Flags().String("end-of-buffer.border", "", "Border Style")
	writeCmd.Flags().String("end-of-buffer.border-background", "", "Border Background Color")
	writeCmd.Flags().String("end-of-buffer.border-foreground", "", "Border Foreground Color")
	writeCmd.Flags().Bool("end-of-buffer.faint", false, "Faint text")
	writeCmd.Flags().String("end-of-buffer.foreground", "", "Foreground Color")
	writeCmd.Flags().String("end-of-buffer.height", "", "Text height")
	writeCmd.Flags().Bool("end-of-buffer.italic", false, "Italicize text")
	writeCmd.Flags().String("end-of-buffer.margin", "", "Text margin")
	writeCmd.Flags().String("end-of-buffer.padding", "", "Text padding")
	writeCmd.Flags().Bool("end-of-buffer.strikethrough", false, "Strikethrough text")
	writeCmd.Flags().Bool("end-of-buffer.underline", false, "Underline text")
	writeCmd.Flags().String("end-of-buffer.width", "", "Text width")
	writeCmd.Flags().String("header", "", "Header value")
	writeCmd.Flags().String("header.align", "", "Text Alignment")
	writeCmd.Flags().String("header.background", "", "Background Color")
	writeCmd.Flags().Bool("header.bold", false, "Bold text")
	writeCmd.Flags().String("header.border", "", "Border Style")
	writeCmd.Flags().String("header.border-background", "", "Border Background Color")
	writeCmd.Flags().String("header.border-foreground", "", "Border Foreground Color")
	writeCmd.Flags().Bool("header.faint", false, "Faint text")
	writeCmd.Flags().String("header.foreground", "", "Foreground Color")
	writeCmd.Flags().String("header.height", "", "Text height")
	writeCmd.Flags().Bool("header.italic", false, "Italicize text")
	writeCmd.Flags().String("header.margin", "", "Text margin")
	writeCmd.Flags().String("header.padding", "", "Text padding")
	writeCmd.Flags().Bool("header.strikethrough", false, "Strikethrough text")
	writeCmd.Flags().Bool("header.underline", false, "Underline text")
	writeCmd.Flags().String("header.width", "", "Text width")
	writeCmd.Flags().String("height", "", "Text area height")
	writeCmd.Flags().String("line-number.align", "", "Text Alignment")
	writeCmd.Flags().String("line-number.background", "", "Background Color")
	writeCmd.Flags().Bool("line-number.bold", false, "Bold text")
	writeCmd.Flags().String("line-number.border", "", "Border Style")
	writeCmd.Flags().String("line-number.border-background", "", "Border Background Color")
	writeCmd.Flags().String("line-number.border-foreground", "", "Border Foreground Color")
	writeCmd.Flags().Bool("line-number.faint", false, "Faint text")
	writeCmd.Flags().String("line-number.foreground", "", "Foreground Color")
	writeCmd.Flags().String("line-number.height", "", "Text height")
	writeCmd.Flags().Bool("line-number.italic", false, "Italicize text")
	writeCmd.Flags().String("line-number.margin", "", "Text margin")
	writeCmd.Flags().String("line-number.padding", "", "Text padding")
	writeCmd.Flags().Bool("line-number.strikethrough", false, "Strikethrough text")
	writeCmd.Flags().Bool("line-number.underline", false, "Underline text")
	writeCmd.Flags().String("line-number.width", "", "Text width")
	writeCmd.Flags().String("max-lines", "", "Maximum number of lines (0 for no limit)")
	writeCmd.Flags().String("placeholder", "", "Placeholder value")
	writeCmd.Flags().String("placeholder.align", "", "Text Alignment")
	writeCmd.Flags().String("placeholder.background", "", "Background Color")
	writeCmd.Flags().Bool("placeholder.bold", false, "Bold text")
	writeCmd.Flags().String("placeholder.border", "", "Border Style")
	writeCmd.Flags().String("placeholder.border-background", "", "Border Background Color")
	writeCmd.Flags().String("placeholder.border-foreground", "", "Border Foreground Color")
	writeCmd.Flags().Bool("placeholder.faint", false, "Faint text")
	writeCmd.Flags().String("placeholder.foreground", "", "Foreground Color")
	writeCmd.Flags().String("placeholder.height", "", "Text height")
	writeCmd.Flags().Bool("placeholder.italic", false, "Italicize text")
	writeCmd.Flags().String("placeholder.margin", "", "Text margin")
	writeCmd.Flags().String("placeholder.padding", "", "Text padding")
	writeCmd.Flags().Bool("placeholder.strikethrough", false, "Strikethrough text")
	writeCmd.Flags().Bool("placeholder.underline", false, "Underline text")
	writeCmd.Flags().String("placeholder.width", "", "Text width")
	writeCmd.Flags().String("prompt", "", "Prompt to display")
	writeCmd.Flags().String("prompt.align", "", "Text Alignment")
	writeCmd.Flags().String("prompt.background", "", "Background Color")
	writeCmd.Flags().Bool("prompt.bold", false, "Bold text")
	writeCmd.Flags().String("prompt.border", "", "Border Style")
	writeCmd.Flags().String("prompt.border-background", "", "Border Background Color")
	writeCmd.Flags().String("prompt.border-foreground", "", "Border Foreground Color")
	writeCmd.Flags().Bool("prompt.faint", false, "Faint text")
	writeCmd.Flags().String("prompt.foreground", "", "Foreground Color")
	writeCmd.Flags().String("prompt.height", "", "Text height")
	writeCmd.Flags().Bool("prompt.italic", false, "Italicize text")
	writeCmd.Flags().String("prompt.margin", "", "Text margin")
	writeCmd.Flags().String("prompt.padding", "", "Text padding")
	writeCmd.Flags().Bool("prompt.strikethrough", false, "Strikethrough text")
	writeCmd.Flags().Bool("prompt.underline", false, "Underline text")
	writeCmd.Flags().String("prompt.width", "", "Text width")
	writeCmd.Flags().Bool("show-cursor-line", false, "Show cursor line")
	writeCmd.Flags().Bool("show-help", false, "Show help key binds")
	writeCmd.Flags().Bool("show-line-numbers", false, "Show line numbers")
	writeCmd.Flags().Bool("strip-ansi", false, "Strip ANSI sequences when reading from STDIN")
	writeCmd.Flags().String("timeout", "", "Timeout until choose returns selected element")
	writeCmd.Flags().String("value", "", "Initial value (can be passed via stdin)")
	writeCmd.Flags().String("width", "", "Text area width (0 for terminal width)")
	rootCmd.AddCommand(writeCmd)

	carapace.Gen(writeCmd).FlagCompletion(carapace.ActionMap{
		"base.align":                           gum.ActionAlignments(),
		"base.background":                      gum.ActionColors(),
		"base.border":                          gum.ActionBorders(),
		"base.border-background":               gum.ActionColors(),
		"base.border-foreground":               gum.ActionColors(),
		"base.foreground":                      gum.ActionColors(),
		"cursor-line-number.align":             gum.ActionAlignments(),
		"cursor-line-number.background":        gum.ActionColors(),
		"cursor-line-number.border":            gum.ActionBorders(),
		"cursor-line-number.border-background": gum.ActionColors(),
		"cursor-line-number.border-foreground": gum.ActionColors(),
		"cursor-line-number.foreground":        gum.ActionColors(),
		"cursor-line.align":                    gum.ActionAlignments(),
		"cursor-line.background":               gum.ActionColors(),
		"cursor-line.border":                   gum.ActionBorders(),
		"cursor-line.border-background":        gum.ActionColors(),
		"cursor-line.border-foreground":        gum.ActionColors(),
		"cursor-line.foreground":               gum.ActionColors(),
		"cursor.align":                         gum.ActionAlignments(),
		"cursor.background":                    gum.ActionColors(),
		"cursor.border":                        gum.ActionBorders(),
		"cursor.border-background":             gum.ActionColors(),
		"cursor.border-foreground":             gum.ActionColors(),
		"cursor.foreground":                    gum.ActionColors(),
		"cursor.mode":                          gum.ActionCursorModes(),
		"end-of-buffer.align":                  gum.ActionAlignments(),
		"end-of-buffer.background":             gum.ActionColors(),
		"end-of-buffer.border":                 gum.ActionBorders(),
		"end-of-buffer.border-background":      gum.ActionColors(),
		"end-of-buffer.border-foreground":      gum.ActionColors(),
		"end-of-buffer.foreground":             gum.ActionColors(),
		"header.align":                         gum.ActionAlignments(),
		"header.background":                    gum.ActionColors(),
		"header.border":                        gum.ActionBorders(),
		"header.border-background":             gum.ActionColors(),
		"header.border-foreground":             gum.ActionColors(),
		"header.foreground":                    gum.ActionColors(),
		"line-number.align":                    gum.ActionAlignments(),
		"line-number.background":               gum.ActionColors(),
		"line-number.border":                   gum.ActionBorders(),
		"line-number.border-background":        gum.ActionColors(),
		"line-number.border-foreground":        gum.ActionColors(),
		"line-number.foreground":               gum.ActionColors(),
		"placeholder.align":                    gum.ActionAlignments(),
		"placeholder.background":               gum.ActionColors(),
		"placeholder.border":                   gum.ActionBorders(),
		"placeholder.border-background":        gum.ActionColors(),
		"placeholder.border-foreground":        gum.ActionColors(),
		"placeholder.foreground":               gum.ActionColors(),
		"prompt.align":                         gum.ActionAlignments(),
		"prompt.background":                    gum.ActionColors(),
		"prompt.border":                        gum.ActionBorders(),
		"prompt.border-background":             gum.ActionColors(),
		"prompt.border-foreground":             gum.ActionColors(),
		"prompt.foreground":                    gum.ActionColors(),
	})
}
