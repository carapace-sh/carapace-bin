package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deleteResourcesCmd = &cobra.Command{
	Use:   "delete-resources",
	Short: "delete resources from the state file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deleteResourcesCmd).Standalone()
	rootCmd.AddCommand(deleteResourcesCmd)
}
