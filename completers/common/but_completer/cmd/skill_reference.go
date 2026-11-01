package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var skill_referenceCmd = &cobra.Command{
	Use:   "reference",
	Short: "Print the command reference: syntax and flags for every but command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(skill_referenceCmd).Standalone()

	skill_referenceCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	skillCmd.AddCommand(skill_referenceCmd)
}
