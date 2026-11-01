package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var skills_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the available bundled agent skills. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(skills_listCmd).Standalone()

	skillsCmd.AddCommand(skills_listCmd)
}
