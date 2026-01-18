package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/flatpak"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:     "run [OPTION…] APP [ARGUMENT…]",
	Short:   "Run an app",
	GroupID: "run",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()
	runCmd.Flags().SetInterspersed(false)

	runCmd.Flags().Bool("a11y-bus", false, "Proxy accessibility bus calls (default except when sandboxed)")
	runCmd.Flags().String("add-policy", "", "Add generic policy option")
	runCmd.Flags().String("allow", "", "Allow feature")
	runCmd.Flags().String("app-path", "", "Use PATH instead of the app's /app")
	runCmd.Flags().String("arch", "", "Arch to use")
	runCmd.Flags().String("branch", "", "Branch to use")
	runCmd.Flags().String("command", "", "Command to run")
	runCmd.Flags().Bool("commit", false, "Run specified commit")
	runCmd.Flags().String("cwd", "", "Directory to run the command in")
	runCmd.Flags().BoolP("devel", "d", false, "Use development runtime")
	runCmd.Flags().String("device", "", "Expose device to app")
	runCmd.Flags().BoolP("die-with-parent", "p", false, "Kill processes when the parent process dies")
	runCmd.Flags().String("disallow", "", "Don't allow feature")
	runCmd.Flags().String("env", "", "Set environment variable")
	runCmd.Flags().String("env-fd", "", "Read environment variables in env -0 format from FD")
	runCmd.Flags().Bool("file-forwarding", false, "Enable file forwarding")
	runCmd.Flags().String("filesystem", "", "Expose filesystem to app (:ro for read-only)")
	runCmd.Flags().BoolP("help", "h", false, "Show help options")
	runCmd.Flags().Bool("help-all", false, "Show all help options")
	runCmd.Flags().String("installation", "", "Work on a non-default system-wide installation")
	runCmd.Flags().Bool("instance-id-fd", false, "Write the instance ID to the given file descriptor")
	runCmd.Flags().Bool("log-a11y-bus", false, "Log accessibility bus calls")
	runCmd.Flags().Bool("log-session-bus", false, "Log session bus calls")
	runCmd.Flags().Bool("log-system-bus", false, "Log system bus calls")
	runCmd.Flags().Bool("no-a11y-bus", false, "Don't proxy accessibility bus calls")
	runCmd.Flags().Bool("no-documents-portal", false, "Don't start portals")
	runCmd.Flags().Bool("no-session-bus", false, "Don't proxy session bus calls")
	runCmd.Flags().String("no-talk-name", "", "Don't allow app to talk to name on the session bus")
	runCmd.Flags().String("nodevice", "", "Don't expose device to app")
	runCmd.Flags().String("nofilesystem", "", "Don't expose filesystem to app")
	runCmd.Flags().String("nosocket", "", "Don't expose socket to app")
	runCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	runCmd.Flags().String("own-name", "", "Allow app to own name on the session bus")
	runCmd.Flags().Bool("parent-expose-pids", false, "Make processes visible in parent namespace")
	runCmd.Flags().String("parent-pid", "", "Use PID as parent pid for sharing namespaces")
	runCmd.Flags().Bool("parent-share-pids", false, "Share process ID namespace with parent")
	runCmd.Flags().String("persist", "", "Persist home directory subpath")
	runCmd.Flags().String("remove-policy", "", "Remove generic policy option")
	runCmd.Flags().String("runtime", "", "Runtime to use")
	runCmd.Flags().Bool("runtime-commit", false, "Use specified runtime commit")
	runCmd.Flags().String("runtime-version", "", "Runtime version to use")
	runCmd.Flags().Bool("sandbox", false, "Run completely sandboxed")
	runCmd.Flags().Bool("session-bus", false, "Proxy session bus calls (default except when sandboxed)")
	runCmd.Flags().String("share", "", "Share with host")
	runCmd.Flags().String("socket", "", "Expose socket to app")
	runCmd.Flags().Bool("system", false, "Work on the system-wide installation (default)")
	runCmd.Flags().String("system-no-talk-name", "", "Don't allow app to talk to name on the system bus")
	runCmd.Flags().String("system-own-name", "", "Allow app to own name on the system bus")
	runCmd.Flags().String("system-talk-name", "", "Allow app to talk to name on the system bus")
	runCmd.Flags().String("talk-name", "", "Allow app to talk to name on the session bus")
	runCmd.Flags().String("unset-env", "", "Remove variable from environment")
	runCmd.Flags().String("unshare", "", "Unshare with host")
	runCmd.Flags().BoolP("user", "u", false, "Work on the user installation")
	runCmd.Flags().String("usr-path", "", "Use PATH instead of the runtime's /usr")
	runCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(runCmd)

	// TODO flags
	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		// 	"add-policy":          carapace.ActionValues(),
		// 	"allow":               carapace.ActionValues(),
		// 	"app-path":            carapace.ActionValues(),
		"arch": flatpak.ActionArches(),
		// 	"branch":              carapace.ActionValues(),
		// 	"command":             carapace.ActionValues(),
		// 	"cwd":                 carapace.ActionValues(),
		// 	"device":              carapace.ActionValues(),
		// 	"disallow":            carapace.ActionValues(),
		"env": env.ActionNameValues(false),
		// 	"env-fd":              carapace.ActionValues(),
		// 	"filesystem":          carapace.ActionValues(),
		// 	"installation":        carapace.ActionValues(),
		// 	"no-talk-name":        carapace.ActionValues(),
		// 	"nodevice":            carapace.ActionValues(),
		// 	"nofilesystem":        carapace.ActionValues(),
		// 	"nosocket":            carapace.ActionValues(),
		// 	"own-name":            carapace.ActionValues(),
		// 	"parent-pid":          carapace.ActionValues(),
		// 	"persist":             carapace.ActionValues(),
		// 	"remove-policy":       carapace.ActionValues(),
		// 	"runtime":             carapace.ActionValues(),
		// 	"runtime-version":     carapace.ActionValues(),
		// 	"share":               carapace.ActionValues(),
		// 	"socket":              carapace.ActionValues(),
		// 	"system-no-talk-name": carapace.ActionValues(),
		// 	"system-own-name":     carapace.ActionValues(),
		// 	"system-talk-name":    carapace.ActionValues(),
		// 	"talk-name":           carapace.ActionValues(),
		// 	"unset-env":           carapace.ActionValues(),
		// 	"unshare":             carapace.ActionValues(),
		// 	"usr-path":            carapace.ActionValues(),
	})

	carapace.Gen(runCmd).PositionalCompletion(
		flatpak.ActionApplications(),
	)

	// TODO complete arguments for application(c.ARGS[0]) but extend path first so that the flatpak binary is executed in any subsequent dynamic completions
}
