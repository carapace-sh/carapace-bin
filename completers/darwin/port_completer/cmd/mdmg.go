package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mdmgCmd = &cobra.Command{
	Use:   "mdmg",
	Short: "Create a disk image containing a metapackage of a port and dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mdmgCmd).Standalone()
	rootCmd.AddCommand(mdmgCmd)
}
