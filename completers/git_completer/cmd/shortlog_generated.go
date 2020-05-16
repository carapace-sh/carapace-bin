package cmd

import (
	"github.com/spf13/cobra"
)

var shortlogCmd = &cobra.Command{
	Use:   "shortlog",
	Short: "Summarize 'git log' output",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	shortlogCmd.Flags().BoolP("committer", "c", false, "Group by committer rather than author")
	shortlogCmd.Flags().BoolP("email", "e", false, "Show the email address of each author")
	shortlogCmd.Flags().BoolP("numbered", "n", false, "sort output according to the number of commits per author")
	shortlogCmd.Flags().BoolP("summary", "s", false, "Suppress commit descriptions, only provides commit count")
	shortlogCmd.Flags().StringP("w", "w", "", "Linewrap output")
	rootCmd.AddCommand(shortlogCmd)
}
