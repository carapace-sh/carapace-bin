package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var services_cleanupCmd = &cobra.Command{
	Use:     "cleanup",
	Aliases: []string{"clean", "cl", "rm"},
	Short:   "Remove all unused services",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(services_cleanupCmd).Standalone()

	services_cleanupCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	services_cleanupCmd.Flags().BoolP("help", "h", false, "Show this message.")
	services_cleanupCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	services_cleanupCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	servicesCmd.AddCommand(services_cleanupCmd)
}
