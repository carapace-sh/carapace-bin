package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var service_startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start service for a repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(service_startCmd).Standalone()

	service_startCmd.Flags().BoolP("help", "h", false, "Print help")
	serviceCmd.AddCommand(service_startCmd)
}
