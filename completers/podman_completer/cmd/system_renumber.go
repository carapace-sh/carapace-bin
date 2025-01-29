package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var system_renumberCmd = &cobra.Command{
	Use:   "renumber",
	Short: "Migrate lock numbers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(system_renumberCmd).Standalone()

	systemCmd.AddCommand(system_renumberCmd)
}
