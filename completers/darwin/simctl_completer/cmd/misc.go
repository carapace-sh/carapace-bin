package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/simctl"
	"github.com/spf13/cobra"
)

var addmediaCmd = &cobra.Command{
	Use:   "addmedia",
	Short: "Add photos, videos, or contacts to a device's library",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var diagnoseCmd = &cobra.Command{
	Use:   "diagnose",
	Short: "Collect diagnostic information and logs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var getenvCmd = &cobra.Command{
	Use:   "getenv",
	Short: "Print an environment variable from a running device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var ioCmd = &cobra.Command{
	Use:   "io",
	Short: "Set up a device IO operation (screenshot, recordVideo)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var icloudSyncCmd = &cobra.Command{
	Use:   "icloud_sync",
	Short: "Trigger iCloud sync on a device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var keychainCmd = &cobra.Command{
	Use:   "keychain",
	Short: "Manipulate a device's keychain",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available devices, device types, runtimes, or device pairs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var logverboseCmd = &cobra.Command{
	Use:   "logverbose",
	Short: "Enable or disable verbose logging for a device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var openurlCmd = &cobra.Command{
	Use:   "openurl",
	Short: "Open a URL in a device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var pairCmd = &cobra.Command{
	Use:   "pair",
	Short: "Create a new watch and phone pair",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var pairActivateCmd = &cobra.Command{
	Use:   "pair_activate",
	Short: "Set a given pair as active",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var pbcopyCmd = &cobra.Command{
	Use:   "pbcopy",
	Short: "Copy standard input onto the device pasteboard",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var pbpasteCmd = &cobra.Command{
	Use:   "pbpaste",
	Short: "Print the contents of the device's pasteboard to stdout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var pbsyncCmd = &cobra.Command{
	Use:   "pbsync",
	Short: "Sync pasteboard content from one device to another",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var privacyCmd = &cobra.Command{
	Use:   "privacy",
	Short: "Grant, revoke, or reset privacy permissions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Send a simulated push notification",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var runtimeCmd = &cobra.Command{
	Use:   "runtime",
	Short: "Manage simulator runtime disk images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var spawnCmd = &cobra.Command{
	Use:   "spawn",
	Short: "Spawn a process by executing a given executable on a device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var statusBarCmd = &cobra.Command{
	Use:   "status_bar",
	Short: "Set or clear status bar overrides",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var uiCmd = &cobra.Command{
	Use:   "ui",
	Short: "Get or set UI options",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var unpairCmd = &cobra.Command{
	Use:   "unpair",
	Short: "Unpair a watch and phone pair",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addmediaCmd).Standalone()
	rootCmd.AddCommand(addmediaCmd)
	carapace.Gen(addmediaCmd).PositionalCompletion(
		simctl.ActionDevices(),
		carapace.ActionFiles(),
	)

	carapace.Gen(diagnoseCmd).Standalone()
	diagnoseCmd.Flags().StringP("log", "l", "", "Write output to a specific path")
	rootCmd.AddCommand(diagnoseCmd)
	carapace.Gen(diagnoseCmd).FlagCompletion(carapace.ActionMap{
		"log": carapace.ActionFiles(),
	})

	carapace.Gen(getenvCmd).Standalone()
	rootCmd.AddCommand(getenvCmd)
	carapace.Gen(getenvCmd).PositionalCompletion(
		simctl.ActionDevicesByState("Booted"),
	)

	carapace.Gen(ioCmd).Standalone()
	rootCmd.AddCommand(ioCmd)
	carapace.Gen(ioCmd).PositionalCompletion(
		simctl.ActionDevices(),
		carapace.ActionValues("screenshot", "recordVideo"),
	)

	carapace.Gen(icloudSyncCmd).Standalone()
	rootCmd.AddCommand(icloudSyncCmd)
	carapace.Gen(icloudSyncCmd).PositionalCompletion(simctl.ActionDevices())

	carapace.Gen(keychainCmd).Standalone()
	rootCmd.AddCommand(keychainCmd)
	carapace.Gen(keychainCmd).PositionalCompletion(
		simctl.ActionDevices(),
		carapace.ActionValues("add-root-cert", "add-cert", "reset"),
	)
	carapace.Gen(keychainCmd).PositionalAnyCompletion(carapace.ActionFiles())

	carapace.Gen(listCmd).Standalone()
	listCmd.Flags().BoolP("json", "j", false, "Produce JSON output")
	rootCmd.AddCommand(listCmd)
	carapace.Gen(listCmd).PositionalCompletion(
		carapace.ActionValues("devices", "devicetypes", "runtimes", "pairs"),
	)

	carapace.Gen(logverboseCmd).Standalone()
	rootCmd.AddCommand(logverboseCmd)
	carapace.Gen(logverboseCmd).PositionalCompletion(
		simctl.ActionDevices(),
		carapace.ActionValues("enable", "disable"),
	)

	carapace.Gen(openurlCmd).Standalone()
	rootCmd.AddCommand(openurlCmd)
	carapace.Gen(openurlCmd).PositionalCompletion(simctl.ActionDevices())

	carapace.Gen(pairCmd).Standalone()
	rootCmd.AddCommand(pairCmd)
	carapace.Gen(pairCmd).PositionalCompletion(
		simctl.ActionDevices(),
		simctl.ActionDevices(),
	)

	carapace.Gen(pairActivateCmd).Standalone()
	rootCmd.AddCommand(pairActivateCmd)
	carapace.Gen(pairActivateCmd).PositionalCompletion(
		carapace.ActionValues(),
	)

	carapace.Gen(pbcopyCmd).Standalone()
	rootCmd.AddCommand(pbcopyCmd)
	carapace.Gen(pbcopyCmd).PositionalCompletion(simctl.ActionDevices())

	carapace.Gen(pbpasteCmd).Standalone()
	rootCmd.AddCommand(pbpasteCmd)
	carapace.Gen(pbpasteCmd).PositionalCompletion(simctl.ActionDevices())

	carapace.Gen(pbsyncCmd).Standalone()
	rootCmd.AddCommand(pbsyncCmd)
	carapace.Gen(pbsyncCmd).PositionalCompletion(
		simctl.ActionDevices(),
		simctl.ActionDevices(),
	)

	carapace.Gen(privacyCmd).Standalone()
	rootCmd.AddCommand(privacyCmd)
	carapace.Gen(privacyCmd).PositionalCompletion(
		simctl.ActionDevices(),
		carapace.ActionValues("grant", "revoke", "reset"),
		carapace.ActionValues(
			"all",
			"calendar",
			"contacts-limited",
			"contacts",
			"location",
			"location-always",
			"photos-add",
			"photos",
			"media-library",
			"microphone",
			"motion",
			"reminders",
			"siri",
		),
	)

	carapace.Gen(pushCmd).Standalone()
	rootCmd.AddCommand(pushCmd)
	carapace.Gen(pushCmd).PositionalCompletion(
		simctl.ActionDevices(),
	)
	carapace.Gen(pushCmd).PositionalAnyCompletion(carapace.ActionFiles())

	carapace.Gen(runtimeCmd).Standalone()
	rootCmd.AddCommand(runtimeCmd)
	carapace.Gen(runtimeCmd).PositionalCompletion(
		carapace.ActionValues("add", "delete"),
	)
	carapace.Gen(runtimeCmd).PositionalAnyCompletion(carapace.ActionFiles())

	carapace.Gen(spawnCmd).Standalone()
	rootCmd.AddCommand(spawnCmd)
	carapace.Gen(spawnCmd).PositionalCompletion(
		simctl.ActionDevicesByState("Booted"),
	)

	carapace.Gen(statusBarCmd).Standalone()
	rootCmd.AddCommand(statusBarCmd)
	carapace.Gen(statusBarCmd).PositionalCompletion(
		simctl.ActionDevices(),
		carapace.ActionValues("list", "clear", "override"),
	)

	carapace.Gen(uiCmd).Standalone()
	rootCmd.AddCommand(uiCmd)
	carapace.Gen(uiCmd).PositionalCompletion(
		simctl.ActionDevices(),
		carapace.ActionValues("appearance"),
		carapace.ActionValues("light", "dark"),
	)

	carapace.Gen(unpairCmd).Standalone()
	rootCmd.AddCommand(unpairCmd)
	carapace.Gen(unpairCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
