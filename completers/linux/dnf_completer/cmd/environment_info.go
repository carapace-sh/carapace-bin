package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var environment_infoCmd = &cobra.Command{
	Use:   "info [environment-spec]...",
	Short: "print detailed information about environments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(environment_infoCmd).Standalone()

	environment_infoCmd.Flags().Bool("available", false, "show only available environments")
	environment_infoCmd.Flags().Bool("installed", false, "show only installed environments")
	environmentCmd.AddCommand(environment_infoCmd)
}
