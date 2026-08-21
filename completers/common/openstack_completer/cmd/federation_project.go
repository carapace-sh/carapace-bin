package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var federation_projectCmd = &cobra.Command{
	Use:   "project",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(federation_projectCmd).Standalone()

	federationCmd.AddCommand(federation_projectCmd)
}
