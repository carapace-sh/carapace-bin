package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/systemd-analyze_completer/cmd/action"
	"github.com/spf13/cobra"
)

var capabilityCmd = &cobra.Command{
	Use:   "capability",
	Short: "List capability definitions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(capabilityCmd).Standalone()

	rootCmd.AddCommand(capabilityCmd)

	carapace.Gen(capabilityCmd).PositionalAnyCompletion(
		action.ActionCapabilities(),
	)
}
