package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dmgCmd = &cobra.Command{
	Use:   "dmg",
	Short: "Create a disk image containing a macOS package of a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dmgCmd).Standalone()
	rootCmd.AddCommand(dmgCmd)
}
