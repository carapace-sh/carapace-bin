package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var offlineCmd = &cobra.Command{
	Use:   "offline <subcommand> [options]",
	Short: "manage offline transactions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var offlineCleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "remove stored offline transaction and delete cached package files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var offlineLogCmd = &cobra.Command{
	Use:   "log",
	Short: "see logs from offline transactions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var offlineRebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "prepare system to perform offline transaction and reboot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var offlineStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "show status of current offline transaction",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(offlineCmd).Standalone()

	offlineCmd.AddCommand(offlineCleanCmd)
	offlineCmd.AddCommand(offlineLogCmd)
	offlineCmd.AddCommand(offlineRebootCmd)
	offlineCmd.AddCommand(offlineStatusCmd)

	offlineLogCmd.Flags().Int("number", 0, "show the log specified by the boot number")
	offlineRebootCmd.Flags().Bool("poweroff", false, "power off after transaction instead of rebooting")

	rootCmd.AddCommand(offlineCmd)
}
