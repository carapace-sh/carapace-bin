package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tap_mirror_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a tap mirror.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tap_mirror_deleteCmd).Standalone()

	tap_mirrorCmd.AddCommand(tap_mirror_deleteCmd)
}
