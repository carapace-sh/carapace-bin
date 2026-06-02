package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var updateAvailCmd = &cobra.Command{
	Use:   "update-avail",
	Short: "Replace available packages info",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateAvailCmd).Standalone()
}
