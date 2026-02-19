package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clipboardCmd = &cobra.Command{
	Use:   "clipboard",
	Short: "Copy/paste with the system clipboard, even over SSH",
}

func init() {
	rootCmd.AddCommand(clipboardCmd)
	carapace.Gen(clipboardCmd).Standalone()

	clipboardCmd.Flags().BoolP("get-clipboard", "g", false, "Output the current contents of the clipboard to STDOUT. Note that by default kitty will prompt for permission to access the clipboard. Can be controlled by clipboard_control.")
	clipboardCmd.Flags().BoolP("use-primary", "p", false, "Use the primary selection rather than the clipboard on systems that support it, such as Linux.")
	clipboardCmd.Flags().StringArrayP("mime", "m", nil, "The mimetype of the specified file. Useful when the auto-detected mimetype is likely to be incorrect or the filename has no extension and therefore no mimetype can be detected. If more than one file is specified, this option should be specified multiple times, once for each specified file. When copying data from the clipboard, you can use wildcards to match MIME types. For example: --mime 'text/*' will match any textual MIME type available on the clipboard, usually the first matching MIME type is copied. The special MIME type . will return the list of available MIME types currently on the system clipboard.")
	clipboardCmd.Flags().StringArrayP("alias", "a", nil, "Specify aliases for MIME types. Aliased MIME types are considered equivalent. When copying to clipboard both the original and alias are made available on the clipboard. When copying from clipboard if the original is not found, the alias is used, as a fallback. Can be specified multiple times to create multiple aliases. For example: --alias text/plain=text/x-rst makes text/plain an alias of text/rst. Aliases are not used in filter mode. An alias for text/plain is automatically created if text/plain is not present in the input data, but some other text/* MIME is present.")
	clipboardCmd.Flags().Bool("wait-for-completion", false, "Wait till the copy to clipboard is complete before exiting. Useful if running the kitten in a dedicated, ephemeral window. Only needed in filter mode.")
	clipboardCmd.Flags().String("password", "", "A password to use when accessing the clipboard. If the user chooses to accept the password future invocations of the kitten will not have a permission prompt in this tty session. Does not work in filter mode. Must be of the form: text:actual-password or fd:integer (a file descriptor number to read the password from) or file:path-to-file (a file from which to read the password). Note that you must also specify a human friendly name using the --human-name flag.")
	clipboardCmd.Flags().String("human-name", "", "A human friendly name to show the user when asking for permission to access the clipboard.")
	clipboardCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(clipboardCmd).FlagCompletion(carapace.ActionMap{
		"password": carapace.ActionValues("text:", "fd:", "file:<file>").MultiPartsP(":", "<.*>", func(placeholder string, matches map[string]string) carapace.Action {
			switch placeholder {
			case "<file>":
				return carapace.ActionFiles()
			default:
				return carapace.ActionValues()
			}
		}),
	})
	carapace.Gen(clipboardCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
