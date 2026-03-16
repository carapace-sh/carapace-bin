package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fzf",
	Short: "a command-line fuzzy finder",
	Long:  "https://github.com/junegunn/fzf",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("accept-nth", "", "Define which fields to print on accept")
	rootCmd.Flags().String("algo", "", "Fuzzy matching algorithm")
	rootCmd.Flags().Bool("ansi", false, "Enable processing of ANSI color codes")
	rootCmd.Flags().Bool("bash", false, "Print script to set up Bash shell integration")
	rootCmd.Flags().String("bind", "", "Custom key bindings")
	rootCmd.Flags().String("border", "", "Draw border around the finder")
	rootCmd.Flags().String("border-label", "", "Label to print on the border")
	rootCmd.Flags().String("border-label-pos", "", "Position of the border label")
	rootCmd.Flags().String("color", "", "Base scheme")
	rootCmd.Flags().Bool("cycle", false, "Enable cyclic scroll")
	rootCmd.Flags().StringP("delimiter", "d", "", "Field delimiter regex")
	rootCmd.Flags().Bool("disabled", false, "Do not perform search")
	rootCmd.Flags().String("ellipsis", "", "Ellipsis to show when line is truncated")
	rootCmd.Flags().BoolP("exact", "e", false, "Enable Exact-match")
	rootCmd.Flags().BoolP("exit-0", "0", false, "Exit immediately when there's no match")
	rootCmd.Flags().String("expect", "", "Comma-separated list of keys to complete fzf")
	rootCmd.Flags().BoolP("extended", "x", false, "Extended-search mode")
	rootCmd.Flags().Bool("filepath-word", false, "Make word-wise movements respect path separators")
	rootCmd.Flags().StringP("filter", "f", "", "Filter mode")
	rootCmd.Flags().Bool("fish", false, "Print script to set up Fish shell integration")
	rootCmd.Flags().String("footer", "", "String to print as footer")
	rootCmd.Flags().String("footer-border", "", "Draw border around the footer section")
	rootCmd.Flags().String("footer-label", "", "Label to print on the footer border")
	rootCmd.Flags().String("footer-label-pos", "", "Position of the footer label")
	rootCmd.Flags().String("freeze-left", "", "Number of fields to freeze on the left")
	rootCmd.Flags().String("freeze-right", "", "Number of fields to freeze on the right")
	rootCmd.Flags().String("gap", "", "Render empty lines between each item")
	rootCmd.Flags().String("gap-line", "", "Draw horizontal line on each gap using the string")
	rootCmd.Flags().String("ghost", "", "Ghost text to display when the input is empty")
	rootCmd.Flags().String("gutter", "", "Character used for the gutter column")
	rootCmd.Flags().String("gutter-raw", "", "Character used for the gutter column in raw mode")
	rootCmd.Flags().String("header", "", "String to print as header")
	rootCmd.Flags().String("header-border", "", "Draw border around the header section")
	rootCmd.Flags().Bool("header-first", false, "Print header before the prompt line")
	rootCmd.Flags().String("header-label-pos", "", "Position of the header label")
	rootCmd.Flags().String("header-lines", "", "The first N lines of the input are treated as header")
	rootCmd.Flags().String("header-lines-border", "", "Display header from --header-lines with a separate border")
	rootCmd.Flags().String("height", "", "Display fzf window below the cursor with the given height instead of using fullscreen")
	rootCmd.Flags().Bool("help", false, "Show this message")
	rootCmd.Flags().Bool("highlight-line", false, "Highlight the whole current line")
	rootCmd.Flags().String("history", "", "File to store fzf search history (*not* shell command history)")
	rootCmd.Flags().String("history-size", "", "Maximum number of history entries")
	rootCmd.Flags().String("hscroll-off", "", "Number of screen columns to keep to the right")
	rootCmd.Flags().BoolP("ignore-case", "i", false, "Case-insensitive match")
	rootCmd.Flags().String("info", "", "Finder info style")
	rootCmd.Flags().String("info-command", "", "Command to generate info line")
	rootCmd.Flags().String("input-border", "", "Draw border around the input section")
	rootCmd.Flags().String("input-label", "", "Label to print on the input border")
	rootCmd.Flags().String("input-label-pos", "", "Position of the input label")
	rootCmd.Flags().String("jump-labels", "", "Label characters for jump and jump-accept")
	rootCmd.Flags().Bool("keep-right", false, "Keep the right end of the line visible on overflow")
	rootCmd.Flags().String("layout", "", "Choose layout")
	rootCmd.Flags().String("list-border", "", "Draw border around the list section")
	rootCmd.Flags().String("list-label", "", "Label to print on the list border")
	rootCmd.Flags().String("list-label-pos", "", "Position of the list label")
	rootCmd.Flags().String("listen", "", "Start HTTP server to receive actions via TCP/Unix domain socket")
	rootCmd.Flags().String("listen-unsafe", "", "allow remote process execution")
	rootCmd.Flags().Bool("literal", false, "Do not normalize latin script letters before matching")
	rootCmd.Flags().Bool("man", false, "Show man page")
	rootCmd.Flags().String("margin", "", "Screen margin")
	rootCmd.Flags().String("marker", "", "Multi-select marker")
	rootCmd.Flags().String("marker-multi-line", "", "Multi-select marker for multi-line entries")
	rootCmd.Flags().String("min-height", "", "Minimum height when --height is given as a percentage")
	rootCmd.Flags().StringP("multi", "m", "", "Enable multi-select with tab/shift-tab")
	rootCmd.Flags().Bool("no-bold", false, "Do not use bold text")
	rootCmd.Flags().Bool("no-clear", false, "Do not clear finder interface on exit")
	rootCmd.Flags().Bool("no-color", false, "Disable colors")
	rootCmd.Flags().Bool("no-extended", false, "Disable extended-search mode")
	rootCmd.Flags().Bool("no-hscroll", false, "Disable horizontal scroll")
	rootCmd.Flags().Bool("no-ignore-case", false, "Case-sensitive match")
	rootCmd.Flags().Bool("no-info", false, "A synonym for --info=hidden")
	rootCmd.Flags().Bool("no-input", false, "Disable and hide the input section")
	rootCmd.Flags().Bool("no-mouse", false, "Disable mouse")
	rootCmd.Flags().Bool("no-multi", false, "Disable multi-select")
	rootCmd.Flags().Bool("no-multi-line", false, "Disable multi-line display of items when using --read0")
	rootCmd.Flags().Bool("no-scrollbar", false, "Hide scrollbar")
	rootCmd.Flags().Bool("no-separator", false, "Hide info line separator")
	rootCmd.Flags().Bool("no-sort", false, "Do not sort the result")
	rootCmd.Flags().Bool("no-tty-default", false, "Make fzf search for the current TTY device via standard error instead of defaulting to /dev/tty")
	rootCmd.Flags().Bool("no-unicode", false, "Use ASCII characters instead of Unicode drawing characters to draw borders, the spinner and the horizontal separator")
	rootCmd.Flags().StringP("nth", "n", "", "Comma-separated list of field index expressions for limiting search scope")
	rootCmd.Flags().String("padding", "", "Padding inside border")
	rootCmd.Flags().Bool("phony", false, "Do not perform search")
	rootCmd.Flags().String("pointer", "", "Pointer to the current line")
	rootCmd.Flags().String("preview", "", "Command to preview highlighted line")
	rootCmd.Flags().String("preview-border", "", "Short for --preview-window=border-STYLE")
	rootCmd.Flags().String("preview-label", "", "Label to print on the horizontal border line of the preview window")
	rootCmd.Flags().String("preview-label-pos", "", "Position of the border label on the border line of the preview window")
	rootCmd.Flags().String("preview-window", "", "Preview window layout")
	rootCmd.Flags().String("preview-wrap-sign", "", "Indicator for wrapped lines in the preview window")
	rootCmd.Flags().Bool("print-query", false, "Print query as the first line")
	rootCmd.Flags().Bool("print0", false, "Print output delimited by ASCII NUL characters")
	rootCmd.Flags().String("prompt", "", "Input prompt")
	rootCmd.Flags().StringP("query", "q", "", "Start the finder with the given query")
	rootCmd.Flags().Bool("raw", false, "Enable raw mode (show non-matching items)")
	rootCmd.Flags().Bool("read0", false, "Read input delimited by ASCII NUL characters")
	rootCmd.Flags().String("scheme", "", "Scoring scheme")
	rootCmd.Flags().String("scroll-off", "", "Number of screen lines to keep above or below when scrolling to the top or to the bottom")
	rootCmd.Flags().String("scrollbar", "", "Scrollbar character(s) (each for list section and preview window)")
	rootCmd.Flags().BoolP("select-1", "1", false, "Automatically select the only match")
	rootCmd.Flags().String("separator", "", "Draw horizontal separator on info line using the string")
	rootCmd.Flags().Bool("smart-case", false, "Smart-case match")
	rootCmd.Flags().String("style", "", "Apply a style preset")
	rootCmd.Flags().Bool("sync", false, "Synchronous search for multi-staged filtering")
	rootCmd.Flags().String("tabstop", "", "Number of spaces for a tab character")
	rootCmd.Flags().Bool("tac", false, "Reverse the order of the input")
	rootCmd.Flags().String("tail", "", "Maximum number of items to keep in memory")
	rootCmd.Flags().String("tiebreak", "", "Comma-separated list of sort criteria to apply")
	rootCmd.Flags().String("tmux", "", "Start fzf in a tmux popup")
	rootCmd.Flags().Bool("track", false, "Track the current selection when the result is updated")
	rootCmd.Flags().Bool("version", false, "Display version information and exit")
	rootCmd.Flags().String("walker", "", "Determines the behavior of the built-in directory walker")
	rootCmd.Flags().String("walker-root", "", "List of directories to walk")
	rootCmd.Flags().String("walker-skip", "", "Comma-separated list of directory names to skip")
	rootCmd.Flags().String("with-nth", "", "Transform the presentation of each line using field index expressions")
	rootCmd.Flags().String("with-shell", "", "Shell command and flags to start child processes with")
	rootCmd.Flags().String("wrap", "", "Enable line wrap")
	rootCmd.Flags().String("wrap-sign", "", "Indicator for wrapped lines")
	rootCmd.Flags().Bool("zsh", false, "Print script to set up Zsh shell integration")

	rootCmd.Flag("border").NoOptDefVal = ""
	rootCmd.Flag("footer-border").NoOptDefVal = ""
	rootCmd.Flag("gap").NoOptDefVal = ""
	rootCmd.Flag("gap-line").NoOptDefVal = ""
	rootCmd.Flag("header-border").NoOptDefVal = ""
	rootCmd.Flag("header-lines-border").NoOptDefVal = ""
	rootCmd.Flag("input-border").NoOptDefVal = ""
	rootCmd.Flag("list-border").NoOptDefVal = ""
	rootCmd.Flag("listen").NoOptDefVal = ""
	rootCmd.Flag("multi").NoOptDefVal = ""
	rootCmd.Flag("preview-border").NoOptDefVal = ""
	rootCmd.Flag("scrollbar").NoOptDefVal = ""
	rootCmd.Flag("tmux").NoOptDefVal = ""
	rootCmd.Flag("wrap").NoOptDefVal = ""

	borderStyles := []string{"rounded", "sharp", "bold", "block", "thinblock", "double", "horizontal", "vertical", "top", "bottom", "left", "right", "line", "none"}

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"algo":                carapace.ActionValues("v1", "v2"),
		"border":              carapace.ActionValues(borderStyles...),
		"color":               carapace.ActionValues("dark", "light", "16", "bw"),
		"footer-border":       carapace.ActionValues(borderStyles...),
		"header-border":       carapace.ActionValues(borderStyles...),
		"header-lines-border": carapace.ActionValues(borderStyles...),
		"history":             carapace.ActionFiles(),
		"info":                carapace.ActionValues("default", "right", "hidden", "inline", "inline-right"),
		"input-border":        carapace.ActionValues(borderStyles...),
		"layout":              carapace.ActionValues("default", "reverse", "reverse-list"),
		"list-border":         carapace.ActionValues(borderStyles...),
		"preview-border":      carapace.ActionValues(borderStyles...),
		"scheme":              carapace.ActionValues("default", "path", "history"),
		"tiebreak":            carapace.ActionValues("length", "chunk", "pathname", "begin", "end", "index").UniqueList(","),
		"walker":              carapace.ActionValues("file", "dir", "follow", "hidden").UniqueList(","),
		"walker-skip":         carapace.ActionDirectories().UniqueList(","),
		"wrap":                carapace.ActionValues("char", "word"),
	})
}
