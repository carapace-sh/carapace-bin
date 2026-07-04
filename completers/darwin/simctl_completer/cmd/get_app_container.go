package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/simctl"
	"github.com/spf13/cobra"
)

var getAppContainerCmd = &cobra.Command{
	Use:   "get_app_container",
	Short: "Print the path of the installed app's container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getAppContainerCmd).Standalone()
	rootCmd.AddCommand(getAppContainerCmd)
	carapace.Gen(getAppContainerCmd).PositionalCompletion(
		simctl.ActionDevices(),
	)
	carapace.Gen(getAppContainerCmd).PositionalAnyCompletion(
		carapace.ActionValues("app", "data", "groups"),
	)
}
