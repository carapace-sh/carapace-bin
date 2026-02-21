package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setUserVarsCmd = &cobra.Command{
	Use:   "set-user-vars",
	Short: "Set user variables on a window",
}

func init() {
	setUserVarsCmd.AddCommand(atCmd)
	carapace.Gen(setUserVarsCmd).Standalone()

	setUserVarsCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	setUserVarsCmd.Flags().StringP("match", "m", "", "The window to match")

	carapace.Gen(setUserVarsCmd).FlagCompletion(carapace.ActionMap{})
}
