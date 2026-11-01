package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/rex"
	"github.com/spf13/cobra"
)

var hostsEditCmd = &cobra.Command{
	Use:   "edit LABEL",
	Short: "Change a host",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hostsEditCmd).Standalone()

	hostsEditCmd.Flags().String("endpoint", "", "reach the host at ENDPOINT (repeatable, replaces all endpoints)")
	hostsEditCmd.Flags().String("icon", "", "show the icon NAME next to the host (empty removes it)")
	hostsEditCmd.Flags().String("label", "", "rename the host to NEW")

	hostsCmd.AddCommand(hostsEditCmd)

	carapace.Gen(hostsEditCmd).PositionalCompletion(
		rex.ActionHosts(),
	)
}
