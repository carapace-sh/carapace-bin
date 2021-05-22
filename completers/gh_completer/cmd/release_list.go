package cmd

import (
	"github.com/spf13/cobra"
)

var release_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List releases in a repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	release_listCmd.Flags().IntP("limit", "L", 30, "Maximum number of items to fetch")
	releaseCmd.AddCommand(release_listCmd)
}
