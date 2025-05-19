package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/brew_completer/cmd/action"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:     "upgrade",
	Short:   "Upgrade outdated casks and outdated, unpinned formulae using the same options they were originally installed with, plus any appended brew formula options",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgradeCmd).Standalone()

	upgradeCmd.Flags().String("appdir", "", "Target location for Applications (default: `/Applications`).")
	upgradeCmd.Flags().String("audio-unit-plugindir", "", "Target location for Audio Unit Plugins (default: `~/Library/Audio/Plug-Ins/Components`).")
	upgradeCmd.Flags().Bool("binaries", false, "Disable/enable linking of helper executables (default: enabled).")
	upgradeCmd.Flags().Bool("build-from-source", false, "Compile <formula> from source even if a bottle is available.")
	upgradeCmd.Flags().Bool("cask", false, "Treat all named arguments as casks. If no named arguments are specified, upgrade only outdated casks.")
	upgradeCmd.Flags().String("colorpickerdir", "", "Target location for Color Pickers (default: `~/Library/ColorPickers`).")
	upgradeCmd.Flags().Bool("debug", false, "If brewing fails, open an interactive debugging session with access to IRB or a shell inside the temporary build directory.")
	upgradeCmd.Flags().Bool("debug-symbols", false, "Generate debug symbols on build. Source will be retained in a cache directory.")
	upgradeCmd.Flags().String("dictionarydir", "", "Target location for Dictionaries (default: `~/Library/Dictionaries`).")
	upgradeCmd.Flags().Bool("display-times", false, "Print install times for each package at the end of the run.")
	upgradeCmd.Flags().Bool("dry-run", false, "Show what would be upgraded, but do not actually upgrade anything.")
	upgradeCmd.Flags().Bool("fetch-HEAD", false, "Fetch the upstream repository to detect if the HEAD installation of the formula is outdated. Otherwise, the repository's HEAD will only be checked for updates when a new stable or development version has been released.")
	upgradeCmd.Flags().String("fontdir", "", "Target location for Fonts (default: `~/Library/Fonts`).")
	upgradeCmd.Flags().Bool("force", false, "Install formulae without checking for previously installed keg-only or non-migrated versions. When installing casks, overwrite existing files (binaries and symlinks are excluded, unless originally from the same cask).")
	upgradeCmd.Flags().Bool("force-bottle", false, "Install from a bottle if it exists for the current or newest version of macOS, even if it would not normally be used for installation.")
	upgradeCmd.Flags().Bool("formula", false, "Treat all named arguments as formulae. If no named arguments are specified, upgrade only outdated formulae.")
	upgradeCmd.Flags().Bool("greedy", false, "Also include casks with `auto_updates true` or `version :latest`.")
	upgradeCmd.Flags().Bool("greedy-auto-updates", false, "Also include casks with `auto_updates true`.")
	upgradeCmd.Flags().Bool("greedy-latest", false, "Also include casks with `version :latest`.")
	upgradeCmd.Flags().Bool("help", false, "Show this message.")
	upgradeCmd.Flags().String("input-methoddir", "", "Target location for Input Methods (default: `~/Library/Input Methods`).")
	upgradeCmd.Flags().Bool("interactive", false, "Download and patch <formula>, then open a shell. This allows the user to run `./configure --help` and otherwise determine how to turn the software package into a Homebrew package.")
	upgradeCmd.Flags().String("internet-plugindir", "", "Target location for Internet Plugins (default: `~/Library/Internet Plug-Ins`).")
	upgradeCmd.Flags().Bool("keep-tmp", false, "Retain the temporary files created during installation.")
	upgradeCmd.Flags().Bool("keyboard-layoutdir", false, "Target location for Keyboard Layouts (default: `/Library/Keyboard Layouts`).")
	upgradeCmd.Flags().Bool("language", false, "Comma-separated list of language codes to prefer for cask installation. The first matching language is used, otherwise it reverts to the cask's default language. The default value is the language of your system.")
	upgradeCmd.Flags().String("mdimporterdir", "", "Target location for Spotlight Plugins (default: `~/Library/Spotlight`).")
	upgradeCmd.Flags().Bool("no-binaries", false, "Disable/enable linking of helper executables (default: enabled).")
	upgradeCmd.Flags().Bool("no-quarantine", false, "Disable/enable quarantining of downloads (default: enabled).")
	upgradeCmd.Flags().String("prefpanedir", "", "Target location for Preference Panes (default: `~/Library/PreferencePanes`).")
	upgradeCmd.Flags().String("qlplugindir", "", "Target location for Quick Look Plugins (default: `~/Library/QuickLook`).")
	upgradeCmd.Flags().Bool("quarantine", false, "Disable/enable quarantining of downloads (default: enabled).")
	upgradeCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	upgradeCmd.Flags().Bool("require-sha", false, "Require all casks to have a checksum.")
	upgradeCmd.Flags().String("screen-saverdir", "", "Target location for Screen Savers (default: `~/Library/Screen Savers`).")
	upgradeCmd.Flags().String("servicedir", "", "Target location for Services (default: `~/Library/Services`).")
	upgradeCmd.Flags().Bool("skip-cask-deps", false, "Skip installing cask dependencies.")
	upgradeCmd.Flags().Bool("verbose", false, "Print the verification and post-install steps.")
	upgradeCmd.Flags().String("vst-plugindir", "", "Target location for VST Plugins (default: `~/Library/Audio/Plug-Ins/VST`).")
	upgradeCmd.Flags().String("vst3-plugindir", "", "Target location for VST3 Plugins (default: `~/Library/Audio/Plug-Ins/VST3`).")
	rootCmd.AddCommand(upgradeCmd)

	carapace.Gen(upgradeCmd).FlagCompletion(carapace.ActionMap{
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

	carapace.Gen(upgradeCmd).PositionalAnyCompletion(
		action.ActionList(upgradeCmd),
	)
}
