package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/psql_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "psql",
	Short: "PostgreSQL interactive terminal",
	Long:  "https://www.postgresql.org/docs/current/app-psql.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("command", "c", "", "run only single command (SQL or internal) and exit")
	rootCmd.Flags().Bool("csv", false, "CSV (Comma-Separated Values) table output mode")
	rootCmd.Flags().StringP("dbname", "d", "", "database name to connect to")
	rootCmd.Flags().BoolP("echo-all", "a", false, "echo all input from script")
	rootCmd.Flags().BoolP("echo-errors", "b", false, "echo failed commands")
	rootCmd.Flags().BoolP("echo-hidden", "E", false, "display queries that internal commands generate")
	rootCmd.Flags().BoolP("echo-queries", "e", false, "echo commands sent to server")
	rootCmd.Flags().BoolP("expanded", "x", false, "turn on expanded table output")
	rootCmd.Flags().StringP("field-separator", "F", "", "field separator for unaligned output")
	rootCmd.Flags().BoolP("field-separator-zero", "z", false, "set field separator for unaligned output to zero byte")
	rootCmd.Flags().StringP("file", "f", "", "execute commands from file, then exit")
	rootCmd.Flags().StringP("help", "?", "", "show help, then exit")
	rootCmd.Flags().StringP("host", "h", "", "database server host or socket directory")
	rootCmd.Flags().BoolP("html", "H", false, "HTML table output mode")
	rootCmd.Flags().BoolP("list", "l", false, "list available databases, then exit")
	rootCmd.Flags().StringP("log-file", "L", "", "send session log to file")
	rootCmd.Flags().BoolP("no-align", "A", false, "unaligned table output mode")
	rootCmd.Flags().BoolP("no-password", "w", false, "never prompt for password")
	rootCmd.Flags().BoolP("no-psqlrc", "X", false, "do not read startup file (~/.psqlrc)")
	rootCmd.Flags().BoolP("no-readline", "n", false, "disable enhanced command line editing (readline)")
	rootCmd.Flags().StringP("output", "o", "", "send query results to file (or |pipe)")
	rootCmd.Flags().BoolP("password", "W", false, "force password prompt (should happen automatically)")
	rootCmd.Flags().StringP("port", "p", "", "database server port")
	rootCmd.Flags().StringP("pset", "P", "", "set printing option VAR to ARG (see \\pset command)")
	rootCmd.Flags().BoolP("quiet", "q", false, "run quietly (no messages, only query output)")
	rootCmd.Flags().StringP("record-separator", "R", "", "record separator for unaligned output (default: newline)")
	rootCmd.Flags().BoolP("record-separator-zero", "0", false, "set record separator for unaligned output to zero byte")
	rootCmd.Flags().StringArrayP("set", "v", nil, "set psql variable NAME to VALUE")
	rootCmd.Flags().BoolP("single-line", "S", false, "single-line mode (end of line terminates SQL command)")
	rootCmd.Flags().BoolP("single-step", "s", false, "single-step mode (confirm each query)")
	rootCmd.Flags().BoolP("single-transaction", "1", false, "execute as a single transaction (if non-interactive)")
	rootCmd.Flags().StringP("table-attr", "T", "", "set HTML table tag attributes, or unset if none")
	rootCmd.Flags().BoolP("tuples-only", "t", false, "print rows only")
	rootCmd.Flags().StringP("username", "U", "", "database user name")
	rootCmd.Flags().StringArray("variable", nil, "set psql variable NAME to VALUE")
	rootCmd.Flags().BoolP("version", "V", false, "output version information, then exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"command": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if strings.HasPrefix(c.Value, "\\") {
				return action.ActionBackslashCommands()
			}
			return carapace.ActionValues()
		}),
		"file":     carapace.ActionFiles(),
		"help":     carapace.ActionValues("commands", "options", "variables"),
		"log-file": carapace.ActionFiles(),
		"output":   carapace.ActionFiles(),
		"pset":     actionPrintOption(),
		"set":      actionVariableValue(),
		"username": os.ActionUsers(),
		"variable": actionVariableValue(),
	})
}

func actionPrintOption() carapace.Action {
	return carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return action.ActionPrintOptions()
		case 1:
			return action.ActionPrintOptionValues(c.Parts[0])
		default:
			return carapace.ActionValues()
		}
	})
}

func actionVariableValue() carapace.Action {
	return carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return action.ActionVariables()
		case 1:
			switch c.Parts[0] {
			case "ECHO":
				return carapace.ActionValues("all", "errors", "none", "queries")
			case "COMP_KEYWORD_CASE":
				return carapace.ActionValues("lower", "preserve-lower", "preserve-upper", "upper")
			case "HISTCONTROL":
				return carapace.ActionValues("ignoreboth", "ignoredups", "ignorespace")
			case "ON_ERROR_ROLLBACK":
				return carapace.ActionValues("interactive", "off", "on")
			case "VERBOSITY":
				return carapace.ActionValues("default", "terse", "verbose")
			default:
				return carapace.ActionValues("off", "on")
			}
		default:
			return carapace.ActionValues()
		}
	})
}
