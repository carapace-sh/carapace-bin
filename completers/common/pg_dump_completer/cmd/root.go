package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pg_dump",
	Short: "Export a PostgreSQL database as an SQL script or to other formats",
	Long:  "https://www.postgresql.org/docs/current/app-pgdump.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("binary-upgrade", false, "for use by upgrade utilities only")
	rootCmd.Flags().Bool("blobs", false, "(same as --large-objects, deprecated)")
	rootCmd.Flags().BoolP("clean", "c", false, "clean (drop) database objects before recreating")
	rootCmd.Flags().Bool("column-inserts", false, "dump data as INSERT commands with column names")
	rootCmd.Flags().StringP("compress", "Z", "", "compress as specified")
	rootCmd.Flags().BoolP("create", "C", false, "include commands to create database in dump")
	rootCmd.Flags().BoolP("data-only", "a", false, "dump only the data, not the schema or statistics")
	rootCmd.Flags().StringP("dbname", "d", "", "database to dump")
	rootCmd.Flags().Bool("disable-dollar-quoting", false, "disable dollar quoting, use SQL standard quoting")
	rootCmd.Flags().Bool("disable-triggers", false, "disable triggers during data-only restore")
	rootCmd.Flags().Bool("enable-row-security", false, "enable row security (dump only content user has access to)")
	rootCmd.Flags().StringP("encoding", "E", "", "dump the data in encoding ENCODING")
	rootCmd.Flags().String("exclude-extension", "", "do NOT dump the specified extension(s)")
	rootCmd.Flags().StringP("exclude-schema", "N", "", "do NOT dump the specified schema(s)")
	rootCmd.Flags().StringP("exclude-table", "T", "", "do NOT dump the specified table(s)")
	rootCmd.Flags().String("exclude-table-and-children", "", "do NOT dump the specified table(s), including child and partition tables")
	rootCmd.Flags().String("exclude-table-data", "", "do NOT dump data for the specified table(s)")
	rootCmd.Flags().String("exclude-table-data-and-children", "", "do NOT dump data for the specified table(s), including child and partition tables")
	rootCmd.Flags().StringP("extension", "e", "", "dump the specified extension(s) only")
	rootCmd.Flags().String("extra-float-digits", "", "override default setting for extra_float_digits")
	rootCmd.Flags().StringP("file", "f", "", "output file or directory name")
	rootCmd.Flags().String("filter", "", "include or exclude objects and data from dump based on expressions in FILENAME")
	rootCmd.Flags().StringP("format", "F", "", "output file format (custom, directory, tar, plain text (default))")
	rootCmd.Flags().BoolP("help", "?", false, "show this help, then exit")
	rootCmd.Flags().StringP("host", "h", "", "database server host or socket directory")
	rootCmd.Flags().Bool("if-exists", false, "use IF EXISTS when dropping objects")
	rootCmd.Flags().String("include-foreign-data", "", "include data of foreign tables on foreign servers matching PATTERN")
	rootCmd.Flags().Bool("inserts", false, "dump data as INSERT commands, rather than COPY")
	rootCmd.Flags().StringP("jobs", "j", "", "use this many parallel jobs to dump")
	rootCmd.Flags().BoolP("large-objects", "b", false, "include large objects in dump")
	rootCmd.Flags().Bool("load-via-partition-root", false, "load partitions via the root table")
	rootCmd.Flags().String("lock-wait-timeout", "", "fail after waiting TIMEOUT for a table lock")
	rootCmd.Flags().Bool("no-blobs", false, "(same as --no-large-objects, deprecated)")
	rootCmd.Flags().Bool("no-comments", false, "do not dump comment commands")
	rootCmd.Flags().Bool("no-data", false, "do not dump data")
	rootCmd.Flags().BoolP("no-large-objects", "B", false, "exclude large objects in dump")
	rootCmd.Flags().BoolP("no-owner", "O", false, "skip restoration of object ownership in plain-text format")
	rootCmd.Flags().BoolP("no-password", "w", false, "never prompt for password")
	rootCmd.Flags().Bool("no-policies", false, "do not dump row security policies")
	rootCmd.Flags().BoolP("no-privileges", "x", false, "do not dump privileges (grant/revoke)")
	rootCmd.Flags().Bool("no-publications", false, "do not dump publications")
	rootCmd.Flags().Bool("no-schema", false, "do not dump schema")
	rootCmd.Flags().Bool("no-security-labels", false, "do not dump security label assignments")
	rootCmd.Flags().Bool("no-statistics", false, "do not dump statistics")
	rootCmd.Flags().Bool("no-subscriptions", false, "do not dump subscriptions")
	rootCmd.Flags().Bool("no-sync", false, "do not wait for changes to be written safely to disk")
	rootCmd.Flags().Bool("no-table-access-method", false, "do not dump table access methods")
	rootCmd.Flags().Bool("no-tablespaces", false, "do not dump tablespace assignments")
	rootCmd.Flags().Bool("no-toast-compression", false, "do not dump TOAST compression methods")
	rootCmd.Flags().Bool("no-unlogged-table-data", false, "do not dump unlogged table data")
	rootCmd.Flags().Bool("on-conflict-do-nothing", false, "add ON CONFLICT DO NOTHING to INSERT commands")
	rootCmd.Flags().BoolP("password", "W", false, "force password prompt (should happen automatically)")
	rootCmd.Flags().StringP("port", "p", "", "database server port number")
	rootCmd.Flags().Bool("quote-all-identifiers", false, "quote all identifiers, even if not key words")
	rootCmd.Flags().String("restrict-key", "", "use provided string as psql \\restrict key")
	rootCmd.Flags().String("role", "", "do SET ROLE before dump")
	rootCmd.Flags().String("rows-per-insert", "", "number of rows per INSERT; implies --inserts")
	rootCmd.Flags().StringP("schema", "n", "", "dump the specified schema(s) only")
	rootCmd.Flags().BoolP("schema-only", "s", false, "dump only the schema, no data or statistics")
	rootCmd.Flags().String("section", "", "dump named section (pre-data, data, or post-data)")
	rootCmd.Flags().Bool("sequence-data", false, "include sequence data in dump")
	rootCmd.Flags().Bool("serializable-deferrable", false, "wait until the dump can run without anomalies")
	rootCmd.Flags().String("snapshot", "", "use given snapshot for the dump")
	rootCmd.Flags().Bool("statistics", false, "dump the statistics")
	rootCmd.Flags().Bool("statistics-only", false, "dump only the statistics, not schema or data")
	rootCmd.Flags().Bool("strict-names", false, "require table and/or schema include patterns to match at least one entity each")
	rootCmd.Flags().StringP("superuser", "S", "", "superuser user name to use in plain-text format")
	rootCmd.Flags().String("sync-method", "", "set method for syncing files to disk")
	rootCmd.Flags().StringP("table", "t", "", "dump only the specified table(s)")
	rootCmd.Flags().String("table-and-children", "", "dump only the specified table(s), including child and partition tables")
	rootCmd.Flags().Bool("use-set-session-authorization", false, "use SET SESSION AUTHORIZATION commands instead of ALTER OWNER commands to set ownership")
	rootCmd.Flags().StringP("username", "U", "", "connect as specified database user")
	rootCmd.Flags().BoolP("verbose", "v", false, "verbose mode")
	rootCmd.Flags().BoolP("version", "V", false, "output version information, then exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"compress": carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues("gzip", "lz4", "none", "zstd")
			default:
				return carapace.ActionValues()
			}
		}),
		"file":   carapace.ActionFiles(),
		"filter": carapace.ActionFiles(),
		"format": carapace.ActionValuesDescribed(
			"c", "custom",
			"d", "directory",
			"p", "plain text",
			"t", "tar",
		),
		"section":     carapace.ActionValues("data", "post-data", "pre-data"),
		"sync-method": carapace.ActionValues("fsync", "syncfs"),
	})
}
