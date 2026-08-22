package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/brew"
	"github.com/spf13/cobra"
)

var services_killCmd = &cobra.Command{
	Use:     "kill",
	Aliases: []string{"k"},
	Short:   "Stop the service <formula> immediately but keep it registered to launch at login (or boot)",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(services_killCmd).Standalone()

	services_killCmd.Flags().Bool("all", false, "Stop all services immediately but keep them registered to launch at login (or boot).")
	services_killCmd.Flags().Bool("debug", false, "Display any debugging information.")
	services_killCmd.Flags().Bool("help", false, "Show this message.")
	services_killCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	services_killCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	servicesCmd.AddCommand(services_killCmd)

	carapace.Gen(services_killCmd).PositionalAnyCompletion(
		brew.ActionServices().FilterArgs(),
	)
}
