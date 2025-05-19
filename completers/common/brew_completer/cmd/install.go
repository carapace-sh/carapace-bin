package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/brew_completer/cmd/action"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install",
	Short:   "Install a <formula> or <cask>",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()

	installCmd.Flags().Bool("HEAD", false, "If <formula> defines it, install the HEAD version, aka. main, trunk, unstable, master.")
	installCmd.Flags().Bool("adopt", false, "Adopt existing artifacts in the destination that are identical to those being installed. Cannot be combined with `--force`.")
	installCmd.Flags().String("appdir", "", "Target location for Applications (default: `/Applications`).")
	installCmd.Flags().String("audio-unit-plugindir", "", "Target location for Audio Unit Plugins (default: `~/Library/Audio/Plug-Ins/Components`).")
	installCmd.Flags().Bool("binaries", false, "Disable/enable linking of helper executables (default: enabled).")
	installCmd.Flags().Bool("bottle-arch", false, "Optimise bottles for the specified architecture rather than the oldest architecture supported by the version of macOS the bottles are built on.")
	installCmd.Flags().Bool("build-bottle", false, "Prepare the formula for eventual bottling during installation, skipping any post-install steps.")
	installCmd.Flags().Bool("build-from-source", false, "Compile <formula> from source even if a bottle is provided. Dependencies will still be installed from bottles if they are available.")
	installCmd.Flags().Bool("cask", false, "Treat all named arguments as casks.")
	installCmd.Flags().Bool("cc", false, "Attempt to compile using the specified <compiler>, which should be the name of the compiler's executable, e.g. `gcc-7` for GCC 7. In order to use LLVM's clang, specify `llvm_clang`. To use the Apple-provided clang, specify `clang`. This option will only accept compilers that are provided by Homebrew or bundled with macOS. Please do not file issues if you encounter errors while using this option.")
	installCmd.Flags().String("colorpickerdir", "", "Target location for Color Pickers (default: `~/Library/ColorPickers`).")
	installCmd.Flags().Bool("debug", false, "If brewing fails, open an interactive debugging session with access to IRB or a shell inside the temporary build directory.")
	installCmd.Flags().Bool("debug-symbols", false, "Generate debug symbols on build. Source will be retained in a cache directory.")
	installCmd.Flags().String("dictionarydir", "", "Target location for Dictionaries (default: `~/Library/Dictionaries`).")
	installCmd.Flags().Bool("display-times", false, "Print install times for each package at the end of the run.")
	installCmd.Flags().Bool("dry-run", false, "Show what would be installed, but do not actually install anything.")
	installCmd.Flags().Bool("fetch-HEAD", false, "Fetch the upstream repository to detect if the HEAD installation of the formula is outdated. Otherwise, the repository's HEAD will only be checked for updates when a new stable or development version has been released.")
	installCmd.Flags().String("fontdir", "", "Target location for Fonts (default: `~/Library/Fonts`).")
	installCmd.Flags().Bool("force", false, "Install formulae without checking for previously installed keg-only or non-migrated versions. When installing casks, overwrite existing files (binaries and symlinks are excluded, unless originally from the same cask).")
	installCmd.Flags().Bool("force-bottle", false, "Install from a bottle if it exists for the current or newest version of macOS, even if it would not normally be used for installation.")
	installCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae.")
	installCmd.Flags().Bool("git", false, "Create a Git repository, useful for creating patches to the software.")
	installCmd.Flags().Bool("help", false, "Show this message.")
	installCmd.Flags().Bool("ignore-dependencies", false, "An unsupported Homebrew development option to skip installing any dependencies of any kind. If the dependencies are not already present, the formula will have issues. If you're not developing Homebrew, consider adjusting your PATH rather than using this option.")
	installCmd.Flags().Bool("include-test", false, "Install testing dependencies required to run `brew test` <formula>.")
	installCmd.Flags().String("input-methoddir", "", "Target location for Input Methods (default: `~/Library/Input Methods`).")
	installCmd.Flags().Bool("interactive", false, "Download and patch <formula>, then open a shell. This allows the user to run `./configure --help` and otherwise determine how to turn the software package into a Homebrew package.")
	installCmd.Flags().String("internet-plugindir", "", "Target location for Internet Plugins (default: `~/Library/Internet Plug-Ins`).")
	installCmd.Flags().Bool("keep-tmp", false, "Retain the temporary files created during installation.")
	installCmd.Flags().String("keyboard-layoutdir", "", "Target location for Keyboard Layouts (default: `/Library/Keyboard Layouts`).")
	installCmd.Flags().Bool("language", false, "Comma-separated list of language codes to prefer for cask installation. The first matching language is used, otherwise it reverts to the cask's default language. The default value is the language of your system.")
	installCmd.Flags().String("mdimporterdir", "", "Target location for Spotlight Plugins (default: `~/Library/Spotlight`).")
	installCmd.Flags().Bool("no-binaries", false, "Disable/enable linking of helper executables (default: enabled).")
	installCmd.Flags().Bool("no-quarantine", false, "Disable/enable quarantining of downloads (default: enabled).")
	installCmd.Flags().Bool("only-dependencies", false, "Install the dependencies with specified options but do not install the formula itself.")
	installCmd.Flags().Bool("overwrite", false, "Delete files that already exist in the prefix while linking.")
	installCmd.Flags().String("prefpanedir", "", "Target location for Preference Panes (default: `~/Library/PreferencePanes`).")
	installCmd.Flags().String("qlplugindir", "", "Target location for Quick Look Plugins (default: `~/Library/QuickLook`).")
	installCmd.Flags().Bool("quarantine", false, "Disable/enable quarantining of downloads (default: enabled).")
	installCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	installCmd.Flags().Bool("require-sha", false, "Require all casks to have a checksum.")
	installCmd.Flags().String("screen-saverdir", "", "Target location for Screen Savers (default: `~/Library/Screen Savers`).")
	installCmd.Flags().String("servicedir", "", "Target location for Services (default: `~/Library/Services`).")
	installCmd.Flags().Bool("skip-cask-deps", false, "Skip installing cask dependencies.")
	installCmd.Flags().Bool("skip-post-install", false, "Install but skip any post-install steps.")
	installCmd.Flags().Bool("verbose", false, "Print the verification and post-install steps.")
	installCmd.Flags().String("vst-plugindir", "", "Target location for VST Plugins (default: `~/Library/Audio/Plug-Ins/VST`).")
	installCmd.Flags().String("vst3-plugindir", "", "Target location for VST3 Plugins (default: `~/Library/Audio/Plug-Ins/VST3`).")
	installCmd.Flags().Bool("zap", false, "For use with `brew reinstall --cask`. Remove all files associated with a cask. *May remove files which are shared between applications.*")
	rootCmd.AddCommand(installCmd)

	carapace.Gen(installCmd).FlagCompletion(carapace.ActionMap{
		"appdir":               carapace.ActionDirectories(),
		"audio-unit-plugindir": carapace.ActionDirectories(),
		"colorpickerdir":       carapace.ActionDirectories(),
		"dictionarydir":        carapace.ActionDirectories(),
		"fontdir":              carapace.ActionDirectories(),
		"input-methoddir":      carapace.ActionDirectories(),
		"internet-plugindir":   carapace.ActionDirectories(),
		"mdimporterdir":        carapace.ActionDirectories(),
		"prefpanedir":          carapace.ActionDirectories(),
		"qlplugindir":          carapace.ActionDirectories(),
		"screen-saverdir":      carapace.ActionDirectories(),
		"servicedir":           carapace.ActionDirectories(),
		"vst-plugindir":        carapace.ActionDirectories(),
		"vst3-plugindir":       carapace.ActionDirectories(),
	})

	carapace.Gen(installCmd).PositionalAnyCompletion(
		action.ActionSearch(installCmd),
	)
}
