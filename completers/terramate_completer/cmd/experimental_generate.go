package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var experimental_generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Experimental generate commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(experimental_generateCmd).Standalone()

	experimentalCmd.AddCommand(experimental_generateCmd)
}
