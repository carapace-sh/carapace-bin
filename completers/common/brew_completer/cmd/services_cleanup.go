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

	services_cleanupCmd.Flags().Bool("debug", false, "Display any debugging information.")
	services_cleanupCmd.Flags().Bool("help", false, "Show this message.")
	services_cleanupCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	services_cleanupCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	servicesCmd.AddCommand(services_cleanupCmd)
}
