package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var distTagCmd = &cobra.Command{
	Use:   "dist-tag",
	Short: "Modify package distribution tags",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(distTagCmd).Standalone()
	distTag_addCmd.PersistentFlags().StringP("workspace", "w", "", "Enable running a command in the context of the given workspace")
	distTag_addCmd.PersistentFlags().Bool("workspaces", false, "Enable running a command in the context fo all workspaces")

	rootCmd.AddCommand(distTagCmd)
}
