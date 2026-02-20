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

	lsCmd.Flags().Bool("all-env-vars", false, "Show all environment variables in output, not just differing ones")
	lsCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	lsCmd.Flags().StringP("match", "m", "", "The window to match")
	lsCmd.Flags().StringP("match-tab", "t", "", "The tab to match")
	lsCmd.Flags().String("output-format", "json", "Output in JSON or kitty session format")
	lsCmd.Flags().Bool("self", false, "Only list the window this command is run in")

	carapace.Gen(lsCmd).FlagCompletion(carapace.ActionMap{
		"output-format": carapace.ActionValues("json", "session"),
	})
}
