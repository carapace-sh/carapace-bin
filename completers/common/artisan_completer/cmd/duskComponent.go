package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var duskComponentCmd = &cobra.Command{
	Use:   "dusk:component <name>",
	Short: "Create a new Dusk component class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(duskComponentCmd).Standalone()

	rootCmd.AddCommand(duskComponentCmd)
}
