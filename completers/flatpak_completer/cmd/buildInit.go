package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/flatpak"
	"github.com/spf13/cobra"
)

var buildInitCmd = &cobra.Command{
	Use:     "build-init [OPTION…] DIRECTORY APPNAME SDK RUNTIME [BRANCH]",
	Short:   "Initialize a directory for building",
	GroupID: "build",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildInitCmd).Standalone()

	buildInitCmd.Flags().String("arch", "", "Arch to use")
	buildInitCmd.Flags().String("base", "", "Initialize apps from named app")
	buildInitCmd.Flags().String("base-extension", "", "Include this base extension")
	buildInitCmd.Flags().String("base-version", "", "Specify version for --base")
	buildInitCmd.Flags().String("extension", "", "Add extension point info")
	buildInitCmd.Flags().String("extension-tag", "", "Extension tag to use if building extension")
	buildInitCmd.Flags().BoolP("help", "h", false, "Show help options")
	buildInitCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	buildInitCmd.Flags().String("sdk-dir", "", "Where to store sdk (defaults to 'usr')")
	buildInitCmd.Flags().String("sdk-extension", "", "Include this sdk extension in /usr")
	buildInitCmd.Flags().String("tag", "", "Add a tag")
	buildInitCmd.Flags().String("type", "", "Specify the build type (app, runtime, extension)")
	buildInitCmd.Flags().Bool("update", false, "Re-initialize the sdk/var")
	buildInitCmd.Flags().StringP("var", "v", "", "Initialize var from named runtime")
	buildInitCmd.Flags().Bool("verbose", false, "Show debug information, -vv for more detail")
	buildInitCmd.Flags().BoolP("writable-sdk", "w", false, "Initialize /usr with a writable copy of the sdk")
	rootCmd.AddCommand(buildInitCmd)

	// TODO flags
	carapace.Gen(buildInitCmd).FlagCompletion(carapace.ActionMap{
		"arch": flatpak.ActionArches(),
		// "base":           carapace.ActionValues(),
		// "base-extension": carapace.ActionValues(),
		// "base-version":   carapace.ActionValues(),
		// "extension":      carapace.ActionValues(),
		// "extension-tag":  carapace.ActionValues(),
		// "sdk-dir":        carapace.ActionDirectories(),
		// "sdk-extension":  carapace.ActionValues(),
		// "tag":            carapace.ActionValues(),
		"type": carapace.ActionValues("app", "runtime", "extension"),
		// "var":            carapace.ActionValues(),
	})

	// build-init [OPTION…] DIRECTORY APPNAME SDK RUNTIME [BRANCH]
	carapace.Gen(buildInitCmd).PositionalCompletion(
		carapace.ActionDirectories(),
		carapace.ActionValues(), // TODO appname
		carapace.ActionValues(), // TODO sdk
		carapace.ActionValues(), // TODO runtime
		carapace.ActionValues(), // TODO branch

	)
}
