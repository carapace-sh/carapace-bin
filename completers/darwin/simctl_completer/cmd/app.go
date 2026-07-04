package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/simctl"
	"github.com/spf13/cobra"
)

var getAppContainerCmd = &cobra.Command{
	Use:   "get_app_container",
	Short: "Print the path of the installed app's container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install an application on a device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var launchCmd = &cobra.Command{
	Use:   "launch",
	Short: "Launch an application by identifier on a device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var terminateCmd = &cobra.Command{
	Use:   "terminate",
	Short: "Terminate an application by identifier on a device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall an application from a device",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getAppContainerCmd).Standalone()
	rootCmd.AddCommand(getAppContainerCmd)
	carapace.Gen(getAppContainerCmd).PositionalCompletion(
		simctl.ActionDevices(),
	)
	carapace.Gen(getAppContainerCmd).PositionalAnyCompletion(
		carapace.ActionValues("app", "data", "groups"),
	)

	carapace.Gen(installCmd).Standalone()
	rootCmd.AddCommand(installCmd)
	carapace.Gen(installCmd).PositionalCompletion(
		simctl.ActionDevices(),
		carapace.ActionDirectories(),
	)

	carapace.Gen(launchCmd).Standalone()
	launchCmd.Flags().Bool("console-pty", false, "Connect stdout/stderr of the app to the terminal")
	rootCmd.AddCommand(launchCmd)
	carapace.Gen(launchCmd).PositionalCompletion(
		simctl.ActionDevices(),
	)

	carapace.Gen(terminateCmd).Standalone()
	rootCmd.AddCommand(terminateCmd)
	carapace.Gen(terminateCmd).PositionalCompletion(
		simctl.ActionDevices(),
	)

	carapace.Gen(uninstallCmd).Standalone()
	rootCmd.AddCommand(uninstallCmd)
	carapace.Gen(uninstallCmd).PositionalCompletion(
		simctl.ActionDevices(),
	)
}
