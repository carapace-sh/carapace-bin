package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"

	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hl"
)

var rootCmd = &cobra.Command{
	Use:   "hl",
	Short: "JSON and logfmt log converter to human readable representation",
	Long:  "https://github.com/pamburus/hl",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("P", "P", false, "Handful alias for --paging=never, overrides --paging option")
	rootCmd.Flags().Bool("allow-prefix", false, "Allow non-JSON prefixes before JSON log entries")
	rootCmd.Flags().String("ascii", "", "Whether to restrict punctuation to ASCII characters only")
	rootCmd.Flags().String("buffer-size", "", "Buffer size")
	rootCmd.Flags().BoolS("c", "c", false, "Handful alias for --color=always, overrides --color option")
	rootCmd.Flags().String("color", "", "Whether to use ANSI colors and styles")
	rootCmd.Flags().StringP("concurrency", "C", "", "Number of processing threads")
	rootCmd.Flags().StringSlice("config", nil, "Configuration file path")
	rootCmd.Flags().String("delimiter", "", "Log entry delimiter")
	rootCmd.Flags().Bool("dump-index", false, "Print debug index metadata (in --sort mode) and exit")
	rootCmd.Flags().StringP("expansion", "x", "", "Whether to expand fields and messages")
	rootCmd.Flags().StringSliceP("filter", "f", nil, "Filter entries by matching field values [k=v, k~=v, k~~=v, 'k!=v', 'k?!=v', etc]")
	rootCmd.Flags().String("flatten", "", "Whether to flatten objects")
	rootCmd.Flags().BoolP("follow", "F", false, "Follow input streams and sort entries chronologically within time frame set by --sync-interval-ms option")
	rootCmd.Flags().String("help", "", "Print help")
	rootCmd.Flags().StringSliceP("hide", "h", nil, "Hide or reveal fields with the specified keys, prefix with ! to reveal, provide '!*' to reveal all")
	rootCmd.Flags().BoolP("hide-empty-fields", "e", false, "Hide empty fields, applies for null, string, object and array fields only")
	rootCmd.Flags().String("input-format", "", "Input format")
	rootCmd.Flags().String("input-info", "", "Input number and filename layouts")
	rootCmd.Flags().String("interrupt-ignore-count", "", "Number of interrupts to ignore, i.e. Ctrl-C (SIGINT)")
	rootCmd.Flags().StringP("level", "l", "", "Display entries with level >= <LEVEL>")
	rootCmd.Flags().String("list-themes", "", "Print available themes optionally filtered by tags")
	rootCmd.Flags().BoolP("local", "L", false, "Use local time zone, overrides --time-zone option")
	rootCmd.Flags().Bool("man-page", false, "Print man page and exit")
	rootCmd.Flags().String("max-message-size", "", "Maximum log entry size")
	rootCmd.Flags().Bool("no-local", false, "Disable local time zone, overrides --local option")
	rootCmd.Flags().Bool("no-raw", false, "Disable raw source entries output, overrides --raw option")
	rootCmd.Flags().StringP("output", "o", "", "Output file")
	rootCmd.Flags().String("output-delimiter", "", "Output entry delimiter")
	rootCmd.Flags().String("paging", "", "Control pager usage (HL_PAGER or PAGER)")
	rootCmd.Flags().StringSliceP("query", "q", nil, "Filter entries using a query expression ['status>=400 or duration>=15', etc]")
	rootCmd.Flags().BoolP("raw", "r", false, "Output raw source entries instead of formatted entries")
	rootCmd.Flags().Bool("raw-fields", false, "Output field values as is, without unescaping or prettifying")
	rootCmd.Flags().String("shell-completions", "", "Print shell auto-completion script and exit")
	rootCmd.Flags().BoolP("show-empty-fields", "E", false, "Show empty fields, overrides --hide-empty-fields option")
	rootCmd.Flags().String("since", "", "Display entries with timestamp >= <TIME>")
	rootCmd.Flags().BoolP("sort", "s", false, "Sort entries chronologically")
	rootCmd.Flags().String("sync-interval-ms", "", "Synchronization interval for live streaming mode enabled by --follow option")
	rootCmd.Flags().String("tail", "", "Number of last entries to preload from each file in --follow mode")
	rootCmd.Flags().String("theme", "", "Color theme")
	rootCmd.Flags().StringP("time-format", "t", "", "Time format, see https://man7.org/linux/man-pages/man1/date.1.html")
	rootCmd.Flags().StringP("time-zone", "Z", "", "Time zone name, see column \"TZ identifier\" at https://en.wikipedia.org/wiki/List_of_tz_database_time_zones")
	rootCmd.Flags().String("unix-timestamp-unit", "", "Unix timestamp unit")
	rootCmd.Flags().String("until", "", "Display entries with timestamp <= <TIME>")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")
	rootCmd.Flag("help").NoOptDefVal = " "
	rootCmd.Flag("list-themes").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"ascii":               carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"color":               carapace.ActionValues("auto", "always", "never").StyleF(style.ForKeyword),
		"concurrency":         carapace.ActionValues(),
		"config":              carapace.ActionFiles(),
		"delimiter":           carapace.ActionValues("auto", "cr", "lf", "crlf", "newline", "nul"),
		"expansion":           carapace.ActionValues("never", "inline", "auto", "always"),
		"flatten":             carapace.ActionValues("never", "always").StyleF(style.ForKeyword),
		"help":                carapace.ActionValues("short", "long"),
		"input-format":        carapace.ActionValues("auto", "json", "logfmt"),
		"input-info":          carapace.ActionValues("auto", "none", "minimal", "compact", "full"),
		"list-themes":         carapace.ActionValues("dark", "light", "16color", "256color", "truecolor", "overlay", "base"),
		"output":              carapace.ActionFiles(),
		"output-delimiter":    carapace.ActionValues("newline", "nul"),
		"paging":              carapace.ActionValues("auto", "always", "never").StyleF(style.ForKeyword),
		"shell-completions":   carapace.ActionValues("bash", "elvish", "fish", "powershell", "zsh"),
		"theme":               hl.ActionThemes(),
		"unix-timestamp-unit": carapace.ActionValues("auto", "s", "ms", "us", "ns"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
