package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_help_detailsCmd = &cobra.Command{
	Use:   "details",
	Short: "Provide details about given branches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_detailsCmd).Standalone()

	branch_helpCmd.AddCommand(branch_help_detailsCmd)
}
