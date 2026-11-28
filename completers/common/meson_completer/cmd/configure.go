package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Change project options",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configureCmd).Standalone()

	configureCmd.Flags().StringS("D", "D", "", "Set the value of an option, can be used several times to set multiple options.")
	configureCmd.Flags().StringS("U", "U", "", "Remove a subproject option.")
	configureCmd.Flags().String("auto-features", "", "Override value of all 'auto' features (default: auto).")
	configureCmd.Flags().String("backend", "", "Backend to use (default: ninja).")
	configureCmd.Flags().String("bindir", "", "Executable directory (default: bin).")
	configureCmd.Flags().String("build.cmake-prefix-path", "", "List of additional prefixes for cmake to search (just for build machine, default: []).")
	configureCmd.Flags().String("build.pkg-config-path", "", "List of additional paths for pkg-config to search (just for build machine, default: []).")
	configureCmd.Flags().String("buildtype", "", "Build type to use (default: debug).")
	configureCmd.Flags().Bool("clearcache", false, "Clear cached state (e.g. found dependencies)")
	configureCmd.Flags().String("cmake-prefix-path", "", "List of additional prefixes for cmake to search (just for host machine, default: []).")
	configureCmd.Flags().String("datadir", "", "Data file directory (default: share).")
	configureCmd.Flags().String("debug", "", "Enable debug symbols and other information.")
	configureCmd.Flags().String("default-both-libraries", "", "Default library type for both_libraries (default: shared).")
	configureCmd.Flags().String("default-library", "", "Default library type (default: shared).")
	configureCmd.Flags().String("errorlogs", "", "Whether to print the logs from failing tests.")
	configureCmd.Flags().String("force-fallback-for", "", "Force fallback for those subprojects (default: []).")
	configureCmd.Flags().String("genvslite", "", "Setup multiple buildtype-suffixed ninja-backend build directories, and a [builddir]_vs containing a Visual Studio meta-backend with multiple configurations that calls into them (default: vs2022).")
	configureCmd.Flags().String("includedir", "", "Header file directory (default: include).")
	configureCmd.Flags().String("infodir", "", "Info page directory (default: share/info).")
	configureCmd.Flags().String("install-umask", "", "Default umask to apply on permissions of installed files (default: 0o22).")
	configureCmd.Flags().String("layout", "", "Build directory layout (default: mirror).")
	configureCmd.Flags().String("libdir", "", "Library directory (default: lib64).")
	configureCmd.Flags().String("libexecdir", "", "Library executable directory (default: libexec).")
	configureCmd.Flags().String("licensedir", "", "Licenses directory (default: ).")
	configureCmd.Flags().String("localedir", "", "Locale data directory (default: share/locale).")
	configureCmd.Flags().String("localstatedir", "", "Localstate data directory (default: var).")
	configureCmd.Flags().String("mandir", "", "Manual page directory (default: share/man).")
	configureCmd.Flags().String("namingscheme", "", "How target file names are formed (default: classic).")
	configureCmd.Flags().Bool("no-pager", false, "Do not redirect output to a pager")
	configureCmd.Flags().String("optimization", "", "Optimization level (default: 0).")
	configureCmd.Flags().String("os2-emxomf", "", "Use OMF format on OS/2.")
	configureCmd.Flags().String("pkg-config-path", "", "List of additional paths for pkg-config to search (just for host machine, default: []).")
	configureCmd.Flags().String("pkgconfig.relocatable", "", "Generate pkgconfig files as relocatable.")
	configureCmd.Flags().String("prefer-static", "", "Whether to try static linking before shared linking.")
	configureCmd.Flags().String("prefix", "", "Installation prefix (default: /usr/local).")
	configureCmd.Flags().String("python.allow-limited-api", "", "Whether to allow use of the Python Limited API.")
	configureCmd.Flags().String("python.build-config", "", "Config file containing the build details for the target Python installation (default: ).")
	configureCmd.Flags().String("python.bytecompile", "", "Whether to compile bytecode (default: 0).")
	configureCmd.Flags().String("python.install-env", "", "Which python environment to install to (default: prefix).")
	configureCmd.Flags().String("python.platlibdir", "", "Directory for site-specific, platform-specific files (default: ).")
	configureCmd.Flags().String("python.purelibdir", "", "Directory for site-specific, non-platform-specific files (default: ).")
	configureCmd.Flags().String("sbindir", "", "System executable directory (default: sbin).")
	configureCmd.Flags().String("sharedstatedir", "", "Architecture-independent data directory (default: com).")
	configureCmd.Flags().String("stdsplit", "", "Split stdout and stderr in test logs.")
	configureCmd.Flags().String("strip", "", "Strip targets on install.")
	configureCmd.Flags().String("sysconfdir", "", "Sysconf data directory (default: etc).")
	configureCmd.Flags().String("unity", "", "Unity build (default: off).")
	configureCmd.Flags().String("unity-size", "", "Unity block size (default: 4).")
	configureCmd.Flags().String("vsenv", "", "Activate Visual Studio environment.")
	configureCmd.Flags().String("warnlevel", "", "Compiler warning level to use (default: 1).")
	configureCmd.Flags().String("werror", "", "Treat warnings as errors.")
	configureCmd.Flags().String("wrap-mode", "", "Wrap mode (default: default).")
	applyCommonFlagCompletion(configureCmd)
	rootCmd.AddCommand(configureCmd)
}
