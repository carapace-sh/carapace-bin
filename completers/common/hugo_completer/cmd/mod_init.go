package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mod_initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize this project as a Hugo Module.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mod_initCmd).Standalone()
	modCmd.AddCommand(mod_initCmd)
}
