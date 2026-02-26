package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_setUserVarsCmd = &cobra.Command{
	Use:   "set-user-vars",
	Short: "Set user variables on a window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_setUserVarsCmd).Standalone()

	at_setUserVarsCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_setUserVarsCmd.Flags().StringP("match", "m", "", "The window to match")
	atCmd.AddCommand(at_setUserVarsCmd)

}
