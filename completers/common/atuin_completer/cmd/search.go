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

	searchCmd.Flags().String("after", "", "Only include results after this date; see `--before` for how it is interpreted")
	searchCmd.Flags().StringSlice("author", nil, "Filter by author. Supports $all-user (non-agents), $all-agent, or literal names")
	searchCmd.Flags().StringP("before", "b", "", "Only include results added before this date")
	searchCmd.Flags().Bool("cmd-only", false, "Show only the text of the command")
	searchCmd.Flags().StringP("cwd", "c", "", "Filter search result by directory")
	searchCmd.Flags().Bool("delete", false, "Delete anything matching this query. Will not print out the match")
	searchCmd.Flags().Bool("delete-it-all", false, "Delete EVERYTHING!")
	searchCmd.Flags().String("exclude-cwd", "", "Exclude directory from results")
	searchCmd.Flags().StringSlice("exclude-exit", nil, "Exclude results with this exit code; repeat to exclude multiple codes")
	searchCmd.Flags().StringSliceP("exit", "e", nil, "Filter by exit code; repeat to include any of the given codes")
	searchCmd.Flags().String("filter-mode", "", "Allow overriding filter mode over config")
	searchCmd.Flags().StringP("format", "f", "", "Available variables: {command}, {directory}, {duration}, {user}, {host}, {time}, {exit} and {relativetime}")
	searchCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	searchCmd.Flags().Bool("human", false, "Use human-readable formatting for time")
	searchCmd.Flags().Bool("include-duplicates", false, "Include duplicate commands in the output (non-interactive only)")
	searchCmd.Flags().String("inline-height", "", "Set the maximum number of lines Atuin's interface should take up")
	searchCmd.Flags().BoolP("interactive", "i", false, "Open interactive search UI")
	searchCmd.Flags().String("keymap-mode", "auto", "Notify the keymap at the shell's side")
	searchCmd.Flags().String("limit", "", "How many entries to return at most")
	searchCmd.Flags().String("offset", "", "Offset from the start of the results")
	searchCmd.Flags().Bool("print0", false, "Terminate the output with a null, for better multiline support")
	searchCmd.Flags().String("result-file", "", "File name to write the result to (hidden from help as this is meant to be used from a script)")
	searchCmd.Flags().BoolP("reverse", "r", false, "Reverse the order of results, oldest first")
	searchCmd.Flags().String("search-mode", "", "Allow overriding search mode over config")
	searchCmd.Flags().StringSlice("shell", nil, "Filter by the shell that was used to run the command")
	searchCmd.Flags().Bool("shell-up-key-binding", false, "Marker argument used to inform atuin that it was invoked from a shell up-key binding (hidden from help to avoid confusion)")
	searchCmd.Flags().String("timezone", "", "Timezone to display command times in and to interpret `--before`/`--after` in, instead")
	searchCmd.Flags().String("tz", "", "Timezone to display command times in and to interpret `--before`/`--after` in, instead")
	searchCmd.Flag("result-file").Hidden = true
	searchCmd.Flag("shell-up-key-binding").Hidden = true
	rootCmd.AddCommand(searchCmd)

	carapace.Gen(searchCmd).FlagCompletion(carapace.ActionMap{
		"after": time.ActionDate(),
		"author": carapace.ActionValuesDescribed(
			"$all-agent", "agents",
			"$all-user", "non-agents",
			// TODO existing users
		),
		"before":      time.ActionDate(),
		"cwd":         carapace.ActionDirectories(),
		"exclude-cwd": carapace.ActionDirectories(),
		"filter-mode": carapace.ActionValues("directory", "global", "host", "session", "session-preload", "workspace"),
		"keymap-mode": carapace.ActionValues("auto", "emacs", "vim-insert", "vim-normal"),
		"result-file": carapace.ActionFiles(),
		"search-mode": carapace.ActionValues("daemon-fuzzy", "fulltext", "fuzzy", "prefix", "skim"),
		"shell":       carapace.ActionValues("bash", "fish", "nu", "powershell", "xonsh", "zsh"), // TODO other shells
		"timezone":    carapace.ActionValues(),                                                   // TODO
		"tz":          carapace.ActionValues(),                                                   // TODO
	})
}
