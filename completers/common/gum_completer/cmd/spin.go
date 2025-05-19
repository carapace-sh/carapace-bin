package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gum_completer/cmd/common"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var spinCmd = &cobra.Command{
	Use:   "spin",
	Short: "Display spinner while running a command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(spinCmd).Standalone()
	spinCmd.Flags().SetInterspersed(false)

	spinCmd.Flags().StringP("align", "a", "", "Alignment of spinner with regard to the title")
	spinCmd.Flags().Bool("show-error", false, "Show output of command only if the command fails")
	spinCmd.Flags().Bool("show-output", false, "Show or pipe output of command during execution (shows both STDOUT and STDERR)")
	spinCmd.Flags().Bool("show-stderr", false, "Show STDERR errput")
	spinCmd.Flags().Bool("show-stdout", false, "Show STDOUT output")
	spinCmd.Flags().StringP("spinner", "s", "", "Spinner type")
	spinCmd.Flags().String("spinner.align", "", "Text Alignment")
	spinCmd.Flags().String("spinner.background", "", "Background Color")
	spinCmd.Flags().Bool("spinner.bold", false, "Bold text")
	spinCmd.Flags().String("spinner.border", "", "Border Style")
	spinCmd.Flags().String("spinner.border-background", "", "Border Background Color")
	spinCmd.Flags().String("spinner.border-foreground", "", "Border Foreground Color")
	spinCmd.Flags().Bool("spinner.faint", false, "Faint text")
	spinCmd.Flags().String("spinner.foreground", "", "Foreground Color")
	spinCmd.Flags().String("spinner.height", "", "Text height")
	spinCmd.Flags().Bool("spinner.italic", false, "Italicize text")
	spinCmd.Flags().String("spinner.margin", "", "Text margin")
	spinCmd.Flags().String("spinner.padding", "", "Text padding")
	spinCmd.Flags().Bool("spinner.strikethrough", false, "Strikethrough text")
	spinCmd.Flags().Bool("spinner.underline", false, "Underline text")
	spinCmd.Flags().String("spinner.width", "", "Text width")
	spinCmd.Flags().String("timeout", "", "Timeout until spin command aborts")
	spinCmd.Flags().String("title", "", "Text to display to user while spinning")
	spinCmd.Flags().String("title.align", "", "Text Alignment")
	spinCmd.Flags().String("title.background", "", "Background Color")
	spinCmd.Flags().Bool("title.bold", false, "Bold text")
	spinCmd.Flags().String("title.border", "", "Border Style")
	spinCmd.Flags().String("title.border-background", "", "Border Background Color")
	spinCmd.Flags().String("title.border-foreground", "", "Border Foreground Color")
	spinCmd.Flags().Bool("title.faint", false, "Faint text")
	spinCmd.Flags().String("title.foreground", "", "Foreground Color")
	spinCmd.Flags().String("title.height", "", "Text height")
	spinCmd.Flags().Bool("title.italic", false, "Italicize text")
	spinCmd.Flags().String("title.margin", "", "Text margin")
	spinCmd.Flags().String("title.padding", "", "Text padding")
	spinCmd.Flags().Bool("title.strikethrough", false, "Strikethrough text")
	spinCmd.Flags().Bool("title.underline", false, "Underline text")
	spinCmd.Flags().String("title.width", "", "Text width")
	rootCmd.AddCommand(spinCmd)

	common.AddFlagCompletion(spinCmd)
	carapace.Gen(spinCmd).FlagCompletion(carapace.ActionMap{
		"align": carapace.ActionValues("left", "right"), // differs from gum.ActionAlignments
		"spinner": carapace.ActionValuesDescribed(
			"line", "/",
			"dot", "⢿",
			"minidot", "⠋",
			"jump", "⡈",
			"pulse", "░",
			"points", "●",
			"globe", "🌍",
			"moon", "🌗",
			"monkey", "🙊",
			"meter", "▰",
			"hamburger", "☲",
		),
	})

	carapace.Gen(spinCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin(),
	)

	// TODO is this still valid?
	carapace.Gen(spinCmd).DashAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if spinCmd.ArgsLenAtDash() > 0 {
				return carapace.ActionValues()
			}
			return bridge.ActionCarapaceBin()
		}),
	)
}
