package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var modifyCmd = &cobra.Command{
	Use:     "modify",
	Short:   "Interactively restructure a stack",
	GroupID: "stack",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(modifyCmd).Standalone()

	modifyCmd.Flags().Bool("abort", false, "Abort the modify session and restore the stack to its pre-modify state")
	modifyCmd.Flags().Bool("continue", false, "Continue after resolving conflicts")
	rootCmd.AddCommand(modifyCmd)
}
