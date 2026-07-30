package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var descriptionCmd = &cobra.Command{
	Use:   "description",
	Short: "set the description for a service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(descriptionCmd).Standalone()
	rootCmd.AddCommand(descriptionCmd)
}
