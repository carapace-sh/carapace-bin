package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/brew"
	"github.com/spf13/cobra"
)

var services_infoCmd = &cobra.Command{
	Use:     "info",
	Aliases: []string{"i"},
	Short:   "List all managed services for the current user (or root)",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(services_infoCmd).Standalone()

	services_infoCmd.Flags().Bool("all", false, "List all managed services.")
	services_infoCmd.Flags().Bool("debug", false, "Display any debugging information.")
	services_infoCmd.Flags().Bool("help", false, "Show this message.")
	services_infoCmd.Flags().Bool("json", false, "Output as JSON.")
	services_infoCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	services_infoCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	servicesCmd.AddCommand(services_infoCmd)

	carapace.Gen(services_infoCmd).PositionalAnyCompletion(
		brew.ActionServices().FilterArgs(),
	)
}
