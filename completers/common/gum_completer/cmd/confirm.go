package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gum_completer/cmd/common"
	"github.com/spf13/cobra"
)

var confirmCmd = &cobra.Command{
	Use:   "confirm",
	Short: "Ask a user to confirm an action",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(confirmCmd).Standalone()

	confirmCmd.Flags().String("affirmative", "", "The title of the affirmative action")
	confirmCmd.Flags().Bool("default", false, "Default confirmation action")
	confirmCmd.Flags().String("negative", "", "The title of the negative action")
	confirmCmd.Flags().String("prompt.align", "", "Text Alignment")
	confirmCmd.Flags().String("prompt.background", "", "Background Color")
	confirmCmd.Flags().Bool("prompt.bold", false, "Bold text")
	confirmCmd.Flags().String("prompt.border", "", "Border Style")
	confirmCmd.Flags().String("prompt.border-background", "", "Border Background Color")
	confirmCmd.Flags().String("prompt.border-foreground", "", "Border Foreground Color")
	confirmCmd.Flags().Bool("prompt.faint", false, "Faint text")
	confirmCmd.Flags().String("prompt.foreground", "", "Foreground Color")
	confirmCmd.Flags().String("prompt.height", "", "Text height")
	confirmCmd.Flags().Bool("prompt.italic", false, "Italicize text")
	confirmCmd.Flags().String("prompt.margin", "", "Text margin")
	confirmCmd.Flags().String("prompt.padding", "", "Text padding")
	confirmCmd.Flags().Bool("prompt.strikethrough", false, "Strikethrough text")
	confirmCmd.Flags().Bool("prompt.underline", false, "Underline text")
	confirmCmd.Flags().String("prompt.width", "", "Text width")
	confirmCmd.Flags().String("selected.align", "", "Text Alignment")
	confirmCmd.Flags().String("selected.background", "", "Background Color")
	confirmCmd.Flags().Bool("selected.bold", false, "Bold text")
	confirmCmd.Flags().String("selected.border", "", "Border Style")
	confirmCmd.Flags().String("selected.border-background", "", "Border Background Color")
	confirmCmd.Flags().String("selected.border-foreground", "", "Border Foreground Color")
	confirmCmd.Flags().Bool("selected.faint", false, "Faint text")
	confirmCmd.Flags().String("selected.foreground", "", "Foreground Color")
	confirmCmd.Flags().String("selected.height", "", "Text height")
	confirmCmd.Flags().Bool("selected.italic", false, "Italicize text")
	confirmCmd.Flags().String("selected.margin", "", "Text margin")
	confirmCmd.Flags().String("selected.padding", "", "Text padding")
	confirmCmd.Flags().Bool("selected.strikethrough", false, "Strikethrough text")
	confirmCmd.Flags().Bool("selected.underline", false, "Underline text")
	confirmCmd.Flags().String("selected.width", "", "Text width")
	confirmCmd.Flags().Bool("show-help", false, "Show help key binds")
	confirmCmd.Flags().Bool("show-output", false, "Print prompt and chosen action to output")
	confirmCmd.Flags().String("timeout", "", "Timeout until confirm returns selected value or default if provided")
	confirmCmd.Flags().String("unselected.align", "", "Text Alignment")
	confirmCmd.Flags().String("unselected.background", "", "Background Color")
	confirmCmd.Flags().Bool("unselected.bold", false, "Bold text")
	confirmCmd.Flags().String("unselected.border", "", "Border Style")
	confirmCmd.Flags().String("unselected.border-background", "", "Border Background Color")
	confirmCmd.Flags().String("unselected.border-foreground", "", "Border Foreground Color")
	confirmCmd.Flags().Bool("unselected.faint", false, "Faint text")
	confirmCmd.Flags().String("unselected.foreground", "", "Foreground Color")
	confirmCmd.Flags().String("unselected.height", "", "Text height")
	confirmCmd.Flags().Bool("unselected.italic", false, "Italicize text")
	confirmCmd.Flags().String("unselected.margin", "", "Text margin")
	confirmCmd.Flags().String("unselected.padding", "", "Text padding")
	confirmCmd.Flags().Bool("unselected.strikethrough", false, "Strikethrough text")
	confirmCmd.Flags().Bool("unselected.underline", false, "Underline text")
	confirmCmd.Flags().String("unselected.width", "", "Text width")
	rootCmd.AddCommand(confirmCmd)

	common.AddFlagCompletion(confirmCmd)
}
