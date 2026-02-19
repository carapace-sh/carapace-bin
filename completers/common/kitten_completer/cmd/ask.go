package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var askCmd = &cobra.Command{
	Use:   "ask",
	Short: "Ask the user for input",
}

func init() {
	rootCmd.AddCommand(askCmd)
	carapace.Gen(askCmd).Standalone()

	askCmd.Flags().StringP("type", "t", "", "Type of input. Defaults to asking for a line of text.")
	askCmd.Flags().StringP("message", "m", "", "The message to display to the user. If not specified a default message is shown.")
	askCmd.Flags().StringP("name", "n", "", "The name for this question. Used to store history of previous answers which can be used for completions and via the browse history readline bindings.")
	askCmd.Flags().String("title", "", "The title for the window in which the question is displayed. Only implemented for yesno and choices types.")
	askCmd.Flags().String("window-title", "", "The title for the window in which the question is displayed. Only implemented for yesno and choices types.")
	askCmd.Flags().StringP("choice", "c", "", "A choice for the choices type. Can be specified multiple times. Every choice has the syntax: ``letter[;color]:text``, where text is the choice text and letter is the selection key. letter is a single letter belonging to text. This letter is highlighted within the choice text. There can be an optional color specification after the letter to indicate what color it should be. For example: y:Yes and n;red:No")
	askCmd.Flags().StringP("default", "d", "", "A default choice or text. If unspecified, it is y for the type yesno, the first choice for choices and empty for others types. The default choice is selected when the user presses the Enter key.")
	askCmd.Flags().StringP("prompt", "p", "", "The prompt to use when inputting a line of text or a password.")
	askCmd.Flags().String("unhide-key", "", "The key to be pressed to unhide hidden text")
	askCmd.Flags().String("hidden-text-placeholder", "", "The text in the message to be replaced by hidden text. The hidden text is read via STDIN.")
	askCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(askCmd).FlagCompletion(carapace.ActionMap{
		"type": carapace.ActionValues("line", "choices", "file", "password", "yesno"),
	})
}
