package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:     "build [OPTION…] DIRECTORY [COMMAND [ARGUMENT…]]",
	Short:   "Build in directory",
	GroupID: "build",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	buildCmd.Flags().String("add-policy", "", "Add generic policy option")
	buildCmd.Flags().String("allow", "", "Allow feature")
	buildCmd.Flags().String("bind-mount", "", "Add bind mount")
	buildCmd.Flags().String("build-dir", "", "Start build in this directory")
	buildCmd.Flags().String("device", "", "Expose device to app")
	buildCmd.Flags().BoolP("die-with-parent", "p", false, "Kill processes when the parent process dies")
	buildCmd.Flags().String("disallow", "", "Don't allow feature")
	buildCmd.Flags().String("env", "", "Set environment variable")
	buildCmd.Flags().String("env-fd", "", "Read environment variables in env -0 format from FD")
	buildCmd.Flags().String("filesystem", "", "Expose filesystem to app (:ro for read-only)")
	buildCmd.Flags().BoolP("help", "h", false, "Show help options")
	buildCmd.Flags().Bool("help-all", false, "Show all help options")
	buildCmd.Flags().Bool("log-session-bus", false, "Log session bus calls")
	buildCmd.Flags().Bool("log-system-bus", false, "Log system bus calls")
	buildCmd.Flags().String("metadata", "", "Use alternative file for the metadata")
	buildCmd.Flags().String("no-talk-name", "", "Don't allow app to talk to name on the session bus")
	buildCmd.Flags().String("nodevice", "", "Don't expose device to app")
	buildCmd.Flags().String("nofilesystem", "", "Don't expose filesystem to app")
	buildCmd.Flags().String("nosocket", "", "Don't expose socket to app")
	buildCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	buildCmd.Flags().String("own-name", "", "Allow app to own name on the session bus")
	buildCmd.Flags().String("persist", "", "Persist home directory subpath")
	buildCmd.Flags().Bool("readonly", false, "Make destination readonly")
	buildCmd.Flags().String("remove-policy", "", "Remove generic policy option")
	buildCmd.Flags().BoolP("runtime", "r", false, "Use Platform runtime rather than Sdk")
	buildCmd.Flags().String("sdk-dir", "", "Where to look for custom sdk dir (defaults to 'usr')")
	buildCmd.Flags().String("share", "", "Share with host")
	buildCmd.Flags().String("socket", "", "Expose socket to app")
	buildCmd.Flags().String("system-no-talk-name", "", "Don't allow app to talk to name on the system bus")
	buildCmd.Flags().String("system-own-name", "", "Allow app to own name on the system bus")
	buildCmd.Flags().String("system-talk-name", "", "Allow app to talk to name on the system bus")
	buildCmd.Flags().String("talk-name", "", "Allow app to talk to name on the session bus")
	buildCmd.Flags().String("unset-env", "", "Remove variable from environment")
	buildCmd.Flags().String("unshare", "", "Unshare with host")
	buildCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	buildCmd.Flags().Bool("with-appdir", false, "Export application homedir directory to build")
	rootCmd.AddCommand(buildCmd)

	// TODO flag
	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		// 	"add-policy":          carapace.ActionValues(),
		// 	"allow":               carapace.ActionValues(),
		// 	"bind-mount":          carapace.ActionValues(),
		"build-dir": carapace.ActionDirectories(),
		// 	"device":              carapace.ActionValues(),
		// 	"disallow":            carapace.ActionValues(),
		"env": env.ActionNameValues(false),
		// 	"env-fd":              carapace.ActionValues(),
		// 	"filesystem":          carapace.ActionValues(),
		// 	"metadata":            carapace.ActionValues(),
		// 	"no-talk-name":        carapace.ActionValues(),
		// 	"nodevice":            carapace.ActionValues(),
		// 	"nofilesystem":        carapace.ActionValues(),
		// 	"nosocket":            carapace.ActionValues(),
		// 	"own-name":            carapace.ActionValues(),
		// 	"persist":             carapace.ActionValues(),
		// 	"remove-policy":       carapace.ActionValues(),
		"sdk-dir": carapace.ActionDirectories(),
		// 	"share":               carapace.ActionValues(),
		// 	"socket":              carapace.ActionValues(),
		// 	"system-no-talk-name": carapace.ActionValues(),
		// 	"system-own-name":     carapace.ActionValues(),
		// 	"system-talk-name":    carapace.ActionValues(),
		// 	"talk-name":           carapace.ActionValues(),
		// 	"unset-env":           carapace.ActionValues(),
		// 	"unshare":             carapace.ActionValues(),
	})

	carapace.Gen(buildCmd).PositionalCompletion(
		carapace.ActionDirectories(),
		// TODO command
	)

	// TODO command arguments
}
