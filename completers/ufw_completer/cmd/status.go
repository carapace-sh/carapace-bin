package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "show firewall status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statusCmd).Standalone()

	carapace.Gen(statusCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"numbered", "show firewall status as numbered list of RULES",
			"verbose", "show verbose firewall status",
		))

	rootCmd.AddCommand(statusCmd)
}
