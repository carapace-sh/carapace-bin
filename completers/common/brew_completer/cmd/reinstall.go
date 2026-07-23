package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/brew_completer/cmd/action"
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

	reinstallCmd.Flags().Bool("HEAD", false, "If <formula> defines it, install the HEAD version, aka. main, trunk, unstable, master.")
	reinstallCmd.Flags().Bool("adopt", false, "Adopt existing artifacts in the destination that are identical to those being installed. Cannot be combined with `--force`.")
	reinstallCmd.Flags().String("appdir", "", "Target location for Applications (default: `/Applications`).")
	reinstallCmd.Flags().String("audio-unit-plugindir", "", "Target location for Audio Unit Plugins (default: `~/Library/Audio/Plug-Ins/Components`).")
	reinstallCmd.Flags().Bool("binaries", false, "Disable/enable linking of helper executables (default: enabled).")
	reinstallCmd.Flags().Bool("build-from-source", false, "Compile <formula> from source even if a bottle is available.")
	reinstallCmd.Flags().Bool("cask", false, "Treat all named arguments as casks.")
	reinstallCmd.Flags().String("colorpickerdir", "", "Target location for Color Pickers (default: `~/Library/ColorPickers`).")
	reinstallCmd.Flags().Bool("debug", false, "If brewing fails, open an interactive debugging session with access to IRB or a shell inside the temporary build directory.")
	reinstallCmd.Flags().Bool("debug-symbols", false, "Generate debug symbols on build. Source will be retained in a cache directory.")
	reinstallCmd.Flags().String("dictionarydir", "", "Target location for Dictionaries (default: `~/Library/Dictionaries`).")
	reinstallCmd.Flags().Bool("display-times", false, "Print install times for each formula at the end of the run.")
	reinstallCmd.Flags().BoolS("dry-run", "n", false, "Show what would be reinstalled, but do not actually reinstall anything.")
	reinstallCmd.Flags().String("fontdir", "", "Target location for Fonts (default: `~/Library/Fonts`).")
	reinstallCmd.Flags().Bool("force", false, "Install without checking for previously installed keg-only or non-migrated versions.")
	reinstallCmd.Flags().Bool("force-bottle", false, "Install from a bottle if it exists for the current or newest version of macOS, even if it would not normally be used for installation.")
	reinstallCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae.")
	reinstallCmd.Flags().Bool("git", false, "Create a Git repository, useful for creating patches to the software.")
	reinstallCmd.Flags().Bool("help", false, "Show this message.")
	reinstallCmd.Flags().String("input-methoddir", "", "Target location for Input Methods (default: `~/Library/Input Methods`).")
	reinstallCmd.Flags().Bool("interactive", false, "Download and patch <formula>, then open a shell. This allows the user to run `./configure --help` and otherwise determine how to turn the software package into a Homebrew package.")
	reinstallCmd.Flags().String("internet-plugindir", "", "Target location for Internet Plugins (default: `~/Library/Internet Plug-Ins`).")
	reinstallCmd.Flags().Bool("keep-tmp", false, "Retain the temporary files created during installation.")
	reinstallCmd.Flags().String("keyboard-layoutdir", "", "Target location for Keyboard Layouts (default: `/Library/Keyboard Layouts`).")
	reinstallCmd.Flags().String("language", "", "Comma-separated list of language codes to prefer for cask installation. The first matching language is used, otherwise it reverts to the cask's default language. The default value is the language of your system.")
	reinstallCmd.Flags().String("mdimporterdir", "", "Target location for Spotlight Plugins (default: `~/Library/Spotlight`).")
	reinstallCmd.Flags().BoolS("no-ask", "y", false, "Do not ask for confirmation before downloading and reinstalling.")
	reinstallCmd.Flags().Bool("no-binaries", false, "Disable/enable linking of helper executables (default: enabled).")
	reinstallCmd.Flags().Bool("no-quarantine", false, "Disable/enable quarantining of downloads (default: enabled).")
	reinstallCmd.Flags().String("prefpanedir", "", "Target location for Preference Panes (default: `~/Library/PreferencePanes`).")
	reinstallCmd.Flags().String("qlplugindir", "", "Target location for Quick Look Plugins (default: `~/Library/QuickLook`).")
	reinstallCmd.Flags().Bool("quarantine", false, "Disable/enable quarantining of downloads (default: enabled).")
	reinstallCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	reinstallCmd.Flags().Bool("require-sha", false, "Require all casks to have a checksum.")
	reinstallCmd.Flags().String("screen-saverdir", "", "Target location for Screen Savers (default: `~/Library/Screen Savers`).")
	reinstallCmd.Flags().String("servicedir", "", "Target location for Services (default: `~/Library/Services`).")
	reinstallCmd.Flags().Bool("skip-cask-deps", false, "Skip installing cask dependencies.")
	reinstallCmd.Flags().Bool("verbose", false, "Print the verification and post-install steps.")
	reinstallCmd.Flags().String("vst-plugindir", "", "Target location for VST Plugins (default: `~/Library/Audio/Plug-Ins/VST`).")
	reinstallCmd.Flags().String("vst3-plugindir", "", "Target location for VST3 Plugins (default: `~/Library/Audio/Plug-Ins/VST3`).")
	reinstallCmd.Flags().Bool("zap", false, "For use with `brew reinstall --cask`. Remove all files associated with a cask. *May remove files which are shared between applications.*")
	rootCmd.AddCommand(reinstallCmd)

	carapace.Gen(reinstallCmd).FlagCompletion(carapace.ActionMap{
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

	carapace.Gen(reinstallCmd).PositionalAnyCompletion(
		action.ActionList(reinstallCmd).FilterArgs(),
	)
}
