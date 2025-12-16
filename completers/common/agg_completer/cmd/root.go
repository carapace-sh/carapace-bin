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

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("cols", "", "Override terminal width")
	rootCmd.Flags().String("font-dir", "", "Use additional font directory")
	rootCmd.Flags().String("font-family", "", "Specify font family")
	rootCmd.Flags().String("font-size", "", "Specify font size (in pixels)")
	rootCmd.Flags().String("fps-cap", "", "Set FPS cap")
	rootCmd.Flags().BoolP("help", "h", false, "Print help information")
	rootCmd.Flags().String("idle-time-limit", "", "Limit idle time to max number of seconds")
	rootCmd.Flags().String("line-height", "", "Specify line height")
	rootCmd.Flags().Bool("no-loop", false, "Disable animation loop")
	rootCmd.Flags().String("renderer", "", "Select frame rendering backend")
	rootCmd.Flags().String("rows", "", "Override terminal height")
	rootCmd.Flags().String("speed", "", "Adjust playback speed")
	rootCmd.Flags().String("theme", "", "Select color theme")
	rootCmd.Flags().BoolP("verbose", "v", false, "Enable verbose logging")
	rootCmd.Flags().BoolP("version", "V", false, "Print version information")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"font-dir":    carapace.ActionDirectories(),
		"font-family": os.ActionFontFamilies().UniqueList(","),
		"renderer":    carapace.ActionValues("fontdue", "resvg"),
		"theme":       carapace.ActionValues("asciinema", "dracula", "monokai", "solarized-dark", "solarized-light", "custom"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
