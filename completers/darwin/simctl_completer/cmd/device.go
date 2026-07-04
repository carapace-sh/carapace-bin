package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/simctl"
	"github.com/spf13/cobra"
)

var bootCmd = &cobra.Command{
	Use:   "boot",
	Short: "Boot a device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone an existing device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete specified devices, unavailable devices, or all devices",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var eraseCmd = &cobra.Command{
	Use:   "erase",
	Short: "Erase a device's contents and settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename a device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var shutdownCmd = &cobra.Command{
	Use:   "shutdown",
	Short: "Shut down a device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade a device to a newer runtime",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootCmd).Standalone()
	rootCmd.AddCommand(bootCmd)
	carapace.Gen(bootCmd).PositionalCompletion(simctl.ActionDevices())

	carapace.Gen(cloneCmd).Standalone()
	rootCmd.AddCommand(cloneCmd)
	carapace.Gen(cloneCmd).PositionalCompletion(
		simctl.ActionDevices(),
	)

	carapace.Gen(createCmd).Standalone()
	rootCmd.AddCommand(createCmd)

	carapace.Gen(deleteCmd).Standalone()
	rootCmd.AddCommand(deleteCmd)
	carapace.Gen(deleteCmd).PositionalAnyCompletion(
		carapace.Batch(
			carapace.ActionValues("all", "unavailable"),
			simctl.ActionDevices(),
		).ToA(),
	)

	carapace.Gen(eraseCmd).Standalone()
	rootCmd.AddCommand(eraseCmd)
	carapace.Gen(eraseCmd).PositionalAnyCompletion(
		carapace.Batch(
			carapace.ActionValues("all"),
			simctl.ActionDevices(),
		).ToA(),
	)

	carapace.Gen(renameCmd).Standalone()
	rootCmd.AddCommand(renameCmd)
	carapace.Gen(renameCmd).PositionalCompletion(
		simctl.ActionDevices(),
	)

	carapace.Gen(shutdownCmd).Standalone()
	rootCmd.AddCommand(shutdownCmd)
	carapace.Gen(shutdownCmd).PositionalAnyCompletion(
		carapace.Batch(
			carapace.ActionValues("all"),
			simctl.ActionDevicesByState("Booted"),
		).ToA(),
	)

	carapace.Gen(upgradeCmd).Standalone()
	rootCmd.AddCommand(upgradeCmd)
	carapace.Gen(upgradeCmd).PositionalCompletion(
		simctl.ActionDevices(),
	)
}
