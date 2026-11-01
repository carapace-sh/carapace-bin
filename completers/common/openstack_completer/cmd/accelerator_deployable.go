package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accelerator_deployableCmd = &cobra.Command{
	Use:   "deployable",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accelerator_deployableCmd).Standalone()

	acceleratorCmd.AddCommand(accelerator_deployableCmd)
}
