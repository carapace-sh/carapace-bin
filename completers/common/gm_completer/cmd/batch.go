package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var batchCmd = &cobra.Command{
	Use:   "batch",
	Short: "issue multiple commands in interactive or batch mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(batchCmd).Standalone()

	batchCmd.Flags().StringS("echo", "echo", "", "echo command back to standard out")
	batchCmd.Flags().StringS("escape", "escape", "", "force use Unix or Windows escape format for command line argument parsing")
	batchCmd.Flags().StringS("fail", "fail", "", "when feedback is on, output the designated text if the command returns error")
	batchCmd.Flags().StringS("feedback", "feedback", "", "print text (see -pass and -fail options) feedback after each command to indicate the result")
	batchCmd.Flags().BoolS("help", "help", false, "print program options")
	batchCmd.Flags().StringS("pass", "pass", "", "when feedback is on, output the designated text if the command executed successfully")
	batchCmd.Flags().StringS("prompt", "prompt", "", "use the given text as command prompt")
	batchCmd.Flags().StringS("stop-on-error", "stop-on-error", "", "when turned on, batch execution quits prematurely when any command returns error")
	batchCmd.Flags().StringS("tap-mode", "tap-mode", "", "Enables or disables Test Anything Protocol (TAP) output for test results")
	rootCmd.AddCommand(batchCmd)

	carapace.Gen(batchCmd).FlagCompletion(carapace.ActionMap{
		"echo":          carapace.ActionValues("on", "off"),
		"escape":        carapace.ActionValues("unix", "windows"),
		"feedback":      carapace.ActionValues("on", "off"),
		"stop-on-error": carapace.ActionValues("on", "off"),
		"tap-mode":      carapace.ActionValues("on", "off"),
	})

	carapace.Gen(batchCmd).PositionalCompletion(carapace.ActionFiles())
}
