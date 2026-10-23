package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mod_initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize new module",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mod_initCmd).Standalone()

	mod_initCmd.Flags().String("version", "", "init module version")
	modCmd.AddCommand(mod_initCmd)
}
