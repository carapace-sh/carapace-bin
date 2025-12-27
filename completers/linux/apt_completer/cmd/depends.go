package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/apt"
	"github.com/spf13/cobra"
)

var dependsCmd = &cobra.Command{
	Use:   "depends [pattern]...",
	Short: "list package dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dependsCmd).Standalone()

	// TODO flag descriptions
	dependsCmd.Flags().Bool("breaks", false, "")
	dependsCmd.Flags().Bool("conflicts", false, "")
	dependsCmd.Flags().Bool("depends", false, "")
	dependsCmd.Flags().Bool("enhances", false, "")
	dependsCmd.Flags().Bool("implicit", false, "")
	dependsCmd.Flags().BoolP("important", "i", false, "")
	dependsCmd.Flags().Bool("installed", false, "")
	dependsCmd.Flags().Bool("pre-depends", false, "")
	dependsCmd.Flags().Bool("recommends", false, "")
	dependsCmd.Flags().Bool("recurse", false, "")
	dependsCmd.Flags().Bool("replaces", false, "")
	dependsCmd.Flags().Bool("suggests", false, "")
	rootCmd.AddCommand(dependsCmd)

	carapace.Gen(dependsCmd).PositionalAnyCompletion(
		apt.ActionPackageSearch(),
	)
}
