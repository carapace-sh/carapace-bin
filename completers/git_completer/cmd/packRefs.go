package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var packRefsCmd = &cobra.Command{
	Use:     "pack-refs",
	Short:   "Pack heads and tags for efficient repository access",
	GroupID: groups[group_manipulator].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(packRefsCmd).Standalone()

	packRefsCmd.Flags().Bool("all", false, "pack everything")
	packRefsCmd.Flags().String("exclude", "", "references to exclude")
	packRefsCmd.Flags().String("include", "", "references to include")
	packRefsCmd.Flags().Bool("prune", false, "prune loose refs (default)")
	rootCmd.AddCommand(packRefsCmd)
}
