package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mod_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Resolves dependencies in your current Hugo Project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mod_getCmd).Standalone()
	modCmd.AddCommand(mod_getCmd)
}
