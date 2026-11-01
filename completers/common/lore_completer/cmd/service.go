package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Manage the repository in a service process",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serviceCmd).Standalone()

	serviceCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(serviceCmd)
}
