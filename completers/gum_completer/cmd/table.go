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

	tableCmd.Flags().String("cell.foreground", "", "Foreground Color")
	tableCmd.Flags().StringP("columns", "c", "", "Column names")
	tableCmd.Flags().StringP("file", "f", "", "file path")
	tableCmd.Flags().String("header.foreground", "", "Foreground Color")
	tableCmd.Flags().String("height", "", "Table height")
	tableCmd.Flags().String("selected.foreground", "", "Foreground Color")
	tableCmd.Flags().StringP("separator", "s", "", "Row separator")
	tableCmd.Flags().StringP("widths", "w", "", "Column widths")
	rootCmd.AddCommand(tableCmd)

	carapace.Gen(tableCmd).FlagCompletion(carapace.ActionMap{
		"cell.foreground":     gum.ActionColors(),
		"file":                carapace.ActionFiles(),
		"header.foreground":   gum.ActionColors(),
		"selected.foreground": gum.ActionColors(),
	})
}
