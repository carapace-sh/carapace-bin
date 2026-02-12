package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/time"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Interactive history search",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	searchCmd.Flags().String("after", "", "Only include results after this date")
	searchCmd.Flags().StringP("before", "b", "", "Only include results added before this date")
	searchCmd.Flags().Bool("cmd-only", false, "Show only the text of the command")
	searchCmd.Flags().StringP("cwd", "c", "", "Filter search result by directory")
	searchCmd.Flags().Bool("delete", false, "Delete anything matching this query. Will not print out the match")
	searchCmd.Flags().Bool("delete-it-all", false, "Delete EVERYTHING!")
	searchCmd.Flags().String("exclude-cwd", "", "Exclude directory from results")
	searchCmd.Flags().String("exclude-exit", "", "Exclude results with this exit code")
	searchCmd.Flags().StringP("exit", "e", "", "Filter search result by exit code")
	searchCmd.Flags().String("filter-mode", "", "Allow overriding filter mode over config")
	searchCmd.Flags().StringP("format", "f", "", "Available variables: {command}, {directory}, {duration}, {user}, {host}, {time}, {exit} and {relativetime}. Example: --format \"{time} - [{duration}] - {directory}$\\t{command}\"")
	searchCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	searchCmd.Flags().Bool("human", false, "Use human-readable formatting for time")
	searchCmd.Flags().Bool("include-duplicates", false, "Include duplicate commands in the output (non-interactive only)")
	searchCmd.Flags().String("inline-height", "", "Set the maximum number of lines Atuin's interface should take up")
	searchCmd.Flags().BoolP("interactive", "i", false, "Open interactive search UI")
	searchCmd.Flags().String("keymap-mode", "", "Notify the keymap at the shell's side")
	searchCmd.Flags().String("limit", "", "How many entries to return at most")
	searchCmd.Flags().String("offset", "", "Offset from the start of the results")
	searchCmd.Flags().Bool("print0", false, "Terminate the output with a null, for better multiline handling")
	searchCmd.Flags().BoolP("reverse", "r", false, "Reverse the order of results, oldest first")
	searchCmd.Flags().String("search-mode", "", "Allow overriding search mode over config")
	searchCmd.Flags().String("timezone", "", "Display the command time in another timezone other than the configured default")
	searchCmd.Flags().String("tz", "", "Display the command time in another timezone other than the configured default")
	rootCmd.AddCommand(searchCmd)

	carapace.Gen(searchCmd).FlagCompletion(carapace.ActionMap{
		"after":       time.ActionDate(),
		"before":      time.ActionDate(),
		"cwd":         carapace.ActionDirectories(),
		"exclude-cwd": carapace.ActionDirectories(),
		"filter-mode": carapace.ActionValues("global", "host", "session", "directory", "workspace", "session-preloa"),
		"keymap-mode": carapace.ActionValues("emacs", "vim-normal", "vim-insert", "auto"),
		"search-mode": carapace.ActionValues("prefix", "full-text", "fuzzy", "skim"),
		"timezone":    carapace.ActionValues(), // TODO
		"tz":          carapace.ActionValues(), // TODO
	})
}
