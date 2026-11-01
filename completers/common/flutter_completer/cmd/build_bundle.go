package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var build_bundleCmd = &cobra.Command{
	Use:   "bundle",
	Short: "Build the Flutter assets directory from your app",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(build_bundleCmd).Standalone()

	build_bundleCmd.Flags().String("asset-dir", "", "The output directory for the kernel_blob.bin file, the native snapshot, the assets, etc. Can be used to redirect the output when driving the Flutter toolchain from another build system.")
	build_bundleCmd.Flags().String("build-number", "", "An identifier used as an internal version number. Each build must have a unique identifier to differentiate it from previous builds.")
	build_bundleCmd.Flags().StringArrayP("dart-define", "D", nil, "Additional key-value pairs that will be available as constants from the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors. Multiple defines can be passed by repeating \"--dart-define\".")
	build_bundleCmd.Flags().StringSlice("dart-define-from-file", nil, "The path of a .json or .env file containing key-value pairs that will be available as environment variables. These can be accessed using the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors. Multiple define files can be passed by repeating \"--dart-define-from-file\".")
	build_bundleCmd.Flags().StringArray("dart-defines", nil, "Additional key-value pairs that will be available as constants from the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors. Multiple defines can be passed by repeating \"--dart-define\".")
	build_bundleCmd.Flags().Bool("debug", false, "Build a debug version of your app.")
	build_bundleCmd.Flags().String("depfile", "", "A file path where a depfile will be written. This contains all build inputs and outputs in a Make-style syntax.")
	build_bundleCmd.Flags().StringSlice("extra-front-end-options", nil, "A comma-separated list of additional command line arguments that will be passed directly to the Dart front end. For example, \"--extra-front-end-options=--enable-experiment=nonfunction-type-aliases\".")
	build_bundleCmd.Flags().StringSlice("extra-gen-snapshot-options", nil, "A comma-separated list of additional command line arguments that will be passed directly to the Dart native compiler. (Requires the \"--release\", \"--profile\", or \"--jit-release\" flag.)")
	build_bundleCmd.Flags().Bool("fatal-warnings", false, "Causes the command to fail if warnings are sent to the console during its execution.")
	build_bundleCmd.Flags().StringSlice("filesystem-root", nil, "Specify the path that is used as the root of a virtual file system during compilation. The input file name should be specified as a URL using the scheme given in \"--filesystem-scheme\". Requires the \"--output-dill\" option to be explicitly specified.")
	build_bundleCmd.Flags().String("filesystem-scheme", "", "Specify the scheme that is used for virtual file system used in compilation. See also the \"--filesystem-root\" option.")
	build_bundleCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	build_bundleCmd.Flags().Bool("jit-release", false, "Build a JIT release version of your app.")
	build_bundleCmd.Flags().Bool("no-fatal-warnings", false, "Causes the command to fail if warnings are sent to the console during its execution.")
	build_bundleCmd.Flags().Bool("no-pub", false, "Whether to run \"flutter pub get\" before executing this command.")
	build_bundleCmd.Flags().Bool("no-track-widget-creation", false, "Track widget creation locations. This enables features such as the widget inspector. This parameter is only functional in debug mode (i.e. when compiling JIT, not AOT).")
	build_bundleCmd.Flags().Bool("no-tree-shake-icons", false, "(deprecated) Icon font tree shaking is not supported by this command.")
	build_bundleCmd.Flags().String("output-dill", "", "Specify the path to frontend server output kernel file.")
	build_bundleCmd.Flags().Bool("profile", false, "Build a version of your app specialized for performance profiling.")
	build_bundleCmd.Flags().Bool("pub", false, "Whether to run \"flutter pub get\" before executing this command.")
	build_bundleCmd.Flags().Bool("release", false, "Build a release version of your app.")
	build_bundleCmd.Flags().StringP("target", "t", "", "The main entrypoint file of the application, as run on the device. If the \"--target\" option is omitted, but a file name is provided on the command line, then that file is used instead.")
	build_bundleCmd.Flags().String("target-platform", "", "The architecture for which to build the application.")
	build_bundleCmd.Flags().Bool("track-widget-creation", false, "Track widget creation locations. This enables features such as the widget inspector. This parameter is only functional in debug mode (i.e. when compiling JIT, not AOT).")
	build_bundleCmd.Flags().Bool("tree-shake-icons", false, "(deprecated) Icon font tree shaking is not supported by this command.")
	buildCmd.AddCommand(build_bundleCmd)

	carapace.Gen(build_bundleCmd).FlagCompletion(carapace.ActionMap{
		"asset-dir":             carapace.ActionDirectories(),
		"dart-define-from-file": carapace.ActionFiles(".json", ".env"),
		"depfile":               carapace.ActionFiles(),
		"filesystem-root":       carapace.ActionDirectories(),
		"output-dill":           carapace.ActionFiles(),
		"target":                carapace.ActionFiles(".dart"),
		"target-platform":       carapace.ActionValues("android-arm", "android-arm64", "android-x64", "darwin", "ios", "linux-arm64", "linux-riscv64", "linux-x64", "windows-arm64", "windows-x64"),
	})
}
