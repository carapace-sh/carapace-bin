package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var integrateCmd = &cobra.Command{
	Use:    "integrate",
	Short:  "integrate with an IDE",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(integrateCmd).Standalone()

	rootCmd.AddCommand(integrateCmd)
}
