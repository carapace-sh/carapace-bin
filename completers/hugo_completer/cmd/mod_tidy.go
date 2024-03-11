package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mod_tidyCmd = &cobra.Command{
	Use:   "tidy",
	Short: "Remove unused entries in go.mod and go.sum.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mod_tidyCmd).Standalone()
	modCmd.AddCommand(mod_tidyCmd)
}
