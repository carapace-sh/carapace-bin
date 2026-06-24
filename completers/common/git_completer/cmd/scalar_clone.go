package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scalar_cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone a repository with Scalar optimizations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scalar_cloneCmd).Standalone()

	scalar_cloneCmd.Flags().StringP("branch", "b", "", "clone only the history leading to the tip of the specified branch")
	scalar_cloneCmd.Flags().Bool("full-clone", false, "do not create a partial clone")
	scalar_cloneCmd.Flags().Bool("maintenance", false, "enable automatic maintenance")
	scalar_cloneCmd.Flags().Bool("single-branch", false, "clone only the history leading to the tip of a single branch")
	scalar_cloneCmd.Flags().Bool("src", false, "show source URL")
	scalar_cloneCmd.Flags().Bool("tags", false, "clone tags")
	scalarCmd.AddCommand(scalar_cloneCmd)
}
