package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var skill_conceptsCmd = &cobra.Command{
	Use:   "concepts",
	Short: "Print the concepts guide: the workspace model behind but",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(skill_conceptsCmd).Standalone()

	skill_conceptsCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	skillCmd.AddCommand(skill_conceptsCmd)
}
