package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sqlite3",
	Short: "A command-line interface for SQLite",
	Long:  "https://sqlite.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("A", "A", "", "run \".archive ARGS\" and exit")
	rootCmd.Flags().BoolS("append", "append", false, "Append the database to the end of the file")
	rootCmd.Flags().BoolS("ascii", "ascii", false, "Set output mode to 'ascii'")
	rootCmd.Flags().BoolS("bail", "bail", false, "Stop after hitting an error")
	rootCmd.Flags().BoolS("batch", "batch", false, "Force batch I/O")
	rootCmd.Flags().BoolS("box", "box", false, "Set output mode to 'box'")
	rootCmd.Flags().StringS("cmd", "cmd", "", "run \"COMMAND\" before reading stdin")
	rootCmd.Flags().BoolS("column", "column", false, "Set output mode to 'column'")
	rootCmd.Flags().BoolS("csv", "csv", false, "Set output mode to 'csv'")
	rootCmd.Flags().BoolS("deserialize", "deserialize", false, "open the database using sqlite3_deserialize()")
	rootCmd.Flags().BoolS("echo", "echo", false, "print inputs before execution")
	rootCmd.Flags().BoolS("header", "header", false, "turn headers on or off")
	rootCmd.Flags().BoolS("help", "help", false, "show this message")
	rootCmd.Flags().BoolS("html", "html", false, "set output mode to HTML")
	rootCmd.Flags().StringS("init", "init", "", "read/process named file")
	rootCmd.Flags().BoolS("interactive", "interactive", false, "force interactive I/O")
	rootCmd.Flags().BoolS("json", "json", false, "set output mode to 'json'")
	rootCmd.Flags().BoolS("line", "line", false, "set output mode to 'line'")
	rootCmd.Flags().BoolS("list", "list", false, "set output mode to 'list'")
	rootCmd.Flags().StringS("lookaside", "lookaside", "", "use N entries of SZ bytes for lookaside memory")
	rootCmd.Flags().BoolS("markdown", "markdown", false, "set output mode to 'markdown'")
	rootCmd.Flags().StringS("maxsize", "maxsize", "", "maximum size for a --deserialize database")
	rootCmd.Flags().BoolS("memtrace", "memtrace", false, "trace all memory allocations and deallocations")
	rootCmd.Flags().StringS("mmap", "mmap", "", "default mmap size set to N")
	rootCmd.Flags().StringS("newline", "newline", "", "set output row separator. Default: '\\n'")
	rootCmd.Flags().BoolS("no-rowid-in-view", "no-rowid-in-view", false, "Disable rowid-in-view using sqlite3_config()")
	rootCmd.Flags().BoolS("nofollow", "nofollow", false, "refuse to open symbolic links to database files")
	rootCmd.Flags().StringS("nonce", "nonce", "", "set the safe-mode escape nonce")
	rootCmd.Flags().StringS("nullvalue", "nullvalue", "", "set text string for NULL values. Default ''")
	rootCmd.Flags().StringS("pagecache", "pagecache", "", "use N slots of SZ bytes each for page cache memory")
	rootCmd.Flags().BoolS("pcachetrace", "pcachetrace", false, "trace all page cache operations")
	rootCmd.Flags().BoolS("quote", "quote", false, "set output mode to 'quote'")
	rootCmd.Flags().BoolS("readonly", "readonly", false, "open the database read-only")
	rootCmd.Flags().BoolS("safe", "safe", false, "enable safe-mode")
	rootCmd.Flags().StringS("separator", "separator", "", "set output column separator. Default: '|")
	rootCmd.Flags().BoolS("stats", "stats", false, "print memory stats before each finalize")
	rootCmd.Flags().BoolS("table", "table", false, "set output mode to 'table'")
	rootCmd.Flags().BoolS("tabs", "tabs", false, "set output mode to 'tabs'")
	rootCmd.Flags().BoolS("unsafe-testing", "unsafe-testing", false, "allow unsafe commands and modes for testing")
	rootCmd.Flags().BoolS("version", "version", false, "show SQLite version")
	rootCmd.Flags().StringS("vfs", "vfs", "", "use NAME as the default VFS")
	rootCmd.Flags().BoolS("vfstrace", "vfstrace", false, "enable tracing of all VFS calls")
	rootCmd.Flags().BoolS("zip", "zip", false, "open the file as a ZIP Archive")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"init": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
