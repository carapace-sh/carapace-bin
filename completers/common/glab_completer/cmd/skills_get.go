package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var skills_getCmd = &cobra.Command{
	Use:   "get <name> [<path>]",
	Short: "Print a bundled agent skill file. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(skills_getCmd).Standalone()

	skillsCmd.AddCommand(skills_getCmd)
}
