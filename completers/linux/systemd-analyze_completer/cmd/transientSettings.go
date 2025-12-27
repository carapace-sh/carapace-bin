package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/systemctl"
	"github.com/spf13/cobra"
)

var transientSettingsCmd = &cobra.Command{
	Use:   "transient-settings",
	Short: "List transient settings for unit TYPE",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transientSettingsCmd).Standalone()

	rootCmd.AddCommand(transientSettingsCmd)

	carapace.Gen(transientSettingsCmd).PositionalAnyCompletion(
		systemctl.ActionUnitTypes(),
	)
}
