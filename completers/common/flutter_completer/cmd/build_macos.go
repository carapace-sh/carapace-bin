package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var build_macosCmd = &cobra.Command{
	Use:   "macos",
	Short: "Build a macOS desktop application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(build_macosCmd).Standalone()

	build_macosCmd.Flags().Bool("analyze-size", false, "Produce additional profile information for artifact output size.")
	build_macosCmd.Flags().StringArray("dart-define", nil, "Additional key-value pairs that will be available as constants.")
	build_macosCmd.Flags().Bool("debug", false, "Build a debug version of your app.")
	build_macosCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	build_macosCmd.Flags().Bool("no-analyze-size", false, "Do not produce additional profile information for artifact output size.")
	build_macosCmd.Flags().Bool("no-null-assertions", false, "Do not perform additional null assertions on the boundaries of migrated and un-migrated code.")
	build_macosCmd.Flags().Bool("no-obfuscate", false, "In a release build, this flag does not remove identifiers and replaces them with randomized values for the purposes of source code obfuscation.")
	build_macosCmd.Flags().Bool("no-pub", false, "Do not run \"flutter pub get\" before executing this command.")
	build_macosCmd.Flags().Bool("no-track-widget-creation", false, "Do not track widget creation locations.")
	build_macosCmd.Flags().Bool("no-tree-shake-icons", false, "Do not tree shake icon fonts so that only glyphs used by the application remain.")
	build_macosCmd.Flags().Bool("null-assertions", false, "Perform additional null assertions on the boundaries of migrated and un-migrated code.")
	build_macosCmd.Flags().Bool("obfuscate", false, "In a release build, this flag removes identifiers and replaces them with randomized values for the purposes of source code obfuscation.")
	build_macosCmd.Flags().Bool("profile", false, "Build a version of your app specialized for performance profiling.")
	build_macosCmd.Flags().Bool("pub", false, "Run \"flutter pub get\" before executing this command.")
	build_macosCmd.Flags().Bool("release", false, "Build a release version of your app (default mode).")
	build_macosCmd.Flags().String("split-debug-info", "", "In a release build, this flag reduces application size by storing Dart program symbols in a separate file on the host rather than in the application.")
	build_macosCmd.Flags().StringP("target", "t", "", "The main entry-point file of the application, as run on the device.")
	build_macosCmd.Flags().Bool("track-widget-creation", false, "Track widget creation locations.")
	build_macosCmd.Flags().Bool("tree-shake-icons", false, "Tree shake icon fonts so that only glyphs used by the application remain.")
	buildCmd.AddCommand(build_macosCmd)

	carapace.Gen(build_macosCmd).FlagCompletion(carapace.ActionMap{
		"split-debug-info": carapace.ActionDirectories(),
		"target":           carapace.ActionFiles(".dart"),
	})
}
