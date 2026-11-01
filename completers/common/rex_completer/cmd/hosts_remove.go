package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/rex"
	"github.com/spf13/cobra"
)

var hostsRemoveCmd = &cobra.Command{
	Use:   "remove LABEL",
	Short: "Remove a host",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hostsRemoveCmd).Standalone()

	hostsCmd.AddCommand(hostsRemoveCmd)

	carapace.Gen(hostsRemoveCmd).PositionalCompletion(
		rex.ActionHosts(),
	)
}
