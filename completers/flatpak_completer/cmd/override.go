package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/flatpak"
	"github.com/spf13/cobra"
)

var overrideCmd = &cobra.Command{
	Use:     "override [OPTION…] [APP]",
	Short:   "Override settings [for application]",
	GroupID: "run",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(overrideCmd).Standalone()

	overrideCmd.Flags().String("add-policy", "", "Add generic policy option")
	overrideCmd.Flags().String("allow", "", "Allow feature")
	overrideCmd.Flags().String("device", "", "Expose device to app")
	overrideCmd.Flags().String("disallow", "", "Don't allow feature")
	overrideCmd.Flags().String("env", "", "Set environment variable")
	overrideCmd.Flags().String("env-fd", "", "Read environment variables in env -0 format from FD")
	overrideCmd.Flags().String("filesystem", "", "Expose filesystem to app (:ro for read-only)")
	overrideCmd.Flags().BoolP("help", "h", false, "Show help options")
	overrideCmd.Flags().Bool("help-all", false, "Show all help options")
	overrideCmd.Flags().String("installation", "", "Work on a non-default system-wide installation")
	overrideCmd.Flags().String("no-talk-name", "", "Don't allow app to talk to name on the session bus")
	overrideCmd.Flags().String("nodevice", "", "Don't expose device to app")
	overrideCmd.Flags().String("nofilesystem", "", "Don't expose filesystem to app")
	overrideCmd.Flags().String("nosocket", "", "Don't expose socket to app")
	overrideCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	overrideCmd.Flags().String("own-name", "", "Allow app to own name on the session bus")
	overrideCmd.Flags().String("persist", "", "Persist home directory subpath")
	overrideCmd.Flags().String("remove-policy", "", "Remove generic policy option")
	overrideCmd.Flags().Bool("reset", false, "Remove existing overrides")
	overrideCmd.Flags().String("share", "", "Share with host")
	overrideCmd.Flags().Bool("show", false, "Show existing overrides")
	overrideCmd.Flags().String("socket", "", "Expose socket to app")
	overrideCmd.Flags().Bool("system", false, "Work on the system-wide installation (default)")
	overrideCmd.Flags().String("system-no-talk-name", "", "Don't allow app to talk to name on the system bus")
	overrideCmd.Flags().String("system-own-name", "", "Allow app to own name on the system bus")
	overrideCmd.Flags().String("system-talk-name", "", "Allow app to talk to name on the system bus")
	overrideCmd.Flags().String("talk-name", "", "Allow app to talk to name on the session bus")
	overrideCmd.Flags().String("unset-env", "", "Remove variable from environment")
	overrideCmd.Flags().String("unshare", "", "Unshare with host")
	overrideCmd.Flags().BoolP("user", "u", false, "Work on the user installation")
	overrideCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(overrideCmd)

	// TODO flags
	carapace.Gen(overrideCmd).FlagCompletion(carapace.ActionMap{
		// "add-policy":          carapace.ActionValues(),
		// "allow":               carapace.ActionValues(),
		// "device":              carapace.ActionValues(),
		// "disallow":            carapace.ActionValues(),
		"env": env.ActionNameValues(false),
		// "env-fd":              carapace.ActionValues(),
		// "filesystem":          carapace.ActionValues(),
		// "installation":        carapace.ActionValues(),
		// "no-talk-name":        carapace.ActionValues(),
		// "nodevice":            carapace.ActionValues(),
		// "nofilesystem":        carapace.ActionValues(),
		// "nosocket":            carapace.ActionValues(),
		// "own-name":            carapace.ActionValues(),
		// "persist":             carapace.ActionValues(),
		// "remove-policy":       carapace.ActionValues(),
		// "share":               carapace.ActionValues(),
		// "socket":              carapace.ActionValues(),
		// "system-no-talk-name": carapace.ActionValues(),
		// "system-own-name":     carapace.ActionValues(),
		// "system-talk-name":    carapace.ActionValues(),
		// "talk-name":           carapace.ActionValues(),
		// "unset-env":           carapace.ActionValues(),
		// "unshare":             carapace.ActionValues(),
	})

	carapace.Gen(overrideCmd).PositionalCompletion(
		flatpak.ActionApplications(),
	)
}
