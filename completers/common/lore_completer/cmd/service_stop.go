package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var service_stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop service for a repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(service_stopCmd).Standalone()

	service_stopCmd.Flags().BoolP("help", "h", false, "Print help")
	serviceCmd.AddCommand(service_stopCmd)
}
