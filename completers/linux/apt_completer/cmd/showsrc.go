package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/apt"
	"github.com/spf13/cobra"
)

var showsrcCmd = &cobra.Command{
	Use:   "showsrc [pattern]...",
	Short: "showsrc package source",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showsrcCmd).Standalone()

	showsrcCmd.Flags().Bool("only-source", false, "display only source package names")
	rootCmd.AddCommand(showsrcCmd)

	carapace.Gen(showsrcCmd).PositionalAnyCompletion(
		apt.ActionPackageSearch(),
	)
}
