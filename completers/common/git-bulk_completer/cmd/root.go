package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-bulk",
	Short: "Run git commands on multiple repositories across registered workspaces",
	Long:  "https://github.com/tj/git-extras/blob/master/Commands.md#git-bulk",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("addcurrent", "", "Register the current directory as a workspace")
	rootCmd.Flags().StringArray("addworkspace", nil, "Register a workspace")
	rootCmd.Flags().BoolP("all", "a", false, "Run command on all registered workspaces")
	rootCmd.Flags().BoolP("g", "g", false, "Guarded mode, ask for confirmation")
	rootCmd.Flags().Bool("help", false, "show help")
	rootCmd.Flags().Bool("listall", false, "List all registered workspaces")
	rootCmd.Flags().Bool("no-follow-hidden", false, "Don't follow hidden directories")
	rootCmd.Flags().Bool("no-follow-symlinks", false, "Don't follow symlinks")
	rootCmd.Flags().Bool("purge", false, "Remove all registered workspaces")
	rootCmd.Flags().BoolP("quiet", "q", false, "Quiet mode")
	rootCmd.Flags().String("removeworkspace", "", "Remove a registered workspace")
	rootCmd.Flags().StringP("workspace", "w", "", "Run command on a specific workspace")
}
