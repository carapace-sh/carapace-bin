package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List tabs/windows",
}

func init() {
	lsCmd.AddCommand(atCmd)
	carapace.Gen(lsCmd).Standalone()

	lsCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(lsCmd).FlagCompletion(carapace.ActionMap{})
}