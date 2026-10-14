package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/brew"
	"github.com/spf13/cobra"
)

var services_runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the service <formula> without registering to launch at login (or boot)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(services_runCmd).Standalone()

	services_runCmd.Flags().Bool("all", false, "Run all services without registering them to launch at login (or boot).")
	services_runCmd.Flags().Bool("debug", false, "Display any debugging information.")
	services_runCmd.Flags().String("file", "", "Use the service file from this location to `run` the service.")
	services_runCmd.Flags().Bool("help", false, "Show this message.")
	services_runCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	services_runCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	servicesCmd.AddCommand(services_runCmd)

	carapace.Gen(services_runCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})

	carapace.Gen(services_runCmd).PositionalAnyCompletion(
		brew.ActionServices().FilterArgs(),
	)
}
