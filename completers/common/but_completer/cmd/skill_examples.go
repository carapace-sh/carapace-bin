package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var skill_examplesCmd = &cobra.Command{
	Use:   "examples",
	Short: "Print the workflow examples",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(skill_examplesCmd).Standalone()

	skill_examplesCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	skillCmd.AddCommand(skill_examplesCmd)
}
