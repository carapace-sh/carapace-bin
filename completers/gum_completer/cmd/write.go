package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gum"
	"github.com/spf13/cobra"
)

var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Prompt for long-form text",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(writeCmd).Standalone()

	writeCmd.Flags().String("base.foreground", "", "Foreground Color")
	writeCmd.Flags().String("char-limit", "", "Maximum value length")
	writeCmd.Flags().String("cursor-line-number.foreground", "", "Foreground Color")
	writeCmd.Flags().String("cursor-line.foreground", "", "Foreground Color")
	writeCmd.Flags().String("cursor.foreground", "", "Foreground Color")
	writeCmd.Flags().String("end-of-buffer.foreground", "", "Foreground Color")
	writeCmd.Flags().String("height", "", "Text area height")
	writeCmd.Flags().String("line-number.foreground", "", "Foreground Color")
	writeCmd.Flags().String("placeholder", "", "Placeholder value")
	writeCmd.Flags().String("placeholder.foreground", "", "Foreground Color")
	writeCmd.Flags().String("prompt", "", "Prompt to display")
	writeCmd.Flags().String("prompt.foreground", "", "Foreground Color")
	writeCmd.Flags().Bool("show-cursor-line", false, "Show cursor line")
	writeCmd.Flags().Bool("show-line-numbers", false, "Show line numbers")
	writeCmd.Flags().String("value", "", "Initial value")
	writeCmd.Flags().String("width", "", "Text area width")
	rootCmd.AddCommand(writeCmd)

	carapace.Gen(writeCmd).FlagCompletion(carapace.ActionMap{
		"base.foreground":               gum.ActionColors(),
		"cursor-line-number.foreground": gum.ActionColors(),
		"cursor-line.foreground":        gum.ActionColors(),
		"end-of-buffer.foreground":      gum.ActionColors(),
		"line-number.foreground":        gum.ActionColors(),
		"placeholder.foreground":        gum.ActionColors(),
		"prompt.foreground":             gum.ActionColors(),
	})
}
