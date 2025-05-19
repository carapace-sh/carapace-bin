package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mergeAvailCmd = &cobra.Command{
	Use:   "merge-avail",
	Short: "Replace available packages info",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mergeAvailCmd).Standalone()
}
