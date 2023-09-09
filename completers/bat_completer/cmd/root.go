package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/bat"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bat",
	Short: "a cat clone with syntax highlighting and Git integration",
	Long:  "https://github.com/sharkdp/bat",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("chop-long-lines", "S", false, "Truncate all lines longer than screen width")
	rootCmd.Flags().String("color", "", "When to use colors")
	rootCmd.Flags().String("decorations", "", "When to show the decorations")
	rootCmd.Flags().BoolP("diff", "d", false, "Only show lines that have been added/removed/modified")
	rootCmd.Flags().String("file-name", "", "Specify the name to display for a file")
	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.Flags().StringP("highlight-line", "H", "", "Highlight lines N through M")
	rootCmd.Flags().String("italic-text", "", "Use italics in output")
	rootCmd.Flags().StringP("language", "l", "", "Set the language for syntax highlighting")
	rootCmd.Flags().StringP("line-range", "r", "", "Only print the lines from N to M")
	rootCmd.Flags().BoolP("list-languages", "L", false, "Display all supported languages")
	rootCmd.Flags().Bool("list-themes", false, "Display all supported highlighting themes")
	rootCmd.Flags().StringSliceP("map-syntax", "m", []string{}, "Use the specified syntax for files matching the glob pattern")
	rootCmd.Flags().String("nonprintable-notation", "", "Set notation for non-printable characters")
	rootCmd.Flags().BoolP("number", "n", false, "Show line numbers")
	rootCmd.Flags().String("paging", "", "Specify when to use the pager, or use `-P` to disable")
	rootCmd.Flags().CountP("plain", "p", "Show plain style")
	rootCmd.Flags().BoolP("show-all", "A", false, "Show non-printable characters")
	rootCmd.Flags().String("style", "", "Comma-separated list of style elements to display")
	rootCmd.Flags().String("tabs", "", "Set the tab width to T spaces")
	rootCmd.Flags().String("theme", "", "Set the color theme for syntax highlighting")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")
	rootCmd.Flags().String("wrap", "", "Specify the text-wrapping mode")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"color":       carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"decorations": carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"italic-text": carapace.ActionValues("never", "always").StyleF(style.ForKeyword),
		"language":    bat.ActionLanguages(),
		"map-syntax": carapace.ActionMultiPartsN(":", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues()
			default:
				return bat.ActionLanguages()
			}
		}),
		"nonprintable-notation": carapace.ActionValuesDescribed(
			"caret", "Use character sequences like ˆG, ˆJ, ˆ@, .. to identify non-printable characters",
			"unicode", "Use special Unicode code points to identify non-printable characters",
		),
		"paging": carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"style":  bat.ActionStyles().UniqueList(","),
		"theme":  bat.ActionThemes(),
		"wrap":   carapace.ActionValues("auto", "never", "character"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
