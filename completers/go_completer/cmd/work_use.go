package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var work_useCmd = &cobra.Command{
	Use:   "use",
	Short: "add modules to workspace file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(work_useCmd).Standalone()
	work_useCmd.Flags().BoolS("r", "r", false, "recursively for modules in the argument directories")

	work_useCmd.Flags().SetInterspersed(false)
	workCmd.AddCommand(work_useCmd)

	carapace.Gen(work_useCmd).PositionalAnyCompletion(
		carapace.ActionDirectories(),
	)
}
