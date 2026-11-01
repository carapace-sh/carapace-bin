package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var abortCmd = &cobra.Command{
	Use:     "abort",
	Short:   "abort an unfinished operation (EXPERIMENTAL)",
	GroupID: groups[group_change_manipulation].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(abortCmd).Standalone()

	abortCmd.Flags().BoolP("dry-run", "n", false, "do not perform actions, just print output")
	rootCmd.AddCommand(abortCmd)
}
