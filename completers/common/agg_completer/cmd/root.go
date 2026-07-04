package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "agg",
	Short: "asciinema gif generator",
	Long:  "https://github.com/asciinema/agg",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("bold-is-bright", false, "Render bold text with bright colors (ANSI 0..7 → 8..15)")
	rootCmd.Flags().String("cols", "", "Override terminal width (number of columns)")
	rootCmd.Flags().String("emoji-font-family", "", "Specify emoji font families")
	rootCmd.Flags().String("font-dir", "", "Use additional font directory; may be specified multiple times")
	rootCmd.Flags().String("font-family", "", "Specify the complete font family list, bypassing automatic fallbacks; must start with a monospace text font")
	rootCmd.Flags().String("font-size", "", "Specify font size (in pixels)")
	rootCmd.Flags().String("fps-cap", "", "Set FPS cap")
	rootCmd.Flags().BoolP("help", "h", false, "Print help (see a summary with '-h')")
	rootCmd.Flags().String("idle-time-limit", "", "Limit idle time to max number of seconds")
	rootCmd.Flags().String("last-frame-duration", "", "Set last frame duration")
	rootCmd.Flags().String("line-height", "", "Specify line height")
	rootCmd.Flags().Bool("no-loop", false, "Disable animation loop")
	rootCmd.Flags().BoolP("quiet", "q", false, "Quiet mode - suppress diagnostic messages and progress bars")
	rootCmd.Flags().String("renderer", "", "Select frame rendering backend")
	rootCmd.Flags().String("rows", "", "Override terminal height")
	rootCmd.Flags().String("select", "", "Select frames to render")
	rootCmd.Flags().String("speed", "", "Adjust playback speed")
	rootCmd.Flags().String("text-font-family", "", "Specify regular text font families")
	rootCmd.Flags().String("theme", "", "Select color theme")
	rootCmd.Flags().CountP("verbose", "v", "Enable verbose logging")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"emoji-font-family": os.ActionFontFamilies().UniqueList(","),
		"font-dir":          carapace.ActionDirectories(),
		"font-family":       os.ActionFontFamilies().UniqueList(","),
		"renderer":          carapace.ActionValues("swash", "resvg"),
		"text-font-family":  os.ActionFontFamilies().UniqueList(","),
		"theme":             carapace.ActionValues("asciinema", "dracula", "github-dark", "github-light", "kanagawa", "kanagawa-dragon", "kanagawa-light", "monokai", "nord", "solarized-dark", "solarized-light", "gruvbox-dark", "custom"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
