package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gum"
	"github.com/spf13/cobra"
)

var tableCmd = &cobra.Command{
	Use:   "table",
	Short: "Render a table of data",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tableCmd).Standalone()

	tableCmd.Flags().String("cell.align", "", "Text Alignment")
	tableCmd.Flags().String("cell.background", "", "Background Color")
	tableCmd.Flags().Bool("cell.bold", false, "Bold text")
	tableCmd.Flags().String("cell.border", "", "Border Style")
	tableCmd.Flags().String("cell.border-background", "", "Border Background Color")
	tableCmd.Flags().String("cell.border-foreground", "", "Border Foreground Color")
	tableCmd.Flags().Bool("cell.faint", false, "Faint text")
	tableCmd.Flags().String("cell.foreground", "", "Foreground Color")
	tableCmd.Flags().String("cell.height", "", "Text height")
	tableCmd.Flags().Bool("cell.italic", false, "Italicize text")
	tableCmd.Flags().String("cell.margin", "", "Text margin")
	tableCmd.Flags().String("cell.padding", "", "Text padding")
	tableCmd.Flags().Bool("cell.strikethrough", false, "Strikethrough text")
	tableCmd.Flags().Bool("cell.underline", false, "Underline text")
	tableCmd.Flags().String("cell.width", "", "Text width")
	tableCmd.Flags().StringSliceP("columns", "c", []string{}, "Column names")
	tableCmd.Flags().StringP("file", "f", "", "file path")
	tableCmd.Flags().String("header.align", "", "Text Alignment")
	tableCmd.Flags().String("header.background", "", "Background Color")
	tableCmd.Flags().Bool("header.bold", false, "Bold text")
	tableCmd.Flags().String("header.border", "", "Border Style")
	tableCmd.Flags().String("header.border-background", "", "Border Background Color")
	tableCmd.Flags().String("header.border-foreground", "", "Border Foreground Color")
	tableCmd.Flags().Bool("header.faint", false, "Faint text")
	tableCmd.Flags().String("header.foreground", "", "Foreground Color")
	tableCmd.Flags().String("header.height", "", "Text height")
	tableCmd.Flags().Bool("header.italic", false, "Italicize text")
	tableCmd.Flags().String("header.margin", "", "Text margin")
	tableCmd.Flags().String("header.padding", "", "Text padding")
	tableCmd.Flags().Bool("header.strikethrough", false, "Strikethrough text")
	tableCmd.Flags().Bool("header.underline", false, "Underline text")
	tableCmd.Flags().String("header.width", "", "Text width")
	tableCmd.Flags().String("height", "", "Table height")
	tableCmd.Flags().String("selected.align", "", "Text Alignment")
	tableCmd.Flags().String("selected.background", "", "Background Color")
	tableCmd.Flags().Bool("selected.bold", false, "Bold text")
	tableCmd.Flags().String("selected.border", "", "Border Style")
	tableCmd.Flags().String("selected.border-background", "", "Border Background Color")
	tableCmd.Flags().String("selected.border-foreground", "", "Border Foreground Color")
	tableCmd.Flags().Bool("selected.faint", false, "Faint text")
	tableCmd.Flags().String("selected.foreground", "", "Foreground Color")
	tableCmd.Flags().String("selected.height", "", "Text height")
	tableCmd.Flags().Bool("selected.italic", false, "Italicize text")
	tableCmd.Flags().String("selected.margin", "", "Text margin")
	tableCmd.Flags().String("selected.padding", "", "Text padding")
	tableCmd.Flags().Bool("selected.strikethrough", false, "Strikethrough text")
	tableCmd.Flags().Bool("selected.underline", false, "Underline text")
	tableCmd.Flags().String("selected.width", "", "Text width")
	tableCmd.Flags().StringP("separator", "s", "", "Row separator")
	tableCmd.Flags().StringSliceP("widths", "w", []string{}, "Column widths")
	rootCmd.AddCommand(tableCmd)

	carapace.Gen(tableCmd).FlagCompletion(carapace.ActionMap{
		"cell.background":            gum.ActionColors(),
		"cell.border":                gum.ActionBorders(),
		"cell.border-background":     gum.ActionColors(),
		"cell.border-foreground":     gum.ActionColors(),
		"cell.foreground":            gum.ActionColors(),
		"file":                       carapace.ActionFiles(),
		"header.align":               gum.ActionAlignments(),
		"header.background":          gum.ActionColors(),
		"header.border":              gum.ActionBorders(),
		"header.border-background":   gum.ActionColors(),
		"header.border-foreground":   gum.ActionColors(),
		"header.foreground":          gum.ActionColors(),
		"selected.align":             gum.ActionAlignments(),
		"selected.background":        gum.ActionColors(),
		"selected.border":            gum.ActionBorders(),
		"selected.border-background": gum.ActionColors(),
		"selected.border-foreground": gum.ActionColors(),
		"selected.foreground":        gum.ActionColors(),
	})
}
