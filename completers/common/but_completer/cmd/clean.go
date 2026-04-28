package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove empty branches from the workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanCmd).Standalone()

	cleanCmd.Flags().Bool("dry-run", false, "Preview which branches would be removed without actually deleting them")
	cleanCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	cleanCmd.Flags().Bool("include-upstream", false, "Also remove branches that have upstream-only commits but no local commits or changes")
	cleanCmd.Flags().Bool("pull", false, "Pull latest changes from the remote before cleaning")
	rootCmd.AddCommand(cleanCmd)
}
