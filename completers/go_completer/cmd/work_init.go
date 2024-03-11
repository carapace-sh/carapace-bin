package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var work_initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize workspace file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(work_initCmd).Standalone()
	work_initCmd.Flags().SetInterspersed(false)

	workCmd.AddCommand(work_initCmd)

	carapace.Gen(work_initCmd).PositionalAnyCompletion(
		carapace.ActionDirectories(),
	)
}
