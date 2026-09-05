package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/brew"
	"github.com/spf13/cobra"
)

var services_restartCmd = &cobra.Command{
	Use:     "restart",
	Aliases: []string{"relaunch", "reload", "r"},
	Short:   "Stop (if necessary) and start the service <formula> immediately and register it to launch at login (or boot)",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(services_restartCmd).Standalone()

	services_restartCmd.Flags().Bool("all", false, "Restart all services.")
	services_restartCmd.Flags().Bool("debug", false, "Display any debugging information.")
	services_restartCmd.Flags().String("file", "", "Use the service file from this location to `start` the service.")
	services_restartCmd.Flags().Bool("help", false, "Show this message.")
	services_restartCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	services_restartCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	servicesCmd.AddCommand(services_restartCmd)

	carapace.Gen(services_restartCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})

	carapace.Gen(services_restartCmd).PositionalAnyCompletion(
		brew.ActionServices().FilterArgs(),
	)
}
