package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_services_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the state of services from `[bootstrap.services]`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_services_statusCmd).Standalone()

	bootstrap_services_statusCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_services_statusCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	bootstrap_services_statusCmd.Flags().Bool("missing", false, "Exit with status 1 if any service differs from the config (all are still listed)")
	bootstrap_servicesCmd.AddCommand(bootstrap_services_statusCmd)
}
