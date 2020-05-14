package cmd

import (
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use: "clean",
	Short: "Remove untracked files from the working tree",
    Run: func(cmd *cobra.Command, args []string) {
    },
}

func init() {
	cleanCmd.Flags().BoolP("d", "d", false, "remove whole directories")
	cleanCmd.Flags().BoolP("exclude", "e", false, "<pattern>    add <pattern> to ignore rules")
	cleanCmd.Flags().BoolP("force", "f", false, "force")
	cleanCmd.Flags().BoolP("interactive", "i", false, "interactive cleaning")
	cleanCmd.Flags().BoolP("dry-run", "n", false, "dry run")
	cleanCmd.Flags().BoolP("quiet", "q", false, "do not print names of files removed")
	cleanCmd.Flags().BoolP("x", "x", false, "remove ignored files, too")
	cleanCmd.Flags().BoolP("X", "X", false, "remove only ignored files")
    rootCmd.AddCommand(cleanCmd)
}
