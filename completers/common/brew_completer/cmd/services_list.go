package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var services_listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List information about all managed services for the current user (or root)",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(services_listCmd).Standalone()

	services_listCmd.Flags().Bool("debug", false, "Display any debugging information.")
	services_listCmd.Flags().Bool("help", false, "Show this message.")
	services_listCmd.Flags().Bool("json", false, "Output as JSON.")
	services_listCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	services_listCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	servicesCmd.AddCommand(services_listCmd)
}
