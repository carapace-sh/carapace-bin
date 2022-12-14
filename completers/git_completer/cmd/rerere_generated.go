package cmd

import (
	"github.com/spf13/cobra"
)

var rerereCmd = &cobra.Command{
	Use:     "rerere",
	Short:   "Reuse recorded resolution of conflicted merges",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_interrogator].ID,
}

func init() {
	rerereCmd.Flags().Bool("rerere-autoupdate", false, "register clean resolutions in index")
	rootCmd.AddCommand(rerereCmd)
}
