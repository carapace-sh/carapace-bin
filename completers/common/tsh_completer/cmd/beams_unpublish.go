package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var beams_unpublishCmd = &cobra.Command{
	Use:   "unpublish",
	Short: "Unpublish a previously published service in a beam.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(beams_unpublishCmd).Standalone()

	beamsCmd.AddCommand(beams_unpublishCmd)
}
