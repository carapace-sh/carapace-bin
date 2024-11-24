package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "clean compilation caches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanCmd).Standalone()

	cleanCmd.Flags().Bool("cascade", false, "clean a project and all projects depending on it")
	cleanCmd.Flags().Bool("include-dependencies", false, "clean a project and all its dependencies")
	cleanCmd.Flags().StringSliceP("project", "p", nil, "the projects to clean")
	cleanCmd.Flags().StringSlice("projects", nil, "the projects to clean")
	cleanCmd.Flags().Bool("propagate", false, "clean a project and all its dependencies")
	rootCmd.AddCommand(cleanCmd)

	cleanCmd.MarkFlagsMutuallyExclusive("include-dependencies", "propagate")

	// TODO completion
}
