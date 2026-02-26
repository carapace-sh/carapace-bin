package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var at_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List tabs/windows",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(at_lsCmd).Standalone()

	at_lsCmd.Flags().Bool("all-env-vars", false, "Show all environment variables in output, not just differing ones")
	at_lsCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	at_lsCmd.Flags().StringP("match", "m", "", "The window to match")
	at_lsCmd.Flags().StringP("match-tab", "t", "", "The tab to match")
	at_lsCmd.Flags().String("output-format", "json", "Output in JSON or kitty session format")
	at_lsCmd.Flags().Bool("self", false, "Only list the window this command is run in")
	atCmd.AddCommand(at_lsCmd)

	carapace.Gen(at_lsCmd).FlagCompletion(carapace.ActionMap{
		"output-format": carapace.ActionValues("json", "session"),
	})
}
