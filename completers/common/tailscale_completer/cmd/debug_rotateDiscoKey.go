package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_rotateDiscoKeyCmd = &cobra.Command{
	Use:   "rotate-disco-key",
	Short: "Rotate the discovery key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_rotateDiscoKeyCmd).Standalone()

	debugCmd.AddCommand(debug_rotateDiscoKeyCmd)
}
