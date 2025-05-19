package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var alphaCmd = &cobra.Command{
	Use:    "alpha [COMMAND]",
	Short:  "Experimental commands",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alphaCmd).Standalone()

	rootCmd.AddCommand(alphaCmd)
}
