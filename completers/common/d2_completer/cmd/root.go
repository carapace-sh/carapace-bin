package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/d2"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "d2",
	Short: "compiles and renders d2 diagrams into svgs",
	Long:  "https://d2lang.com/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("animate-interval", "", "multiple boards are packaged as 1 SVG")
	rootCmd.Flags().String("ascii-mode", "", "ASCII rendering mode for text outputs")
	rootCmd.Flags().String("browser", "", "browser executable that watch opens")
	rootCmd.Flags().BoolP("bundle", "b", false, "bundle all assets and layers into the output file")
	rootCmd.Flags().BoolP("center", "c", false, "center the SVG in the containing viewbox")
	rootCmd.Flags().Bool("check", false, "check that the specified files are formatted correctly")
	rootCmd.Flags().String("dark-theme", "", "the theme to use when the viewer's browser is in dark mode")
	rootCmd.Flags().BoolP("debug", "d", false, "print debug logs")
	rootCmd.Flags().String("font-bold", "", "path to .ttf file to use for the bold font")
	rootCmd.Flags().String("font-italic", "", "path to .ttf file to use for the italic font")
	rootCmd.Flags().String("font-mono", "", "path to .ttf file to use for the monospace font")
	rootCmd.Flags().String("font-mono-bold", "", "path to .ttf file to use for the monospace bold font")
	rootCmd.Flags().String("font-mono-italic", "", "path to .ttf file to use for the monospace italic font")
	rootCmd.Flags().String("font-mono-semibold", "", "path to .ttf file to use for the monospace semibold font")
	rootCmd.Flags().String("font-regular", "", "path to .ttf file to use for the regular font")
	rootCmd.Flags().String("font-semibold", "", "path to .ttf file to use for the semibold font")
	rootCmd.Flags().Bool("force-appendix", false, "add an appendix to SVG exports")
	rootCmd.Flags().StringP("host", "h", "", "host listening address when used with watch")
	rootCmd.Flags().Bool("img-cache", false, "images used in icons are cached for subsequent compilations")
	rootCmd.Flags().StringP("layout", "l", "", "the layout engine used")
	rootCmd.Flags().Bool("no-xml-tag", false, "omit XML tag (<?xml ...?>) from output SVG files")
	rootCmd.Flags().Bool("omit-version", false, "omit D2 version from generated image")
	rootCmd.Flags().String("pad", "", "pixels padded around the rendered diagram")
	rootCmd.Flags().StringP("port", "p", "", "port listening address when used with watch")
	rootCmd.Flags().String("salt", "", "Add a salt value to ensure the output uses unique IDs")
	rootCmd.Flags().String("scale", "", "scale the output")
	rootCmd.Flags().BoolP("sketch", "s", false, "render the diagram to look like it was sketched by hand")
	rootCmd.Flags().String("stdout-format", "", "output format when writing to stdout")
	rootCmd.Flags().String("target", "", "target board to render")
	rootCmd.Flags().StringP("theme", "t", "", "the diagram theme ID")
	rootCmd.Flags().String("timeout", "", "the maximum number of seconds that D2 runs")
	rootCmd.Flags().BoolP("version", "v", false, "get the version")
	rootCmd.Flags().BoolP("watch", "w", false, "watch for changes to input and live reload")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"ascii-mode": carapace.ActionValuesDescribed(
			"standard", "basic ASCII chars",
			"extended", "Unicode chars",
		),
		"browser": carapace.Batch(
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
		"dark-theme":         d2.ActionThemes(d2.ThemeOpts{Dark: true}),
		"font-bold":          carapace.ActionFiles(".ttf"),
		"font-italic":        carapace.ActionFiles(".ttf"),
		"font-mono":          carapace.ActionFiles(".ttf"),
		"font-mono-bold":     carapace.ActionFiles(".ttf"),
		"font-mono-italic":   carapace.ActionFiles(".ttf"),
		"font-mono-semibold": carapace.ActionFiles(".ttf"),
		"font-regular":       carapace.ActionFiles(".ttf"),
		"font-semibold":      carapace.ActionFiles(".ttf"),
		"layout":             d2.ActionLayouts(),
		"pad":                carapace.ActionValues(),
		"port":               net.ActionPorts(),
		"stdout-format":      carapace.ActionValues("svg", "png", "ascii", "txt", "pdf", "pptx", "gif"),
		"target":             carapace.ActionValues(), // TODO
		"theme":              d2.ActionThemes(d2.ThemeOpts{}.Default()),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(".d2"),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
