package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tap_mirrorCmd = &cobra.Command{
	Use:   "mirror",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tap_mirrorCmd).Standalone()

	tapCmd.AddCommand(tap_mirrorCmd)
}
