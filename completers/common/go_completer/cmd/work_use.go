package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var work_useCmd = &cobra.Command{
	Use:   "use",
	Short: "add modules to workspace file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(work_useCmd).Standalone()
	work_useCmd.Flags().SetInterspersed(false)

	work_useCmd.Flags().BoolS("r", "r", false, "recursively for modules in the argument directories")
	workCmd.AddCommand(work_useCmd)

	carapace.Gen(work_useCmd).PositionalAnyCompletion(
		carapace.ActionDirectories(),
	)
}
