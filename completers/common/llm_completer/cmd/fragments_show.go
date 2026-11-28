package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fragments_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Display the fragment stored under an alias or hash",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fragments_showCmd).Standalone()

	fragmentsCmd.AddCommand(fragments_showCmd)
}
