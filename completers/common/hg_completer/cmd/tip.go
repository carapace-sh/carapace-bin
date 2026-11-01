package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tipCmd = &cobra.Command{
	Use:     "tip",
	Short:   "show the tip revision (DEPRECATED)",
	GroupID: groups[group_change_navigation].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tipCmd).Standalone()

	tipCmd.Flags().BoolP("git", "g", false, "use git extended diff format")
	tipCmd.Flags().BoolP("patch", "p", false, "show patch")
	tipCmd.Flags().String("style", "", "display using template map file (DEPRECATED)")
	tipCmd.Flags().StringP("template", "T", "", "display with template")
	rootCmd.AddCommand(tipCmd)
}
