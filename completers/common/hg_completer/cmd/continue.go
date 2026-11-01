package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var continueCmd = &cobra.Command{
	Use:     "continue",
	Short:   "resumes an interrupted operation (EXPERIMENTAL)",
	GroupID: groups[group_change_manipulation].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(continueCmd).Standalone()

	continueCmd.Flags().BoolP("dry-run", "n", false, "do not perform actions, just print output")
	rootCmd.AddCommand(continueCmd)
}
