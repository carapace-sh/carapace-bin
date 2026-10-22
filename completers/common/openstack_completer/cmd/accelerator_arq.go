package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var accelerator_arqCmd = &cobra.Command{
	Use:   "arq",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(accelerator_arqCmd).Standalone()

	acceleratorCmd.AddCommand(accelerator_arqCmd)
}
