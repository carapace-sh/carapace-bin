package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/meson_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "meson",
	Short: "The Meson Build System",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

}

// applyCommonFlagCompletion registers completion for flags shared across meson commands
func applyCommonFlagCompletion(cmd *cobra.Command) {
	actions := carapace.ActionMap{
		"C":                       carapace.ActionDirectories(),
		"backend":                 action.ActionBackends(),
		"bindir":                  carapace.ActionDirectories(),
		"build.cmake-prefix-path": carapace.ActionDirectories(),
		"build.pkg-config-path":   carapace.ActionDirectories(),
		"builddir":                carapace.ActionDirectories(),
		"buildtype":               action.ActionBuildTypes(),
		"cmake-prefix-path":       carapace.ActionDirectories(),
		"cross-file":              carapace.ActionFiles(),
		"datadir":                 carapace.ActionDirectories(),
		"default-both-libraries":  action.ActionBothLibraries(),
		"default-library":         action.ActionDefaultLibraries(),
		"destdir":                 carapace.ActionDirectories(),
		"formats":                 action.ActionDistFormats(),
		"genvslite":               action.ActionGenvslite(),
		"includedir":              carapace.ActionDirectories(),
		"infodir":                 carapace.ActionDirectories(),
		"language":                action.ActionLanguages(),
		"layout":                  action.ActionLayouts(),
		"libdir":                  carapace.ActionDirectories(),
		"libexecdir":              carapace.ActionDirectories(),
		"licensedir":              carapace.ActionDirectories(),
		"localedir":               carapace.ActionDirectories(),
		"localstatedir":           carapace.ActionDirectories(),
		"mandir":                  carapace.ActionDirectories(),
		"namingscheme":            action.ActionNamingSchemes(),
		"native-file":             carapace.ActionFiles(),
		"optimization":            action.ActionOptimizations(),
		"pkg-config-path":         carapace.ActionDirectories(),
		"prefix":                  carapace.ActionDirectories(),
		"sourcedir":               carapace.ActionDirectories(),
		"sysconfdir":              carapace.ActionDirectories(),
		"type":                    action.ActionProjectTypes(),
		"unity":                   action.ActionUnityBuilds(),
		"warnlevel":               action.ActionWarningLevels(),
		"wrap-mode":               action.ActionWrapModes(),
	}

	available := make(carapace.ActionMap)
	for name, a := range actions {
		if cmd.LocalFlags().Lookup(name) != nil {
			available[name] = a
		}
	}
	carapace.Gen(cmd).FlagCompletion(available)
}
