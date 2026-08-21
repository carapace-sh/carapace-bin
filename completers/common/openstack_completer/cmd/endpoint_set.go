package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var endpoint_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set endpoint properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(endpoint_setCmd).Standalone()

	endpoint_setCmd.Flags().Bool("disable", false, "Disable endpoint")
	endpoint_setCmd.Flags().Bool("enable", false, "Enable endpoint")
	endpoint_setCmd.Flags().String("interface", "", "New endpoint interface type (admin, public or internal)")
	endpoint_setCmd.Flags().String("region", "", "New endpoint region ID")
	endpoint_setCmd.Flags().String("service", "", "New endpoint service (name or ID)")
	endpoint_setCmd.Flags().String("url", "", "New endpoint URL")
	endpointCmd.AddCommand(endpoint_setCmd)
}
