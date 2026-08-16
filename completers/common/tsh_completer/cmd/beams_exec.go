package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var beams_execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Run a command in a beam, via SSH.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(beams_execCmd).Standalone()

	beamsCmd.AddCommand(beams_execCmd)
}
