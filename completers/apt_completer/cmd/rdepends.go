package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/apt"
	"github.com/spf13/cobra"
)

var rdependsCmd = &cobra.Command{
	Use:   "rdepends [pattern]...",
	Short: "list package dependents",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rdependsCmd).Standalone()

	// TODO flag descriptions
	rdependsCmd.Flags().Bool("breaks", false, "")
	rdependsCmd.Flags().Bool("conflicts", false, "")
	rdependsCmd.Flags().Bool("enhances", false, "")
	rdependsCmd.Flags().Bool("implicit", false, "")
	rdependsCmd.Flags().BoolP("important", "i", false, "")
	rdependsCmd.Flags().Bool("installed", false, "")
	rdependsCmd.Flags().Bool("pre-rdepends", false, "")
	rdependsCmd.Flags().Bool("rdepends", false, "")
	rdependsCmd.Flags().Bool("recommends", false, "")
	rdependsCmd.Flags().Bool("recurse", false, "")
	rdependsCmd.Flags().Bool("replaces", false, "")
	rdependsCmd.Flags().Bool("suggests", false, "")
	rootCmd.AddCommand(rdependsCmd)

	carapace.Gen(rdependsCmd).PositionalAnyCompletion(
		apt.ActionPackageSearch(),
	)
}
