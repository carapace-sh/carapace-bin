package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-p4",
	Short: "Import from and submit to Perforce repositories",
	Long:  "https://git-scm.com/docs/git-p4",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("git-dir", "", "Set the GIT_DIR environment variable")
	rootCmd.Flags().BoolP("verbose", "v", false, "Provide more progress information")
}
