package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dbMonitorCmd = &cobra.Command{
	Use:   "db:monitor [--databases [DATABASES]] [--max [MAX]]",
	Short: "Monitor the number of connections on the specified database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dbMonitorCmd).Standalone()

	dbMonitorCmd.Flags().String("databases", "", "The database connections to monitor")
	dbMonitorCmd.Flags().String("max", "", "The maximum number of connections that can be open before an event is dispatched")
	rootCmd.AddCommand(dbMonitorCmd)
}
