package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var killCmd = &cobra.Command{
	Use:   "kill",
	Short: "Kill the application specified",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(killCmd).Standalone()
	killCmd.Flags().Bool("childapps", false, "Kill child applications")
	killCmd.Flags().Bool("coalition", false, "Kill all pids in the coalition")
	killCmd.Flags().Bool("force", false, "Immediately remove application from list")
	killCmd.Flags().Bool("hard", false, "Send SIGKILL instead of SIGTERM")
	killCmd.Flags().Bool("launchdjobs", false, "Include launchd jobs in coalition")
	rootCmd.AddCommand(killCmd)
}
