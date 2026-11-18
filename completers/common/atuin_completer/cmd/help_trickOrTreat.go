package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_trickOrTreatCmd = &cobra.Command{
	Use:    "trick-or-treat",
	Short:  "Dig a forgotten command, or a ghost, out of your history",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_trickOrTreatCmd).Standalone()

	helpCmd.AddCommand(help_trickOrTreatCmd)
}
