package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/bridge"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gum"
	"github.com/spf13/cobra"
)

var spinCmd = &cobra.Command{
	Use:                "spin",
	Short:              "Display spinner while running a command",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(spinCmd).Standalone()
	rootCmd.AddCommand(spinCmd)

	carapace.Gen(spinCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			firstPass := _spinCmd()
			firstPass.FParseErrWhitelist = cobra.FParseErrWhitelist{
				UnknownFlags: true,
			}
			args := []string{}
			args = append(args, c.Args...)
			args = append(args, c.CallbackValue)
			if err := firstPass.ParseFlags(args); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			secondPass := _spinCmd()
			if args := firstPass.Flags().Args(); len(args) > 0 && !(len(args) == 1 && args[0] == c.CallbackValue) {
				subcmd := subCmd(args[0])
				secondPass.AddCommand(subcmd)
			}
			return carapace.ActionExecute(secondPass)
		}),
	)
}

func _spinCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "spin",
		Run: func(cmd *cobra.Command, args []string) {},
	}

	carapace.Gen(cmd).Standalone()

	cmd.Flags().StringP("align", "a", "", "Alignment of spinner with regard to the title")
	cmd.Flags().Bool("show-output", false, "Show output of command")
	cmd.Flags().StringP("spinner", "s", "", "Spinner type")
	cmd.Flags().String("spinner.align", "", "Text Alignment")
	cmd.Flags().String("spinner.background", "", "Background Color")
	cmd.Flags().Bool("spinner.bold", false, "Bold text")
	cmd.Flags().String("spinner.border", "", "Border Style")
	cmd.Flags().String("spinner.border-background", "", "Border Background Color")
	cmd.Flags().String("spinner.border-foreground", "", "Border Foreground Color")
	cmd.Flags().Bool("spinner.faint", false, "Faint text")
	cmd.Flags().String("spinner.foreground", "", "Foreground Color")
	cmd.Flags().String("spinner.height", "", "Text height")
	cmd.Flags().Bool("spinner.italic", false, "Italicize text")
	cmd.Flags().String("spinner.margin", "", "Text margin")
	cmd.Flags().String("spinner.padding", "", "Text padding")
	cmd.Flags().Bool("spinner.strikethrough", false, "Strikethrough text")
	cmd.Flags().Bool("spinner.underline", false, "Underline text")
	cmd.Flags().String("spinner.width", "", "Text width")
	cmd.Flags().String("title", "", "Text to display to user while spinning")
	cmd.Flags().String("title.align", "", "Text Alignment")
	cmd.Flags().String("title.background", "", "Background Color")
	cmd.Flags().Bool("title.bold", false, "Bold text")
	cmd.Flags().String("title.border", "", "Border Style")
	cmd.Flags().String("title.border-background", "", "Border Background Color")
	cmd.Flags().String("title.border-foreground", "", "Border Foreground Color")
	cmd.Flags().Bool("title.faint", false, "Faint text")
	cmd.Flags().String("title.foreground", "", "Foreground Color")
	cmd.Flags().String("title.height", "", "Text height")
	cmd.Flags().Bool("title.italic", false, "Italicize text")
	cmd.Flags().String("title.margin", "", "Text margin")
	cmd.Flags().String("title.padding", "", "Text padding")
	cmd.Flags().Bool("title.strikethrough", false, "Strikethrough text")
	cmd.Flags().Bool("title.underline", false, "Underline text")
	cmd.Flags().String("title.width", "", "Text width")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"align": carapace.ActionValues("left", "right"),
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
		"spinner.background":        gum.ActionColors(),
		"spinner.border":            gum.ActionBorders(),
		"spinner.border-background": gum.ActionColors(),
		"spinner.border-foreground": gum.ActionColors(),
		"spinner.foreground":        gum.ActionColors(),
		"title.align":               gum.ActionAlignments(),
		"title.background":          gum.ActionColors(),
		"title.border":              gum.ActionBorders(),
		"title.border-background":   gum.ActionColors(),
		"title.border-foreground":   gum.ActionColors(),
		"title.foreground":          gum.ActionColors(),
	})

	carapace.Gen(cmd).PositionalCompletion(
		carapace.Batch(
			os.ActionPathExecutables(),
			carapace.ActionFiles(),
		).ToA(),
	)

	return cmd
}

func subCmd(name string) *cobra.Command {
	cmd := &cobra.Command{
		Use:                name,
		Run:                func(cmd *cobra.Command, args []string) {},
		DisableFlagParsing: true,
	}

	carapace.Gen(cmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin(name),
	)
	return cmd
}
