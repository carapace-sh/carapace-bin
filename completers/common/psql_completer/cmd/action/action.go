package action

import (
	"github.com/carapace-sh/carapace"
)

// generated from the psql help output (PostgreSQL 18)

var backslashCommands = []string{
	"\\!", "execute command in shell or start interactive shell",
	"\\?", "show help on backslash commands",
	"\\C", "set table title, or unset if none",
	"\\H", "",
	"\\T", "set HTML <table> tag attributes, or unset if none",
	"\\a", "toggle between unaligned and aligned output mode",
	"\\bind", "set query parameters",
	"\\bind_named", "",
	"\\c", "",
	"\\cd", "change the current working directory",
	"\\close_prepared", "",
	"\\connect", "",
	"\\conninfo", "display information about current connection",
	"\\copy", "perform SQL COPY with data stream to the client host",
	"\\copyright", "show PostgreSQL usage and distribution terms",
	"\\crosstabview", "",
	"\\d", "list tables, views, and sequences",
	"\\dA", "",
	"\\dAc", "list operator classes",
	"\\dAf", "list operator families",
	"\\dAo", "list operators of operator families",
	"\\dAp", "list support functions of operator families",
	"\\dC", "",
	"\\dD", "list domains",
	"\\dE", "list foreign tables",
	"\\dF", "",
	"\\dFd", "list text search dictionaries",
	"\\dFp", "list text search parsers",
	"\\dFt", "list text search templates",
	"\\dL", "list procedural languages",
	"\\dO", "list collations",
	"\\dP", "list [only index/table] partitioned relations [n=nested]",
	"\\dRp", "list replication publications",
	"\\dRs", "list replication subscriptions",
	"\\dT", "list data types",
	"\\dX", "",
	"\\da", "",
	"\\db", "",
	"\\dc", "list conversions",
	"\\dconfig", "",
	"\\dd", "",
	"\\ddp", "",
	"\\des", "list foreign servers",
	"\\det", "list foreign tables",
	"\\deu", "list user mappings",
	"\\dew", "list foreign-data wrappers",
	"\\df", "",
	"\\dg", "list roles",
	"\\di", "list indexes",
	"\\dl", "list large objects, same as \\lo_list",
	"\\dm", "list materialized views",
	"\\dn", "list schemas",
	"\\do", "",
	"\\dp", "",
	"\\drds", "",
	"\\drg", "list role grants",
	"\\ds", "list sequences",
	"\\dt", "list tables",
	"\\du", "list roles",
	"\\dv", "list views",
	"\\dx", "",
	"\\dy", "",
	"\\e", "edit the query buffer (or file) with external editor",
	"\\echo", "write string to standard output (-n for no newline)",
	"\\edit", "",
	"\\ef", "edit function definition with external editor",
	"\\elif", "alternative within current conditional block",
	"\\else", "final alternative within current conditional block",
	"\\encoding", "show or set client encoding",
	"\\endif", "end conditional block",
	"\\endpipeline", "exit pipeline mode",
	"\\errverbose", "show most recent error message at maximum verbosity",
	"\\ev", "edit view definition with external editor",
	"\\f", "show or set field separator for unaligned query output",
	"\\flush", "flush output data to the server",
	"\\flushrequest", "send request to the server to flush its output buffer",
	"\\g", "",
	"\\gdesc", "describe result of query, without executing it",
	"\\getenv", "",
	"\\getresults", "read NUM_RES pending results, or all if no argument",
	"\\gexec", "execute query, then execute each value in its result",
	"\\gset", "execute query and store result in psql variables",
	"\\gx", "",
	"\\h", "help on syntax of SQL commands, * for all commands",
	"\\help", "",
	"\\html", "",
	"\\i", "execute commands from file",
	"\\if", "begin conditional block",
	"\\include", "",
	"\\include_relative", "",
	"\\ir", "as \\i, but relative to location of current script",
	"\\l", "",
	"\\list", "",
	"\\listx", "",
	"\\lo_export", "",
	"\\lo_import", "",
	"\\lo_list", "list large objects",
	"\\lo_unlink", "delete a large object",
	"\\lx", "",
	"\\o", "send all query results to file or |pipe",
	"\\out", "",
	"\\p", "show the contents of the query buffer",
	"\\parse", "create a prepared statement",
	"\\password", "securely change the password for a user",
	"\\print", "",
	"\\prompt", "prompt user to set internal variable",
	"\\pset", "set table output option",
	"\\q", "quit psql",
	"\\qecho", "write string to \\o output stream (-n for no newline)",
	"\\quit", "",
	"\\r", "reset (clear) the query buffer",
	"\\reset", "",
	"\\restrict", "",
	"\\s", "display history or save it to file",
	"\\sendpipeline", "send an extended query to an ongoing pipeline",
	"\\set", "set internal variable, or list all if no parameters",
	"\\setenv", "set or unset environment variable",
	"\\sf", "",
	"\\startpipeline", "enter pipeline mode",
	"\\sv", "",
	"\\syncpipeline", "add a synchronisation point to an ongoing pipeline",
	"\\t", "",
	"\\timing", "",
	"\\unrestrict", "",
	"\\unset", "unset (delete) internal variable",
	"\\w", "write query buffer to file",
	"\\warn", "write string to standard error (-n for no newline)",
	"\\watch", "",
	"\\write", "",
	"\\x", "",
	"\\z", "",
	"\\zS", "",
	"\\zSx", "",
	"\\zx", "",
}

var variables = []string{
	"AUTOCOMMIT", "if set, successful SQL commands are automatically committed",
	"COLUMNS", "number of columns for wrapped format",
	"COMP_KEYWORD_CASE", "determines the case used to complete SQL key words",
	"DBNAME", "the currently connected database name",
	"ECHO", "controls what input is written to standard output",
	"ECHO_HIDDEN", "if set, display internal queries executed by backslash commands;",
	"ENCODING", "current client character set encoding",
	"ERROR", "\"true\" if last query failed, else \"false\"",
	"FETCH_COUNT", "the number of result rows to fetch and display at a time (0 = unlimited)",
	"HIDE_TABLEAM", "if set, table access methods are not displayed",
	"HIDE_TOAST_COMPRESSION", "if set, compression methods are not displayed",
	"HISTCONTROL", "controls command history [ignorespace, ignoredups, ignoreboth]",
	"HISTFILE", "file name used to store the command history",
	"HISTSIZE", "maximum number of commands to store in the command history",
	"HOST", "the currently connected database server host",
	"IGNOREEOF", "number of EOFs needed to terminate an interactive session",
	"LASTOID", "value of the last affected OID",
	"LAST_ERROR_MESSAGE", "LAST_ERROR_SQLSTATE",
	"ON_ERROR_ROLLBACK", "if set, an error doesn't stop a transaction (uses implicit savepoints)",
	"ON_ERROR_STOP", "stop batch execution after error",
	"PGAPPNAME", "same as the application_name connection parameter",
	"PGDATABASE", "same as the dbname connection parameter",
	"PGHOST", "same as the host connection parameter",
	"PGPASSFILE", "password file name",
	"PGPASSWORD", "connection password (not recommended)",
	"PGPORT", "same as the port connection parameter",
	"PGUSER", "same as the user connection parameter",
	"PORT", "server port of the current connection",
	"PROMPT1", "specifies the standard psql prompt",
	"PROMPT2", "specifies the prompt used when a statement continues from a previous line",
	"PROMPT3", "specifies the prompt used during COPY ... FROM STDIN",
	"PSQLRC", "alternative location for the user's .psqlrc file",
	"PSQL_EDITOR_LINENUMBER_ARG", "how to specify a line number when invoking the editor",
	"PSQL_HISTORY", "alternative location for the command history file",
	"PSQL_WATCH_PAGER", "name of external pager program used for \\watch",
	"QUIET", "run quietly (same as -q option)",
	"ROW_COUNT", "number of rows returned or affected by last query, or 0",
	"SERVER_VERSION_NAME", "SERVER_VERSION_NUM",
	"SHELL", "shell used by the \\! command",
	"SHELL_ERROR", "\"true\" if the last shell command failed, \"false\" if it succeeded",
	"SHELL_EXIT_CODE", "exit status of the last shell command",
	"SHOW_ALL_RESULTS", "show all results of a combined query (\\;) instead of only the last",
	"SHOW_CONTEXT", "controls display of message context fields [never, errors, always]",
	"SINGLELINE", "if set, end of line terminates SQL commands (same as -S option)",
	"SINGLESTEP", "single-step mode (same as -s option)",
	"SQLSTATE", "SQLSTATE of last query, or \"00000\" if no error",
	"TMPDIR", "directory for temporary files",
	"USER", "the currently connected database user",
	"VERBOSITY", "controls verbosity of error reports [default, verbose, terse, sqlstate]",
	"VERSION", "VERSION_NAME",
}

var printOptions = []string{
	"border",
	"columns",
	"csv_fieldsep",
	"expanded",
	"fieldsep",
	"fieldsep_zero",
	"footer",
	"format",
	"linestyle",
	"null",
	"numericlocale",
	"pager",
	"pager_min_lines",
	"recordsep",
	"recordsep_zero",
	"tableattr",
	"title",
	"tuples_only",
	"unicode_border_linestyle",
	"unicode_column_linestyle",
	"unicode_header_linestyle",
	"xheader_width",
}

// ActionBackslashCommands completes psql backslash commands
func ActionBackslashCommands() carapace.Action {
	return carapace.ActionValuesDescribed(backslashCommands...).Tag("backslash commands")
}

// ActionVariables completes psql special variables
func ActionVariables() carapace.Action {
	return carapace.ActionValuesDescribed(variables...).Tag("variables")
}

// ActionPrintOptions completes the options of \pset
func ActionPrintOptions() carapace.Action {
	return carapace.ActionValues(printOptions...).Tag("print options")
}

// ActionPrintOptionValues completes the values of a \pset option
func ActionPrintOptionValues(option string) carapace.Action {
	switch option {
	case "border":
		return carapace.ActionValues("0", "1", "2", "3")
	case "expanded":
		return carapace.ActionValues("auto", "off", "on")
	case "format":
		return carapace.ActionValues(
			"aligned",
			"asciidoc",
			"csv",
			"html",
			"latex",
			"latex-longtable",
			"troff-ms",
			"unaligned",
			"wrapped",
		)
	case "linestyle":
		return carapace.ActionValues("ascii", "old-ascii", "unicode")
	case "pager":
		return carapace.ActionValues("always", "off", "on")
	case "unicode_border_linestyle", "unicode_column_linestyle", "unicode_header_linestyle":
		return carapace.ActionValues("double", "single")
	case "xheader_width":
		return carapace.ActionValues("column", "full", "page")
	default:
		return carapace.ActionValues("off", "on")
	}
}
