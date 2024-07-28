package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "_complete [-s|--shell SHELL] [-i|--input INPUT] [-c|--current CURRENT] [-a|--api-version API-VERSION] [-S|--symfony SYMFONY]",
	Short: "Internal command to provide shell completion suggestions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completeCmd).Standalone()

	completeCmd.Flags().String("api-version", "", "The API version of the completion script")
	completeCmd.Flags().String("current", "", "The index of the \"input\" array that the cursor is in (e.g. COMP_CWORD)")
	completeCmd.Flags().String("input", "", "An array of input tokens (e.g. COMP_WORDS or argv)")
	completeCmd.Flags().String("shell", "", "The shell type (\"bash\", \"fish\", \"zsh\")")
	completeCmd.Flags().String("symfony", "", "deprecated")
	rootCmd.AddCommand(completeCmd)

	carapace.Gen(completeCmd).FlagCompletion(carapace.ActionMap{
		"shell": carapace.ActionValues("bash", "fish", "zsh"),
	})
}
