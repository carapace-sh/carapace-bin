package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ip_availabilityCmd = &cobra.Command{
	Use:   "availability",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ip_availabilityCmd).Standalone()

	ipCmd.AddCommand(ip_availabilityCmd)
}
