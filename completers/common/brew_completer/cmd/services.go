package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var servicesCmd = &cobra.Command{
	Use:     "services",
	Short:   "Manage background services with macOS' `launchctl`(1) daemon manager or Linux's `systemctl`(1) service manager",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicesCmd).Standalone()

	servicesCmd.PersistentFlags().String("sudo-service-user", "", "When run as root on macOS, run the service(s) as this user.")

	servicesCmd.Flags().Bool("debug", false, "Display any debugging information.")
	servicesCmd.Flags().Bool("help", false, "Show this message.")
	servicesCmd.Flags().Bool("json", false, "Output as JSON.")
	servicesCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	servicesCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(servicesCmd)

	carapace.Gen(servicesCmd).FlagCompletion(carapace.ActionMap{
		"sudo-service-user": os.ActionUsers(),
	})
}
