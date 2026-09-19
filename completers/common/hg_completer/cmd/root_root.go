package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var root_rootCmd = &cobra.Command{
	Use:     "root",
	Short:   "print the root (top) of the current working directory",
	GroupID: groups[group_working_directory_management].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(root_rootCmd).Standalone()

	root_rootCmd.Flags().Bool("share-source", false, "print the share source root instead")
	root_rootCmd.Flags().StringP("template", "T", "", "display with template")
	rootCmd.AddCommand(root_rootCmd)
}
