package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tsh"
	"github.com/spf13/cobra"
)

var beams_sshCmd = &cobra.Command{
	Use:     "ssh",
	Short:   "Start an interactive shell in a beam, via SSH.",
	Aliases: []string{"console"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(beams_sshCmd).Standalone()

	beamsCmd.AddCommand(beams_sshCmd)

	carapace.Gen(beams_sshCmd).PositionalCompletion(
		tsh.ActionBeams(),
	)
}
