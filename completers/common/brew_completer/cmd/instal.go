package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var instalCmd = &cobra.Command{
	Use:   "instal",
	Short: "Install a <formula> or <cask>",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(instalCmd).Standalone()

	instalCmd.Flags().Bool("HEAD", false, "If <formula> defines it, install the HEAD version, aka. main, trunk, unstable, master.")
	instalCmd.Flags().Bool("adopt", false, "Adopt existing artifacts in the destination that are identical to those being installed. Cannot be combined with `--force`.")
	instalCmd.Flags().Bool("appdir", false, "Target location for Applications (default: `/Applications`).")
	instalCmd.Flags().Bool("audio-unit-plugindir", false, "Target location for Audio Unit Plugins (default: `~/Library/Audio/Plug-Ins/Components`).")
	instalCmd.Flags().Bool("binaries", false, "Disable/enable linking of helper executables (default: enabled).")
	instalCmd.Flags().Bool("bottle-arch", false, "Optimise bottles for the specified architecture rather than the oldest architecture supported by the version of macOS the bottles are built on.")
	instalCmd.Flags().Bool("build-bottle", false, "Prepare the formula for eventual bottling during installation, skipping any post-install steps.")
	instalCmd.Flags().Bool("build-from-source", false, "Compile <formula> from source even if a bottle is provided. Dependencies will still be installed from bottles if they are available.")
	instalCmd.Flags().Bool("cask", false, "Treat all named arguments as casks.")
	instalCmd.Flags().Bool("cc", false, "Attempt to compile using the specified <compiler>, which should be the name of the compiler's executable, e.g. `gcc-7` for GCC 7. In order to use LLVM's clang, specify `llvm_clang`. To use the Apple-provided clang, specify `clang`. This option will only accept compilers that are provided by Homebrew or bundled with macOS. Please do not file issues if you encounter errors while using this option.")
	instalCmd.Flags().Bool("colorpickerdir", false, "Target location for Color Pickers (default: `~/Library/ColorPickers`).")
	instalCmd.Flags().Bool("debug", false, "If brewing fails, open an interactive debugging session with access to IRB or a shell inside the temporary build directory.")
	instalCmd.Flags().Bool("debug-symbols", false, "Generate debug symbols on build. Source will be retained in a cache directory.")
	instalCmd.Flags().Bool("dictionarydir", false, "Target location for Dictionaries (default: `~/Library/Dictionaries`).")
	instalCmd.Flags().Bool("display-times", false, "Print install times for each package at the end of the run.")
	instalCmd.Flags().Bool("dry-run", false, "Show what would be installed, but do not actually install anything.")
	instalCmd.Flags().Bool("fetch-HEAD", false, "Fetch the upstream repository to detect if the HEAD installation of the formula is outdated. Otherwise, the repository's HEAD will only be checked for updates when a new stable or development version has been released.")
	instalCmd.Flags().Bool("fontdir", false, "Target location for Fonts (default: `~/Library/Fonts`).")
	instalCmd.Flags().Bool("force", false, "Install formulae without checking for previously installed keg-only or non-migrated versions. When installing casks, overwrite existing files (binaries and symlinks are excluded, unless originally from the same cask).")
	instalCmd.Flags().Bool("force-bottle", false, "Install from a bottle if it exists for the current or newest version of macOS, even if it would not normally be used for installation.")
	instalCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae.")
	instalCmd.Flags().Bool("git", false, "Create a Git repository, useful for creating patches to the software.")
	instalCmd.Flags().Bool("help", false, "Show this message.")
	instalCmd.Flags().Bool("ignore-dependencies", false, "An unsupported Homebrew development option to skip installing any dependencies of any kind. If the dependencies are not already present, the formula will have issues. If you're not developing Homebrew, consider adjusting your PATH rather than using this option.")
	instalCmd.Flags().Bool("include-test", false, "Install testing dependencies required to run `brew test` <formula>.")
	instalCmd.Flags().Bool("input-methoddir", false, "Target location for Input Methods (default: `~/Library/Input Methods`).")
	instalCmd.Flags().Bool("interactive", false, "Download and patch <formula>, then open a shell. This allows the user to run `./configure --help` and otherwise determine how to turn the software package into a Homebrew package.")
	instalCmd.Flags().Bool("internet-plugindir", false, "Target location for Internet Plugins (default: `~/Library/Internet Plug-Ins`).")
	instalCmd.Flags().Bool("keep-tmp", false, "Retain the temporary files created during installation.")
	instalCmd.Flags().Bool("keyboard-layoutdir", false, "Target location for Keyboard Layouts (default: `/Library/Keyboard Layouts`).")
	instalCmd.Flags().Bool("language", false, "Comma-separated list of language codes to prefer for cask installation. The first matching language is used, otherwise it reverts to the cask's default language. The default value is the language of your system.")
	instalCmd.Flags().Bool("mdimporterdir", false, "Target location for Spotlight Plugins (default: `~/Library/Spotlight`).")
	instalCmd.Flags().Bool("no-binaries", false, "Disable/enable linking of helper executables (default: enabled).")
	instalCmd.Flags().Bool("no-quarantine", false, "Disable/enable quarantining of downloads (default: enabled).")
	instalCmd.Flags().Bool("only-dependencies", false, "Install the dependencies with specified options but do not install the formula itself.")
	instalCmd.Flags().Bool("overwrite", false, "Delete files that already exist in the prefix while linking.")
	instalCmd.Flags().Bool("prefpanedir", false, "Target location for Preference Panes (default: `~/Library/PreferencePanes`).")
	instalCmd.Flags().Bool("qlplugindir", false, "Target location for Quick Look Plugins (default: `~/Library/QuickLook`).")
	instalCmd.Flags().Bool("quarantine", false, "Disable/enable quarantining of downloads (default: enabled).")
	instalCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	instalCmd.Flags().Bool("require-sha", false, "Require all casks to have a checksum.")
	instalCmd.Flags().Bool("screen-saverdir", false, "Target location for Screen Savers (default: `~/Library/Screen Savers`).")
	instalCmd.Flags().Bool("servicedir", false, "Target location for Services (default: `~/Library/Services`).")
	instalCmd.Flags().Bool("skip-cask-deps", false, "Skip installing cask dependencies.")
	instalCmd.Flags().Bool("skip-post-install", false, "Install but skip any post-install steps.")
	instalCmd.Flags().Bool("verbose", false, "Print the verification and post-install steps.")
	instalCmd.Flags().Bool("vst-plugindir", false, "Target location for VST Plugins (default: `~/Library/Audio/Plug-Ins/VST`).")
	instalCmd.Flags().Bool("vst3-plugindir", false, "Target location for VST3 Plugins (default: `~/Library/Audio/Plug-Ins/VST3`).")
	instalCmd.Flags().Bool("zap", false, "For use with `brew reinstall --cask`. Remove all files associated with a cask. *May remove files which are shared between applications.*")
	rootCmd.AddCommand(instalCmd)
}
