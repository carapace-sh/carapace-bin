package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var maintenance_startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start running maintenance on the current repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(maintenance_startCmd).Standalone()

	maintenance_startCmd.Flags().String("scheduler", "", "specify scheduler")
	maintenanceCmd.AddCommand(maintenance_startCmd)

	carapace.Gen(maintenance_startCmd).FlagCompletion(carapace.ActionMap{
		"scheduler": carapace.ActionValuesDescribed(
			"auto", "",
			"crontab", "POSIX",
			"systemd-timer", "Linux",
			"launchctl", "macOS",
			"schtasks", "Windows",
		),
	})
}
