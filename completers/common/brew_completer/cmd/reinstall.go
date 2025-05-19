package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var reinstallCmd = &cobra.Command{
	Use:     "reinstall",
	Short:   "Uninstall and then reinstall a <formula> or <cask> using the same options it was originally installed with, plus any appended options specific to a <formula>",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reinstallCmd).Standalone()

	reinstallCmd.Flags().Bool("adopt", false, "Adopt existing artifacts in the destination that are identical to those being installed. Cannot be combined with `--force`.")
	reinstallCmd.Flags().Bool("appdir", false, "Target location for Applications (default: `/Applications`).")
	reinstallCmd.Flags().Bool("audio-unit-plugindir", false, "Target location for Audio Unit Plugins (default: `~/Library/Audio/Plug-Ins/Components`).")
	reinstallCmd.Flags().Bool("binaries", false, "Disable/enable linking of helper executables (default: enabled).")
	reinstallCmd.Flags().Bool("build-from-source", false, "Compile <formula> from source even if a bottle is available.")
	reinstallCmd.Flags().Bool("cask", false, "Treat all named arguments as casks.")
	reinstallCmd.Flags().Bool("colorpickerdir", false, "Target location for Color Pickers (default: `~/Library/ColorPickers`).")
	reinstallCmd.Flags().Bool("debug", false, "If brewing fails, open an interactive debugging session with access to IRB or a shell inside the temporary build directory.")
	reinstallCmd.Flags().Bool("debug-symbols", false, "Generate debug symbols on build. Source will be retained in a cache directory.")
	reinstallCmd.Flags().Bool("dictionarydir", false, "Target location for Dictionaries (default: `~/Library/Dictionaries`).")
	reinstallCmd.Flags().Bool("display-times", false, "Print install times for each formula at the end of the run.")
	reinstallCmd.Flags().Bool("fontdir", false, "Target location for Fonts (default: `~/Library/Fonts`).")
	reinstallCmd.Flags().Bool("force", false, "Install without checking for previously installed keg-only or non-migrated versions.")
	reinstallCmd.Flags().Bool("force-bottle", false, "Install from a bottle if it exists for the current or newest version of macOS, even if it would not normally be used for installation.")
	reinstallCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae.")
	reinstallCmd.Flags().Bool("git", false, "Create a Git repository, useful for creating patches to the software.")
	reinstallCmd.Flags().Bool("help", false, "Show this message.")
	reinstallCmd.Flags().Bool("input-methoddir", false, "Target location for Input Methods (default: `~/Library/Input Methods`).")
	reinstallCmd.Flags().Bool("interactive", false, "Download and patch <formula>, then open a shell. This allows the user to run `./configure --help` and otherwise determine how to turn the software package into a Homebrew package.")
	reinstallCmd.Flags().Bool("internet-plugindir", false, "Target location for Internet Plugins (default: `~/Library/Internet Plug-Ins`).")
	reinstallCmd.Flags().Bool("keep-tmp", false, "Retain the temporary files created during installation.")
	reinstallCmd.Flags().Bool("keyboard-layoutdir", false, "Target location for Keyboard Layouts (default: `/Library/Keyboard Layouts`).")
	reinstallCmd.Flags().Bool("language", false, "Comma-separated list of language codes to prefer for cask installation. The first matching language is used, otherwise it reverts to the cask's default language. The default value is the language of your system.")
	reinstallCmd.Flags().Bool("mdimporterdir", false, "Target location for Spotlight Plugins (default: `~/Library/Spotlight`).")
	reinstallCmd.Flags().Bool("no-binaries", false, "Disable/enable linking of helper executables (default: enabled).")
	reinstallCmd.Flags().Bool("no-quarantine", false, "Disable/enable quarantining of downloads (default: enabled).")
	reinstallCmd.Flags().Bool("prefpanedir", false, "Target location for Preference Panes (default: `~/Library/PreferencePanes`).")
	reinstallCmd.Flags().Bool("qlplugindir", false, "Target location for Quick Look Plugins (default: `~/Library/QuickLook`).")
	reinstallCmd.Flags().Bool("quarantine", false, "Disable/enable quarantining of downloads (default: enabled).")
	reinstallCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	reinstallCmd.Flags().Bool("require-sha", false, "Require all casks to have a checksum.")
	reinstallCmd.Flags().Bool("screen-saverdir", false, "Target location for Screen Savers (default: `~/Library/Screen Savers`).")
	reinstallCmd.Flags().Bool("servicedir", false, "Target location for Services (default: `~/Library/Services`).")
	reinstallCmd.Flags().Bool("skip-cask-deps", false, "Skip installing cask dependencies.")
	reinstallCmd.Flags().Bool("verbose", false, "Print the verification and post-install steps.")
	reinstallCmd.Flags().Bool("vst-plugindir", false, "Target location for VST Plugins (default: `~/Library/Audio/Plug-Ins/VST`).")
	reinstallCmd.Flags().Bool("vst3-plugindir", false, "Target location for VST3 Plugins (default: `~/Library/Audio/Plug-Ins/VST3`).")
	reinstallCmd.Flags().Bool("zap", false, "For use with `brew reinstall --cask`. Remove all files associated with a cask. *May remove files which are shared between applications.*")
	rootCmd.AddCommand(reinstallCmd)
}
