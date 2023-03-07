package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var experimental_globalsCmd = &cobra.Command{
	Use:   "globals",
	Short: "List globals for all stacks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(experimental_globalsCmd).Standalone()

	experimentalCmd.AddCommand(experimental_globalsCmd)
}
