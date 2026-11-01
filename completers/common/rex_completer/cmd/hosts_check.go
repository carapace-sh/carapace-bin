package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/rex"
	"github.com/spf13/cobra"
)

var hostsCheckCmd = &cobra.Command{
	Use:   "check [LABEL...]",
	Short: "Check which hosts can be reached from here",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hostsCheckCmd).Standalone()

	hostsCheckCmd.Flags().Bool("json", false, "output the results as JSON")

	hostsCmd.AddCommand(hostsCheckCmd)

	carapace.Gen(hostsCheckCmd).PositionalAnyCompletion(
		rex.ActionHosts(),
	)
}
