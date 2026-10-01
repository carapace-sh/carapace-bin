package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/brew"
	"github.com/spf13/cobra"
)

var services_startCmd = &cobra.Command{
	Use:     "start",
	Aliases: []string{"launch", "load", "s", "l"},
	Short:   "Start the service <formula> immediately and register it to launch at login (or boot)",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(services_startCmd).Standalone()

	services_startCmd.Flags().Bool("all", false, "Start all services and register them to launch at login (or boot).")
	services_startCmd.Flags().Bool("debug", false, "Display any debugging information.")
	services_startCmd.Flags().String("file", "", "Use the service file from this location to `start` the service.")
	services_startCmd.Flags().Bool("help", false, "Show this message.")
	services_startCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	services_startCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	servicesCmd.AddCommand(services_startCmd)

	carapace.Gen(services_startCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})

	carapace.Gen(services_startCmd).PositionalAnyCompletion(
		brew.ActionServices().FilterArgs(),
	)
}
