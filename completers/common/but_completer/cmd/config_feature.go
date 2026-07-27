package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_featureCmd = &cobra.Command{
	Use:   "feature",
	Short: "View and configure feature flags",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_featureCmd).Standalone()

	config_featureCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	configCmd.AddCommand(config_featureCmd)

	carapace.Gen(config_featureCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"unapply-v3-pgm", "Use the V3 unapply compatibility mode",
			"single-branch", "Enable single-branch mode",
		),
		carapace.ActionValues("enable", "disable"),
	)
}
