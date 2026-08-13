package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var link_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add projects to an existing repository link",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(link_addCmd).Standalone()

	link_addCmd.Flags().Bool("yes", false, "Skip questions")

	linkCmd.AddCommand(link_addCmd)
}
