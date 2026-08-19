package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var DeleteResourcesCmd = &cobra.Command{
	Use:   "delete-resources",
	Short: "Delete Resources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(DeleteResourcesCmd).Standalone()
	rootCmd.AddCommand(DeleteResourcesCmd)
}
