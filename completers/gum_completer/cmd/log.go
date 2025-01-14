package cmd

import (
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gum"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Log messages to output",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logCmd).Standalone()

	logCmd.Flags().StringP("file", "o", "", "Log to file")
	logCmd.Flags().BoolP("format", "f", false, "Format message using printf")
	logCmd.Flags().String("formatter", "", "The log formatter to use")
	logCmd.Flags().String("key.align", "", "Text Alignment")
	logCmd.Flags().String("key.background", "", "Background Color")
	logCmd.Flags().Bool("key.bold", false, "Bold text")
	logCmd.Flags().String("key.border", "", "Border Style")
	logCmd.Flags().String("key.border-background", "", "Border Background Color")
	logCmd.Flags().String("key.border-foreground", "", "Border Foreground Color")
	logCmd.Flags().Bool("key.faint", false, "Faint text")
	logCmd.Flags().String("key.foreground", "", "Foreground Color")
	logCmd.Flags().String("key.height", "", "Text height")
	logCmd.Flags().Bool("key.italic", false, "Italicize text")
	logCmd.Flags().String("key.margin", "", "Text margin")
	logCmd.Flags().String("key.padding", "", "Text padding")
	logCmd.Flags().Bool("key.strikethrough", false, "Strikethrough text")
	logCmd.Flags().Bool("key.underline", false, "Underline text")
	logCmd.Flags().String("key.width", "", "Text width")
	logCmd.Flags().StringP("level", "l", "", "The log level to use")
	logCmd.Flags().String("level.align", "", "Text Alignment")
	logCmd.Flags().String("level.background", "", "Background Color")
	logCmd.Flags().Bool("level.bold", false, "Bold text")
	logCmd.Flags().String("level.border", "", "Border Style")
	logCmd.Flags().String("level.border-background", "", "Border Background Color")
	logCmd.Flags().String("level.border-foreground", "", "Border Foreground Color")
	logCmd.Flags().Bool("level.faint", false, "Faint text")
	logCmd.Flags().String("level.foreground", "", "Foreground Color")
	logCmd.Flags().String("level.height", "", "Text height")
	logCmd.Flags().Bool("level.italic", false, "Italicize text")
	logCmd.Flags().String("level.margin", "", "Text margin")
	logCmd.Flags().String("level.padding", "", "Text padding")
	logCmd.Flags().Bool("level.strikethrough", false, "Strikethrough text")
	logCmd.Flags().Bool("level.underline", false, "Underline text")
	logCmd.Flags().String("level.width", "", "Text width")
	logCmd.Flags().String("message.align", "", "Text Alignment")
	logCmd.Flags().String("message.background", "", "Background Color")
	logCmd.Flags().Bool("message.bold", false, "Bold text")
	logCmd.Flags().String("message.border", "", "Border Style")
	logCmd.Flags().String("message.border-background", "", "Border Background Color")
	logCmd.Flags().String("message.border-foreground", "", "Border Foreground Color")
	logCmd.Flags().Bool("message.faint", false, "Faint text")
	logCmd.Flags().String("message.foreground", "", "Foreground Color")
	logCmd.Flags().String("message.height", "", "Text height")
	logCmd.Flags().Bool("message.italic", false, "Italicize text")
	logCmd.Flags().String("message.margin", "", "Text margin")
	logCmd.Flags().String("message.padding", "", "Text padding")
	logCmd.Flags().Bool("message.strikethrough", false, "Strikethrough text")
	logCmd.Flags().Bool("message.underline", false, "Underline text")
	logCmd.Flags().String("message.width", "", "Text width")
	logCmd.Flags().String("min-level", "", "Minimal level to show")
	logCmd.Flags().String("prefix", "", "Prefix to print before the message")
	logCmd.Flags().String("prefix.align", "", "Text Alignment")
	logCmd.Flags().String("prefix.background", "", "Background Color")
	logCmd.Flags().Bool("prefix.bold", false, "Bold text")
	logCmd.Flags().String("prefix.border", "", "Border Style")
	logCmd.Flags().String("prefix.border-background", "", "Border Background Color")
	logCmd.Flags().String("prefix.border-foreground", "", "Border Foreground Color")
	logCmd.Flags().Bool("prefix.faint", false, "Faint text")
	logCmd.Flags().String("prefix.foreground", "", "Foreground Color")
	logCmd.Flags().String("prefix.height", "", "Text height")
	logCmd.Flags().Bool("prefix.italic", false, "Italicize text")
	logCmd.Flags().String("prefix.margin", "", "Text margin")
	logCmd.Flags().String("prefix.padding", "", "Text padding")
	logCmd.Flags().Bool("prefix.strikethrough", false, "Strikethrough text")
	logCmd.Flags().Bool("prefix.underline", false, "Underline text")
	logCmd.Flags().String("prefix.width", "", "Text width")
	logCmd.Flags().String("separator.align", "", "Text Alignment")
	logCmd.Flags().String("separator.background", "", "Background Color")
	logCmd.Flags().Bool("separator.bold", false, "Bold text")
	logCmd.Flags().String("separator.border", "", "Border Style")
	logCmd.Flags().String("separator.border-background", "", "Border Background Color")
	logCmd.Flags().String("separator.border-foreground", "", "Border Foreground Color")
	logCmd.Flags().Bool("separator.faint", false, "Faint text")
	logCmd.Flags().String("separator.foreground", "", "Foreground Color")
	logCmd.Flags().String("separator.height", "", "Text height")
	logCmd.Flags().Bool("separator.italic", false, "Italicize text")
	logCmd.Flags().String("separator.margin", "", "Text margin")
	logCmd.Flags().String("separator.padding", "", "Text padding")
	logCmd.Flags().Bool("separator.strikethrough", false, "Strikethrough text")
	logCmd.Flags().Bool("separator.underline", false, "Underline text")
	logCmd.Flags().String("separator.width", "", "Text width")
	logCmd.Flags().BoolP("structured", "s", false, "Use structured logging")
	logCmd.Flags().StringP("time", "t", "", "The time format to use (kitchen, layout, ansic, rfc822, etc...)")
	logCmd.Flags().String("time.align", "", "Text Alignment")
	logCmd.Flags().String("time.background", "", "Background Color")
	logCmd.Flags().Bool("time.bold", false, "Bold text")
	logCmd.Flags().String("time.border", "", "Border Style")
	logCmd.Flags().String("time.border-background", "", "Border Background Color")
	logCmd.Flags().String("time.border-foreground", "", "Border Foreground Color")
	logCmd.Flags().Bool("time.faint", false, "Faint text")
	logCmd.Flags().String("time.foreground", "", "Foreground Color")
	logCmd.Flags().String("time.height", "", "Text height")
	logCmd.Flags().Bool("time.italic", false, "Italicize text")
	logCmd.Flags().String("time.margin", "", "Text margin")
	logCmd.Flags().String("time.padding", "", "Text padding")
	logCmd.Flags().Bool("time.strikethrough", false, "Strikethrough text")
	logCmd.Flags().Bool("time.underline", false, "Underline text")
	logCmd.Flags().String("time.width", "", "Text width")
	logCmd.Flags().String("value.align", "", "Text Alignment")
	logCmd.Flags().String("value.background", "", "Background Color")
	logCmd.Flags().Bool("value.bold", false, "Bold text")
	logCmd.Flags().String("value.border", "", "Border Style")
	logCmd.Flags().String("value.border-background", "", "Border Background Color")
	logCmd.Flags().String("value.border-foreground", "", "Border Foreground Color")
	logCmd.Flags().Bool("value.faint", false, "Faint text")
	logCmd.Flags().String("value.foreground", "", "Foreground Color")
	logCmd.Flags().String("value.height", "", "Text height")
	logCmd.Flags().Bool("value.italic", false, "Italicize text")
	logCmd.Flags().String("value.margin", "", "Text margin")
	logCmd.Flags().String("value.padding", "", "Text padding")
	logCmd.Flags().Bool("value.strikethrough", false, "Strikethrough text")
	logCmd.Flags().Bool("value.underline", false, "Underline text")
	logCmd.Flags().String("value.width", "", "Text width")
	rootCmd.AddCommand(logCmd)

	carapace.Gen(logCmd).FlagCompletion(carapace.ActionMap{
		"file":                        carapace.ActionFiles(),
		"formatter":                   carapace.ActionValues("json", "logfmt", "text"),
		"key.align":                   gum.ActionAlignments(),
		"key.background":              gum.ActionColors(),
		"key.border":                  gum.ActionBorders(),
		"key.border-background":       gum.ActionColors(),
		"key.border-foreground":       gum.ActionColors(),
		"key.foreground":              gum.ActionColors(),
		"level":                       carapace.ActionValues("none", "debug", "info", "warn", "error", "fatal").StyleF(style.ForLogLevel),
		"level.align":                 gum.ActionAlignments(),
		"level.background":            gum.ActionColors(),
		"level.border":                gum.ActionBorders(),
		"level.border-background":     gum.ActionColors(),
		"level.border-foreground":     gum.ActionColors(),
		"level.foreground":            gum.ActionColors(),
		"message.align":               gum.ActionAlignments(),
		"message.background":          gum.ActionColors(),
		"message.border":              gum.ActionBorders(),
		"message.border-background":   gum.ActionColors(),
		"message.border-foreground":   gum.ActionColors(),
		"message.foreground":          gum.ActionColors(),
		"prefix.align":                gum.ActionAlignments(),
		"prefix.background":           gum.ActionColors(),
		"prefix.border":               gum.ActionBorders(),
		"prefix.border-background":    gum.ActionColors(),
		"prefix.border-foreground":    gum.ActionColors(),
		"prefix.foreground":           gum.ActionColors(),
		"separator.align":             gum.ActionAlignments(),
		"separator.background":        gum.ActionColors(),
		"separator.border":            gum.ActionBorders(),
		"separator.border-background": gum.ActionColors(),
		"separator.border-foreground": gum.ActionColors(),
		"separator.foreground":        gum.ActionColors(),
		"time": carapace.ActionValuesDescribed(
			"layout", time.Layout,
			"ansic", time.ANSIC,
			"unixdate", time.UnixDate,
			"rubydate", time.RubyDate,
			"rfc822", time.RFC822,
			"rfc822z", time.RFC822Z,
			"rfc850", time.RFC850,
			"rfc1123", time.RFC1123,
			"rfc1123z", time.RFC1123Z,
			"rfc3339", time.RFC3339,
			"rfc3339nano", time.RFC3339Nano,
			"kitchen", time.Kitchen,
			"stamp", time.Stamp,
			"stampmilli", time.StampMilli,
			"stampmicro", time.StampMicro,
			"stampnano", time.StampNano,
			"datetime", "2006-01-02 15:04:05",
			"dateonly", "2006-01-02",
			"timeonly", "15:04:05"),
		"time.align":              gum.ActionAlignments(),
		"time.background":         gum.ActionColors(),
		"time.border":             gum.ActionBorders(),
		"time.border-background":  gum.ActionColors(),
		"time.border-foreground":  gum.ActionColors(),
		"time.foreground":         gum.ActionColors(),
		"value.align":             gum.ActionAlignments(),
		"value.background":        gum.ActionColors(),
		"value.border":            gum.ActionBorders(),
		"value.border-background": gum.ActionColors(),
		"value.border-foreground": gum.ActionColors(),
		"value.foreground":        gum.ActionColors(),
	})
}
