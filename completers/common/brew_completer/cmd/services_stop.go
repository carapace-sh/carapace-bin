package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/brew"
	"github.com/spf13/cobra"
)

var services_stopCmd = &cobra.Command{
	Use:     "stop",
	Aliases: []string{"unload", "terminate", "term", "t", "u"},
	Short:   "Stop the service <formula> immediately and unregister it from launching at login (or boot)",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(services_stopCmd).Standalone()

	services_stopCmd.Flags().Bool("all", false, "Stop all services and unregister them from launching at login (or boot), unless `--keep` is specified.")
	services_stopCmd.Flags().Bool("debug", false, "Display any debugging information.")
	services_stopCmd.Flags().Bool("help", false, "Show this message.")
	services_stopCmd.Flags().Bool("keep", false, "When stopped, don't unregister the service from launching at login (or boot).")
	services_stopCmd.Flags().String("max-wait", "", "Wait at most this many seconds for `stop` to finish stopping a service. Defaults to 60. Set this to zero (0) seconds to wait indefinitely.")
	services_stopCmd.Flags().Bool("no-wait", false, "Don't wait for `stop` to finish stopping the service.")
	services_stopCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	services_stopCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	servicesCmd.AddCommand(services_stopCmd)

	carapace.Gen(services_stopCmd).PositionalAnyCompletion(
		brew.ActionServices().FilterArgs(),
	)
}
