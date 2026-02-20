package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hintsCmd = &cobra.Command{
	Use:   "hints",
	Short: "Select text from the screen using the keyboard",
}

func init() {
	rootCmd.AddCommand(hintsCmd)
	carapace.Gen(hintsCmd).Standalone()

	hintsCmd.Flags().String("add-trailing-space", "", "Add trailing space after matched text. Defaults to auto, which adds the space when used together with --multiple.")
	hintsCmd.Flags().String("alphabet", "", "The list of characters to use for hints. The default is to use numbers and lowercase English alphabets. Specify your preference as a string of characters. Note that you need to specify the --hints-offset as zero to use the first character to highlight the first match, otherwise it will start with the second character by default.")
	hintsCmd.Flags().Bool("ascending", false, "Make the hints increase from top to bottom, instead of decreasing from top to bottom.")
	hintsCmd.Flags().String("customize-processing", "", "Name of a python file in the kitty config directory which will be imported to provide custom implementations for pattern finding and performing actions on selected matches. You can also specify absolute paths to load the script from elsewhere.")
	hintsCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	hintsCmd.Flags().String("hints-background-color", "", "The background color for hints. You can use color names or hex values. For the eight basic named terminal colors you can also use the bright- prefix to get the bright variant of the color.")
	hintsCmd.Flags().String("hints-foreground-color", "", "The foreground color for hints. You can use color names or hex values. For the eight basic named terminal colors you can also use the bright- prefix to get the bright variant of the color.")
	hintsCmd.Flags().Int("hints-offset", 0, "The offset (from zero) at which to start hint numbering. Note that only numbers greater than or equal to zero are respected.")
	hintsCmd.Flags().String("hints-text-color", "", "The foreground color for text pointed to by the hints. You can use color names or hex values. For the eight basic named terminal colors you can also use the bright- prefix to get the bright variant of the color. The default is to pick a suitable color automatically.")
	hintsCmd.Flags().String("linenum-action", "", "Where to perform the action on matched errors. self means the current window, window a new kitty window, tab a new tab, os_window a new OS window and background run in the background. remote-control is like background but the program can use kitty remote control without needing to turn on remote control globally.")
	hintsCmd.Flags().Int("minimum-match-length", 0, "The minimum number of characters to consider a match.")
	hintsCmd.Flags().Bool("multiple", false, "Select multiple matches and perform the action on all of them together at the end. In this mode, press Esc to finish selecting.")
	hintsCmd.Flags().String("multiple-joiner", "", "String for joining multiple selections when copying to the clipboard or inserting into the terminal. The special values are: space - a space character, newline - a newline, empty - an empty joiner, json - a JSON serialized list, auto - an automatic choice, based on the type of text being selected. In addition, integers are interpreted as zero-based indices into the list of selections. You can use 0 for the first selection and -1 for the last.")
	hintsCmd.Flags().StringArray("program", nil, "What program to use to open matched text. Defaults to the default open program for the operating system. Various special values are supported: - (paste into terminal), @ (copy to clipboard), * (copy to primary selection), @NAME (copy to specified buffer), default (run default open program), launch (run the launch command). Can be specified multiple times to run multiple programs.")
	hintsCmd.Flags().String("regex", "", "The regular expression to use when option --type is set to regex, in Perl 5 syntax. If you specify a numbered group in the regular expression, only the group will be matched. This allows you to match text ignoring a prefix/suffix, as needed. The default expression matches lines.")
	hintsCmd.Flags().String("type", "", "The type of text to search for. A value of linenum is special, it looks for error messages using the pattern specified with --regex, which must have the named groups: path and line. If not specified, will look for path:line. The --linenum-action option controls where to display the selected error message, other options are ignored.")
	hintsCmd.Flags().String("url-excluded-characters", "", "Characters to exclude when matching URLs. Defaults to the list of characters defined by the url_excluded_characters option in kitty.conf. The syntax for this option is the same as for url_excluded_characters.")
	hintsCmd.Flags().String("url-prefixes", "", "Comma separated list of recognized URL prefixes. Defaults to the list of prefixes defined by the url_prefixes option in kitty.conf.")
	hintsCmd.Flags().String("window-title", "", "The title for the hints window, default title is based on the type of text being hinted.")
	hintsCmd.Flags().String("word-characters", "", "Characters to consider as part of a word. In addition, all characters marked as alphanumeric in the Unicode database will be considered as word characters. Defaults to the select_by_word_characters option from kitty.conf.")

	carapace.Gen(hintsCmd).FlagCompletion(carapace.ActionMap{
		"program":        carapace.ActionValues("-", "@", "*", "default", "launch"),
		"type":           carapace.ActionValues("url", "hash", "hyperlink", "ip", "line", "linenum", "path", "regex", "word"),
		"linenum-action": carapace.ActionValues("self", "background", "os_window", "remote-control", "tab", "window"),
		"multiple-joiner": carapace.ActionValuesDescribed(
			"space", "a space character",
			"newline", "a newline",
			"empty", "an empty joiner",
			"json", "a JSON serialized list",
			"auto", "an automatic choice, based on the type of text being selected",
		),
		"add-trailing-space":   carapace.ActionValues("auto", "always", "never"),
		"customize-processing": carapace.ActionFiles().Chdir("~/.config/kitty"),
	})
}
