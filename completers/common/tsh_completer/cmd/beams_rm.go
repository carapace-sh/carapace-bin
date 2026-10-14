package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var beams_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Delete a beam.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(beams_rmCmd).Standalone()

	beamsCmd.AddCommand(beams_rmCmd)
}
