package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:     "link",
	Short:   "Connect the local project to another one",
	GroupID: "general",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(linkCmd).Standalone()

	linkCmd.Flags().BoolP("all", "A", false, "Link all workspaces belonging to the target project to the current one")
	linkCmd.Flags().BoolP("private", "p", false, "Also link private workspaces belonging to the target project to the current one")
	linkCmd.Flags().BoolP("relative", "r", false, "Link workspaces using relative paths instead of absolute paths")
	rootCmd.AddCommand(linkCmd)

	carapace.Gen(linkCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
