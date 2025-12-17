package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var alpha_kubercCmd = &cobra.Command{
	Use:   "kuberc SUBCOMMAND",
	Short: "Manage kuberc configuration files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alpha_kubercCmd).Standalone()

	alphaCmd.AddCommand(alpha_kubercCmd)
}
