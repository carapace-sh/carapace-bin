package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start up the launchservicesd server in process",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverCmd).Standalone()
	serverCmd.Flags().String("duration", "", "Terminate server after given seconds")
	serverCmd.Flags().String("file", "", "Terminate when file at path exists")
	serverCmd.Flags().Bool("forever", false, "Never terminate")
	serverCmd.Flags().String("gone", "", "Terminate when file at path is deleted")
	serverCmd.Flags().Bool("local", false, "Process XPC requests from future commands")
	serverCmd.Flags().String("xpcservicename", "", "XPC service name")
	rootCmd.AddCommand(serverCmd)
}
