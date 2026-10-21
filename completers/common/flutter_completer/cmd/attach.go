package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/adb"
	"github.com/spf13/cobra"
)

var attachCmd = &cobra.Command{
	Use:   "attach",
	Short: "Attach to a running app",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(attachCmd).Standalone()

	attachCmd.Flags().Bool("adb-log-filtering", false, "Filter adb logs so that logs not emitted by the current flutter application are not displayed.")
	attachCmd.Flags().String("app-id", "", "The package name (Android) or bundle identifier (iOS) for the app. This can be specified to avoid being prompted if multiple Dart VM Service ports are advertised. If you have multiple devices or emulators running, you should include the device hostname as well, e.g. \"com.example.myApp@my-iphone\". This parameter is case-insensitive.")
	attachCmd.Flags().Bool("assume-initialize-from-dill-up-to-date", false, "If set, assumes that the file passed in initialize-from-dill is up to date and skip the check and potential invalidation of files.")
	attachCmd.Flags().StringArrayP("dart-define", "D", nil, "Additional key-value pairs that will be available as constants from the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors. Multiple defines can be passed by repeating \"--dart-define\".")
	attachCmd.Flags().StringSlice("dart-define-from-file", nil, "The path of a .json or .env file containing key-value pairs that will be available as environment variables. These can be accessed using the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors. Multiple define files can be passed by repeating \"--dart-define-from-file\".")
	attachCmd.Flags().StringArray("dart-defines", nil, "Additional key-value pairs that will be available as constants from the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors. Multiple defines can be passed by repeating \"--dart-define\".")
	attachCmd.Flags().Bool("dds", false, "Enable the Dart Developer Service (DDS). It may be necessary to disable this when attaching to an application with an existing DDS instance (e.g., attaching to an application currently connected to by \"flutter run\"), or when running certain tests. Disabling this feature may degrade IDE functionality if a DDS instance is not already connected to the target application.")
	attachCmd.Flags().Int("dds-port", 0, "When this value is provided, the Dart Development Service (DDS) will be bound to the provided port. Specifying port 0 (the default) will find a random free port.")
	attachCmd.Flags().Bool("debug", false, "Build a debug version of your app.")
	attachCmd.Flags().String("debug-port", "", "(deprecated) Device port where the Dart VM Service is listening. Requires \"--disable-service-auth-codes\" to also be provided to the Flutter application at launch, otherwise this command will fail to connect to the application. In general, \"--debug-url\" should be used instead.")
	attachCmd.Flags().String("debug-uri", "", "The URL at which the Dart VM Service is listening.")
	attachCmd.Flags().String("debug-url", "", "The URL at which the Dart VM Service is listening.")
	attachCmd.Flags().String("device-connection", "", "Discover devices based on connection type.")
	attachCmd.Flags().String("device-timeout", "", "Time in seconds to wait for devices to attach. Longer timeouts may be necessary for networked devices.")
	attachCmd.Flags().String("device-user", "", "Identifier number for a user or work profile on Android only. Run \"adb shell pm list users\" for available identifiers.")
	attachCmd.Flags().String("device-vmservice-port", "", "Look for vmservice connections only from the specified port. Specifying port 0 (the default) will accept the first vmservice discovered.")
	attachCmd.Flags().Bool("devtools", false, "Enable (or disable, with \"--no-devtools\") the launching of the Flutter DevTools debugger and profiler. If \"--no-devtools\" is specified, \"--devtools-server-address\" is ignored.")
	attachCmd.Flags().String("devtools-server-address", "", "When this value is provided, the Flutter tool will not spin up a new DevTools server instance, and will instead use the one provided at the given address. Ignored if \"--no-devtools\" is specified.")
	attachCmd.Flags().Bool("disable-dds", false, "(deprecated; use \"--no-dds\" instead) Disable the Dart Developer Service (DDS).")
	attachCmd.Flags().StringSlice("enable-experiment", nil, "The name of an experimental Dart feature to enable. Multiple experiments can be enabled by repeating \"--enable-experiment\".")
	attachCmd.Flags().StringSlice("filesystem-root", nil, "Specify the path that is used as the root of a virtual file system during compilation. The input file name should be specified as a URL using the scheme given in \"--filesystem-scheme\". Requires the \"--output-dill\" option to be explicitly specified.")
	attachCmd.Flags().String("filesystem-scheme", "", "Specify the scheme that is used for virtual file system used in compilation. See also the \"--filesystem-root\" option.")
	attachCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	attachCmd.Flags().String("host-vmservice-port", "", "When a device-side vmservice port is forwarded to a host-side port, use this value as the host port. Specifying port 0 (the default) will find a random free host port.")
	attachCmd.Flags().String("initialize-from-dill", "", "Initializes the resident compiler with a specific kernel file instead of the default cached location.")
	attachCmd.Flags().Bool("ipv6", false, "Binds to IPv6 localhost instead of IPv4 when the flutter tool forwards the host port to a device port.")
	attachCmd.Flags().Bool("machine", false, "Outputs in a machine readable structured JSON format.")
	attachCmd.Flags().StringP("module", "m", "", "The name of the module (required if attaching to a fuchsia device).")
	attachCmd.Flags().String("native-assets-yaml-file", "", "Initializes the resident compiler with a custom native assets yaml file instead of the default cached location.")
	attachCmd.Flags().Bool("no-adb-log-filtering", false, "Filter adb logs so that logs not emitted by the current flutter application are not displayed.")
	attachCmd.Flags().Bool("no-assume-initialize-from-dill-up-to-date", false, "If set, assumes that the file passed in initialize-from-dill is up to date and skip the check and potential invalidation of files.")
	attachCmd.Flags().Bool("no-dds", false, "Enable the Dart Developer Service (DDS). It may be necessary to disable this when attaching to an application with an existing DDS instance (e.g., attaching to an application currently connected to by \"flutter run\"), or when running certain tests. Disabling this feature may degrade IDE functionality if a DDS instance is not already connected to the target application.")
	attachCmd.Flags().Bool("no-devtools", false, "Enable (or disable, with \"--no-devtools\") the launching of the Flutter DevTools debugger and profiler. If \"--no-devtools\" is specified, \"--devtools-server-address\" is ignored.")
	attachCmd.Flags().Bool("no-disable-dds", false, "(deprecated; use \"--no-dds\" instead) Disable the Dart Developer Service (DDS).")
	attachCmd.Flags().Bool("no-report-ready", false, "Print \"ready\" to the console after handling a keyboard command. This is primarily useful for tests and other automation, but consider using \"--machine\" instead.")
	attachCmd.Flags().Bool("no-track-widget-creation", false, "Track widget creation locations. This enables features such as the widget inspector. This parameter is only functional in debug mode (i.e. when compiling JIT, not AOT).")
	attachCmd.Flags().String("output-dill", "", "Specify the path to frontend server output kernel file.")
	attachCmd.Flags().String("pid-file", "", "Specify a file to write the process ID to. You can send SIGUSR1 to trigger a hot reload and SIGUSR2 to trigger a hot restart. The file is created when the signal handlers are hooked and deleted when they are removed.")
	attachCmd.Flags().Bool("profile", false, "Build a version of your app specialized for performance profiling.")
	attachCmd.Flags().String("project-root", "", "Specify the project root directory.")
	attachCmd.Flags().Bool("report-ready", false, "Print \"ready\" to the console after handling a keyboard command. This is primarily useful for tests and other automation, but consider using \"--machine\" instead.")
	attachCmd.Flags().StringP("target", "t", "", "The main entrypoint file of the application, as run on the device. If the \"--target\" option is omitted, but a file name is provided on the command line, then that file is used instead.")
	attachCmd.Flags().String("target-model", "", "Target model that determines what core libraries are available.")
	attachCmd.Flags().Bool("track-widget-creation", false, "Track widget creation locations. This enables features such as the widget inspector. This parameter is only functional in debug mode (i.e. when compiling JIT, not AOT).")
	attachCmd.Flags().String("vm-service-port", "", "(deprecated; use host-vmservice-port instead) Listen to the given port for a Dart VM Service connection. Specifying port 0 (the default) will find a random free port. if the Dart Development Service (DDS) is enabled, this will not be the port of the VmService instance advertised on the command line.")
	rootCmd.AddCommand(attachCmd)

	carapace.Gen(attachCmd).FlagCompletion(carapace.ActionMap{
		"app-id": adb.ActionPackages(), // TODO ios
		"device-connection": carapace.ActionValuesDescribed(
			"attached", "Only searches for devices connected by USB or built-in (such as simulators/emulators, MacOS/Windows, Chrome)",
			"wireless", "Only searches for devices connected wirelessly. Discovering wireless devices may take longer.",
			"both", "Searches for both attached and wireless devices.",
		),
		"device-user":  adb.ActionUsers(),
		"pid-file":     carapace.ActionFiles(),
		"target":       carapace.ActionFiles(".dart"),
		"target-model": carapace.ActionValues("flutter", "flutter_runner"),
	})
}
