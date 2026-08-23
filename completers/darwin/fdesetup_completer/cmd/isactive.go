package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var isactiveCmd = &cobra.Command{
	Use:   "isactive",
	Short: "Check if the volume is encrypted",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(isactiveCmd).Standalone()
}
