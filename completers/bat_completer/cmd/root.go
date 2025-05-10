package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/bat"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace/pkg/style"
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

	rootCmd.Flags().Bool("acknowledgements", false, "Show acknowledgements")
	rootCmd.Flags().BoolP("chop-long-lines", "S", false, "Truncate all lines longer than screen width")
	rootCmd.Flags().String("color", "", "When to use colors")
	rootCmd.Flags().String("decorations", "", "When to show the decorations")
	rootCmd.Flags().Bool("diagnostic", false, "Show diagnostic information for bug reports")
	rootCmd.Flags().BoolP("diff", "d", false, "Only show lines that have been added/removed/modified")
	rootCmd.Flags().String("diff-context", "", "Include N lines of context around added/removed/modified")
	rootCmd.Flags().StringSlice("file-name", nil, "Specify the name to display for a file")
	rootCmd.Flags().BoolP("force-colorization", "f", false, "Alias for '--decorations=always --color=always'")
	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.Flags().StringSliceP("highlight-line", "H", nil, "Highlight lines N through M")
	rootCmd.Flags().StringSlice("ignored-suffix", nil, "Ignore extension")
	rootCmd.Flags().String("italic-text", "", "Use italics in output")
	rootCmd.Flags().StringP("language", "l", "", "Set the language for syntax highlighting")
	rootCmd.Flags().StringSliceP("line-range", "r", nil, "Only print the lines from N to M")
	rootCmd.Flags().BoolP("list-languages", "L", false, "Display all supported languages")
	rootCmd.Flags().Bool("list-themes", false, "Display all supported highlighting themes")
	rootCmd.Flags().StringSliceP("map-syntax", "m", nil, "Use the specified syntax for files matching the glob pattern")
	rootCmd.Flags().BoolP("no-pager", "", false, "Disable paging")
	rootCmd.Flags().BoolP("no-paging", "", false, "Disable paging")
	rootCmd.Flags().String("nonprintable-notation", "", "Set notation for non-printable characters")
	rootCmd.Flags().BoolP("number", "n", false, "Show line numbers")
	rootCmd.Flags().String("pager", "", "Determine which pager to use")
	rootCmd.Flags().String("paging", "", "Specify when to use the pager, or use `-P` to disable")
	rootCmd.Flags().CountP("plain", "p", "Show plain style")
	rootCmd.Flags().Bool("cache-dir", false, "Show bat's cache directory")
	rootCmd.Flags().Bool("config-dir", false, "Show bat's configuration directory")
	rootCmd.Flags().Bool("config-file", false, "Show path to the configuration file")
	rootCmd.Flags().Bool("generate-config-file", false, "Generates a default configuration file")
	rootCmd.Flags().Bool("lessopen", false, "Enable the $LESSOPEN preprocessor")
	rootCmd.Flags().Bool("no-config", false, "Do not use the configuration file")
	rootCmd.Flags().Bool("no-custom-assets", false, "Do not load custom assets")
	rootCmd.Flags().Bool("no-lessopen", false, "Disable the $LESSOPEN preprocessor if enabled")
	rootCmd.Flags().BoolP("show-all", "A", false, "Show non-printable characters")
	rootCmd.Flags().String("style", "", "Comma-separated list of style elements to display")
	rootCmd.Flags().String("tabs", "", "Set the tab width to T spaces")
	rootCmd.Flags().String("terminal-width", "", "Explicitly set the width of the terminal")
	rootCmd.Flags().String("theme", "", "Set the color theme for syntax highlighting")
	rootCmd.Flags().BoolP("unbuffered", "u", false, "This option exists for POSIX-compliance reasons")
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
		"pager":  bridge.ActionCarapaceBin().Split(),
		"paging": carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"style":  bat.ActionStyles().UniqueList(","),
		"theme":  bat.ActionThemes(),
		"wrap":   carapace.ActionValues("auto", "never", "character"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
