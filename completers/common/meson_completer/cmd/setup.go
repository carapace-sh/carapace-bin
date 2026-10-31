package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Configure the project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setupCmd).Standalone()

	setupCmd.Flags().StringS("D", "D", "", "Set the value of an option, can be used several times to set multiple options.")
	setupCmd.Flags().String("auto-features", "", "Override value of all 'auto' features (default: auto).")
	setupCmd.Flags().String("backend", "", "Backend to use (default: ninja).")
	setupCmd.Flags().String("bindir", "", "Executable directory (default: bin).")
	setupCmd.Flags().String("build.cmake-prefix-path", "", "List of additional prefixes for cmake to search (just for build machine, default: []).")
	setupCmd.Flags().String("build.pkg-config-path", "", "List of additional paths for pkg-config to search (just for build machine, default: []).")
	setupCmd.Flags().String("buildtype", "", "Build type to use (default: debug).")
	setupCmd.Flags().Bool("clearcache", false, "Clear cached state (e.g. found dependencies).")
	setupCmd.Flags().String("cmake-prefix-path", "", "List of additional prefixes for cmake to search (just for host machine, default: []).")
	setupCmd.Flags().String("cross-file", "", "File describing cross compilation environment.")
	setupCmd.Flags().String("datadir", "", "Data file directory (default: share).")
	setupCmd.Flags().String("debug", "", "Enable debug symbols and other information.")
	setupCmd.Flags().String("default-both-libraries", "", "Default library type for both_libraries (default: shared).")
	setupCmd.Flags().String("default-library", "", "Default library type (default: shared).")
	setupCmd.Flags().String("errorlogs", "", "Whether to print the logs from failing tests.")
	setupCmd.Flags().Bool("fatal-meson-warnings", false, "Make all Meson warnings fatal")
	setupCmd.Flags().String("force-fallback-for", "", "Force fallback for those subprojects (default: []).")
	setupCmd.Flags().String("genvslite", "", "Setup multiple buildtype-suffixed ninja-backend build directories, and a [builddir]_vs containing a Visual Studio meta-backend with multiple configurations that calls into them (default: vs2022).")
	setupCmd.Flags().String("includedir", "", "Header file directory (default: include).")
	setupCmd.Flags().String("infodir", "", "Info page directory (default: share/info).")
	setupCmd.Flags().String("install-umask", "", "Default umask to apply on permissions of installed files (default: 0o22).")
	setupCmd.Flags().String("layout", "", "Build directory layout (default: mirror).")
	setupCmd.Flags().String("libdir", "", "Library directory (default: lib64).")
	setupCmd.Flags().String("libexecdir", "", "Library executable directory (default: libexec).")
	setupCmd.Flags().String("licensedir", "", "Licenses directory (default: ).")
	setupCmd.Flags().String("localedir", "", "Locale data directory (default: share/locale).")
	setupCmd.Flags().String("localstatedir", "", "Localstate data directory (default: var).")
	setupCmd.Flags().String("mandir", "", "Manual page directory (default: share/man).")
	setupCmd.Flags().String("namingscheme", "", "How target file names are formed (default: classic).")
	setupCmd.Flags().String("native-file", "", "File containing overrides for native compilation environment.")
	setupCmd.Flags().String("optimization", "", "Optimization level (default: 0).")
	setupCmd.Flags().String("os2-emxomf", "", "Use OMF format on OS/2.")
	setupCmd.Flags().String("pkg-config-path", "", "List of additional paths for pkg-config to search (just for host machine, default: []).")
	setupCmd.Flags().String("pkgconfig.relocatable", "", "Generate pkgconfig files as relocatable.")
	setupCmd.Flags().String("prefer-static", "", "Whether to try static linking before shared linking.")
	setupCmd.Flags().String("prefix", "", "Installation prefix (default: /usr/local).")
	setupCmd.Flags().Bool("profile-self", false, "==SUPPRESS==")
	setupCmd.Flags().String("python.allow-limited-api", "", "Whether to allow use of the Python Limited API.")
	setupCmd.Flags().String("python.build-config", "", "Config file containing the build details for the target Python installation (default: ).")
	setupCmd.Flags().String("python.bytecompile", "", "Whether to compile bytecode (default: 0).")
	setupCmd.Flags().String("python.install-env", "", "Which python environment to install to (default: prefix).")
	setupCmd.Flags().String("python.platlibdir", "", "Directory for site-specific, platform-specific files (default: ).")
	setupCmd.Flags().String("python.purelibdir", "", "Directory for site-specific, non-platform-specific files (default: ).")
	setupCmd.Flags().Bool("reconfigure", false, "Set options and reconfigure the project.")
	setupCmd.Flags().String("sbindir", "", "System executable directory (default: sbin).")
	setupCmd.Flags().String("sharedstatedir", "", "Architecture-independent data directory (default: com).")
	setupCmd.Flags().String("stdsplit", "", "Split stdout and stderr in test logs.")
	setupCmd.Flags().String("strip", "", "Strip targets on install.")
	setupCmd.Flags().String("sysconfdir", "", "Sysconf data directory (default: etc).")
	setupCmd.Flags().String("unity", "", "Unity build (default: off).")
	setupCmd.Flags().String("unity-size", "", "Unity block size (default: 4).")
	setupCmd.Flags().StringP("version", "v", "", "show program's version number and exit")
	setupCmd.Flags().String("vsenv", "", "Activate Visual Studio environment.")
	setupCmd.Flags().String("warnlevel", "", "Compiler warning level to use (default: 1).")
	setupCmd.Flags().String("werror", "", "Treat warnings as errors.")
	setupCmd.Flags().Bool("wipe", false, "Wipe build directory and reconfigure using previous command line options.")
	setupCmd.Flags().String("wrap-mode", "", "Wrap mode (default: default).")
	applyCommonFlagCompletion(setupCmd)
	rootCmd.AddCommand(setupCmd)
}
