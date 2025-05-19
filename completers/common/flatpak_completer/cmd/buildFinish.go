package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/spf13/cobra"
)

var buildFinishCmd = &cobra.Command{
	Use:     "build-finish [OPTION…] DIRECTORY",
	Short:   "Finalize a build directory",
	GroupID: "build",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildFinishCmd).Standalone()

	buildFinishCmd.Flags().String("add-policy", "", "Add generic policy option")
	buildFinishCmd.Flags().String("allow", "", "Allow feature")
	buildFinishCmd.Flags().String("command", "", "Command to set")
	buildFinishCmd.Flags().String("device", "", "Expose device to app")
	buildFinishCmd.Flags().String("disallow", "", "Don't allow feature")
	buildFinishCmd.Flags().String("env", "", "Set environment variable")
	buildFinishCmd.Flags().String("env-fd", "", "Read environment variables in env -0 format from FD")
	buildFinishCmd.Flags().String("extension", "", "Add extension point info")
	buildFinishCmd.Flags().String("extension-priority", "", "Set extension priority (only for extensions)")
	buildFinishCmd.Flags().Bool("extra-data", false, "Extra data info")
	buildFinishCmd.Flags().String("filesystem", "", "Expose filesystem to app (:ro for read-only)")
	buildFinishCmd.Flags().BoolP("help", "h", false, "Show help options")
	buildFinishCmd.Flags().Bool("help-all", false, "Show all help options")
	buildFinishCmd.Flags().String("metadata", "", "Set generic metadata option")
	buildFinishCmd.Flags().Bool("no-exports", false, "Don't process exports")
	buildFinishCmd.Flags().Bool("no-inherit-permissions", false, "Don't inherit permissions from runtime")
	buildFinishCmd.Flags().String("no-talk-name", "", "Don't allow app to talk to name on the session bus")
	buildFinishCmd.Flags().String("nodevice", "", "Don't expose device to app")
	buildFinishCmd.Flags().String("nofilesystem", "", "Don't expose filesystem to app")
	buildFinishCmd.Flags().String("nosocket", "", "Don't expose socket to app")
	buildFinishCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	buildFinishCmd.Flags().String("own-name", "", "Allow app to own name on the session bus")
	buildFinishCmd.Flags().String("persist", "", "Persist home directory subpath")
	buildFinishCmd.Flags().String("remove-extension", "", "Remove extension point info")
	buildFinishCmd.Flags().String("remove-policy", "", "Remove generic policy option")
	buildFinishCmd.Flags().String("require-version", "", "Flatpak version to require")
	buildFinishCmd.Flags().String("runtime", "", "Change the runtime used for the app")
	buildFinishCmd.Flags().String("sdk", "", "Change the sdk used for the app")
	buildFinishCmd.Flags().String("share", "", "Share with host")
	buildFinishCmd.Flags().String("socket", "", "Expose socket to app")
	buildFinishCmd.Flags().String("system-no-talk-name", "", "Don't allow app to talk to name on the system bus")
	buildFinishCmd.Flags().String("system-own-name", "", "Allow app to own name on the system bus")
	buildFinishCmd.Flags().String("system-talk-name", "", "Allow app to talk to name on the system bus")
	buildFinishCmd.Flags().String("talk-name", "", "Allow app to talk to name on the session bus")
	buildFinishCmd.Flags().String("unset-env", "", "Remove variable from environment")
	buildFinishCmd.Flags().String("unshare", "", "Unshare with host")
	buildFinishCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(buildFinishCmd)

	// TODO flags
	carapace.Gen(buildFinishCmd).FlagCompletion(carapace.ActionMap{
		// "add-policy": carapace.ActionValues().
		// "allow": carapace.ActionValues().
		// "command": carapace.ActionValues().
		// "device": carapace.ActionValues().
		// "disallow": carapace.ActionValues().
		"env": env.ActionNameValues(false),
		// "env-fd": carapace.ActionValues().
		// "extension": carapace.ActionValues().
		// "extension-priority": carapace.ActionValues().
		// "filesystem": carapace.ActionValues().
		// "metadata": carapace.ActionValues().
		// "no-talk-name": carapace.ActionValues().
		// "nodevice": carapace.ActionValues().
		// "nofilesystem": carapace.ActionValues().
		// "nosocket": carapace.ActionValues().
		// "own-name": carapace.ActionValues().
		// "persist": carapace.ActionValues().
		// "remove-extension": carapace.ActionValues().
		// "remove-policy": carapace.ActionValues().
		// "require-version": carapace.ActionValues().
		// "runtime": carapace.ActionValues().
		// "sdk": carapace.ActionValues().
		// "share": carapace.ActionValues().
		// "socket": carapace.ActionValues().
		// "system-no-talk-name": carapace.ActionValues().
		// "system-own-name": carapace.ActionValues().
		// "system-talk-name": carapace.ActionValues().
		// "talk-name": carapace.ActionValues().
		// "unset-env": carapace.ActionValues().
		// "unshare"		: carapace.ActionValues().
	})

	carapace.Gen(buildFinishCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
