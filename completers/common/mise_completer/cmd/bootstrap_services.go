package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_servicesCmd = &cobra.Command{
	Use:   "services",
	Short: "Manage services from `[bootstrap.services]`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_servicesCmd).Standalone()

	bootstrap_servicesCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrapCmd.AddCommand(bootstrap_servicesCmd)
}
