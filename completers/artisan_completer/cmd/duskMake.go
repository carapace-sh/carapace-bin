package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var duskMakeCmd = &cobra.Command{
	Use:   "dusk:make <name>",
	Short: "Create a new Dusk test class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(duskMakeCmd).Standalone()

	rootCmd.AddCommand(duskMakeCmd)
}
