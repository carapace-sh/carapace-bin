package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "actool",
	Short: "Asset catalog compiler",
	Long:  "https://keith.github.io/xcode-manpages/actool.1.html",
	Run:   func(*cobra.Command, []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("accent-color", "", "select a named color to use as the target's accent/tint color")
	rootCmd.Flags().String("alternate-app-icon", "", "specify an additional app icon set name to include")
	rootCmd.Flags().String("app-icon", "", "select a primary app icon set")
	rootCmd.Flags().String("asset-pack-output-specifications", "", "write info about on-demand resources to file")
	rootCmd.Flags().String("compile", "", "compile the asset catalog and write to path")
	rootCmd.Flags().Bool("compress-pngs", false, "process PNGs with pngcrush")
	rootCmd.Flags().Bool("enable-on-demand-resources", false, "process on-demand resources")
	rootCmd.Flags().Bool("errors", false, "include document error messages in output")
	rootCmd.Flags().String("filter-for-device-model", "", "filter files by device model")
	rootCmd.Flags().String("filter-for-device-os-version", "", "filter files by device OS version")
	rootCmd.Flags().Bool("include-all-app-icons", false, "include all app icon assets from all named catalogs")
	rootCmd.Flags().String("include-language", "", "filter compilation to include only the specified language")
	rootCmd.Flags().Bool("include-partial-info-plist-localizations", true, "include localization info in partial Info.plist")
	rootCmd.Flags().Bool("include-sticker-content", false, "include sticker pack content")
	rootCmd.Flags().String("launch-image", "", "select a launch image to compile")
	rootCmd.Flags().String("minimum-deployment-target", "", "minimum deployment target version")
	rootCmd.Flags().Bool("notices", false, "include document notice messages in output")
	rootCmd.Flags().String("output-format", "", "output format: binary1, xml1, or human-readable-text")
	rootCmd.Flags().String("output-partial-info-plist", "", "emit plist with Info.plist keys")
	rootCmd.Flags().String("platform", "", "target platform to compile for")
	rootCmd.Flags().Bool("print-contents", false, "include a listing of the catalog's content in output")
	rootCmd.Flags().String("product-type", "", "set the product type for sticker-specific behaviors")
	rootCmd.Flags().Bool("skip-app-store-deployment", false, "skip App Store-specific behaviors")
	rootCmd.Flags().String("standalone-icon-behavior", "", "control loose PNG/ICNS file creation: default, all, none")
	rootCmd.Flags().String("sticker-pack-identifier-prefix", "", "set the default prefix for sticker path identifiers")
	rootCmd.Flags().String("sticker-pack-strings-file", "", "specify a strings file for sticker name translations")
	rootCmd.Flags().String("stickers-icon-role", "", "select icon sizes for Messages-style app icons")
	rootCmd.Flags().String("target-device", "", "target device to compile for")
	rootCmd.Flags().Bool("version", false, "print the version of actool")
	rootCmd.Flags().Bool("warnings", false, "include document warning messages in output")
	rootCmd.Flags().String("widget-background-color", "", "select a named color for the widget background")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"asset-pack-output-specifications": carapace.ActionFiles(),
		"compile":                          carapace.ActionDirectories(),
		"output-format":                    carapace.ActionValues("binary1", "xml1", "human-readable-text"),
		"output-partial-info-plist":        carapace.ActionFiles(),
		"platform":                         carapace.ActionValues("ios", "macos", "tvos", "watchos"),
		"standalone-icon-behavior":         carapace.ActionValues("default", "all", "none"),
		"sticker-pack-strings-file":        carapace.ActionFiles(),
		"stickers-icon-role":               carapace.ActionValues("app-host", "extension"),
		"target-device":                    carapace.ActionValues("iphone", "ipad", "mac", "tv", "watch"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
