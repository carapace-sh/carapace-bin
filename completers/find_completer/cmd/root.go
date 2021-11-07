package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/fs"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "find",
	Short: "search for files in a directory hierarchy",
	Long:  "https://linux.die.net/man/1/find",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	carapace.Override(carapace.Opts{
		LongShorthand: true,
	})
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("-help", false, "Print a summary of the command-line usage of find and exit.")
	rootCmd.Flags().Bool("-version", false, "Print the find version number and exit.")
	rootCmd.Flags().String("D", "", "Print diagnostic information")
	rootCmd.Flags().Bool("H", false, "Do not follow symbolic links")
	rootCmd.Flags().Bool("L", false, "Follow symbolic links")
	rootCmd.Flags().Bool("O0", false, "Equivalent to optimisation level 1")
	rootCmd.Flags().Bool("O1", false, "This  is  the default optimisation level and corresponds to the traditional behaviour.")
	rootCmd.Flags().Bool("O2", false, "Any -type or -xtype tests are performed after any tests based only on the names of files.")
	rootCmd.Flags().Bool("O3", false, "At  this  optimisation level, the full cost-based query optimiser is enabled.")
	rootCmd.Flags().Bool("P", false, "Never follow symbolic links")
	rootCmd.Flags().String("amin", "", "File was last accessed less than, more than or exactly n minutes ago.")
	rootCmd.Flags().String("anewer", "", "Time of the last access of the current file is more recent than that  of  the  last  data modification of the reference file.")
	rootCmd.Flags().String("atime", "", "File was last accessed less than, more than or exactly n*24 hours ago.")
	rootCmd.Flags().String("cmin", "", "File's status was last changed less than, more than or exactly n minutes ago.")
	rootCmd.Flags().String("cnewer", "", "Time  of  the last status change of the current file is more recent than that of the last data modification of the reference file.")
	rootCmd.Flags().String("context", "", "Security context of the file matches glob pattern.")
	rootCmd.Flags().String("ctime", "", "File's status was last changed less than, more than or exactly n*24 hours ago.")
	rootCmd.Flags().Bool("daystart", false, "Measure times from the beginning of today rather than from 24 hours ago.")
	rootCmd.Flags().Bool("delete", false, "Delete  files.")
	rootCmd.Flags().Bool("depth", false, "Process each directory's contents before the directory itself.")
	rootCmd.Flags().String("exec", "", "Execute command.")
	rootCmd.Flags().String("execdir", "", "Like -exec, but the specified command is run from the subdirectory containing the matched file.")
	rootCmd.Flags().Bool("executable", false, "Matches files which are executable and directories which are searchable by the current user.")
	rootCmd.Flags().Bool("false", false, "Always false.")
	rootCmd.Flags().Bool("fls", false, "Unusual characters are always escaped.")
	rootCmd.Flags().Bool("fprint", false, "Quoting is handled in the same way as for -printf and -fprintf.")
	rootCmd.Flags().Bool("fprint0", false, "Always print the exact filename, unchanged, even if the output is going to a terminal.")
	rootCmd.Flags().Bool("fprintf", false, "If the output is not going to a terminal, it is printed as-is.")
	rootCmd.Flags().String("fstype", "", "File  is  on  a filesystem of type type.")
	rootCmd.Flags().String("gid", "", "File's numeric group ID is less than, more than or exactly n.")
	rootCmd.Flags().String("group", "", "File belongs to group gname (numeric group ID allowed).")
	rootCmd.Flags().Bool("help", false, "Print a summary of the command-line usage of find and exit.")
	rootCmd.Flags().Bool("ignore_readdir_race", false, "Do not  emit an error message when find fails to stat a file.")
	rootCmd.Flags().String("ilname", "", "Like -lname, but the match is case insensitive.")
	rootCmd.Flags().String("iname", "", "Like -name, but the match is case insensitive.")
	rootCmd.Flags().String("inum", "", "File  has inode number smaller than, greater than or exactly n.")
	rootCmd.Flags().String("ipath", "", "Like -path.  but the match is case insensitive.")
	rootCmd.Flags().String("iregex", "", "Like -regex, but the match is case insensitive.")
	rootCmd.Flags().String("iwholename", "", "See -ipath.  This alternative is less portable than -ipath.")
	rootCmd.Flags().String("links", "", "File has less than, more than or exactly n hard links.")
	rootCmd.Flags().String("lname", "", "File is a symbolic link whose contents match shell pattern pattern.")
	rootCmd.Flags().Bool("ls", false, "List current file in ls -dils format on standard output.")
	rootCmd.Flags().String("maxdepth", "", "Descend at most levels of directories below the startingpoints.")
	rootCmd.Flags().String("mindepth", "", "Do  not  apply  any tests or actions at levels less than levels.")
	rootCmd.Flags().String("mmin", "", "File's data was last modified less than, more than or exactly n minutes ago.")
	rootCmd.Flags().Bool("mount", false, "Don't descend directories on other filesystems.")
	rootCmd.Flags().String("mtime", "", "File's data was last modified less than, more than or exactly n*24 hours  ago.")
	rootCmd.Flags().String("name", "", "Base of file name matches  shell  pattern pattern.")
	rootCmd.Flags().String("newer", "", "Time of the last data modification of the current file is more recent than that of the last data modification of the reference file.")
	// TODO XY completion
	//rootCmd.Flags().String("newerXY", "", "Succeeds if timestamp X of the file being considered is newer than timestamp  Y  of  the file reference.")
	rootCmd.Flags().Bool("nogroup", false, "No group corresponds to file's numeric group ID.")
	rootCmd.Flags().Bool("noignore_readdir_race", false, "Turns off the effect of -ignore_readdir_race.")
	rootCmd.Flags().Bool("noleaf", false, "Do  not  optimize  by assuming that directories contain 2 fewer subdirectories than their hard link count.")
	rootCmd.Flags().Bool("nouser", false, "No user corresponds to file's numeric user ID.")
	rootCmd.Flags().Bool("nowarn", false, "Turn  warning  messages  off.")
	rootCmd.Flags().String("ok", "", "Like  -exec but ask the user first.")
	rootCmd.Flags().String("okdir", "", "Like -execdir but ask the user first in the same way as for -ok.")
	rootCmd.Flags().String("path", "", "File name matches shell pattern pattern.")
	rootCmd.Flags().String("perm", "", "File's permission bits are exactly mode (octal or symbolic).")
	rootCmd.Flags().Bool("print", false, "Print the full file name on the standard output, followed by a newline.")
	rootCmd.Flags().Bool("print0", false, "Print the full file name on the standard output, followed by a null character.")
	rootCmd.Flags().String("printf", "", "Print  format on the standard output, interpreting `\\' escapes and `%' directives.")
	rootCmd.Flags().Bool("prune", false, "If the file is a directory, do not descend into it.")
	rootCmd.Flags().Bool("quit", false, "Exit immediately.")
	rootCmd.Flags().Bool("readable", false, "Matches  files  which  are  readable by the current user.")
	rootCmd.Flags().String("regex", "", "File  name  matches regular expression pattern.")
	rootCmd.Flags().String("regextype", "", "Changes the regular expression syntax understood by -regex and -iregex tests which  occur later  on the command line.")
	rootCmd.Flags().String("samefile", "", "File  refers  to the same inode as name.")
	rootCmd.Flags().String("size", "", "File uses less than, more than or exactly n units of space, rounding up.")
	rootCmd.Flags().Bool("true", false, "Always true.")
	rootCmd.Flags().String("type", "", "file type")
	rootCmd.Flags().String("used", "", "File  was  last accessed less than, more than or exactly n days after its status was last changed.")
	rootCmd.Flags().String("user", "", "File is owned by user uname.")
	rootCmd.Flags().Bool("version", false, "Print the find version number and exit.")
	rootCmd.Flags().Bool("warn", false, "Turn  warning  messages  on.")
	rootCmd.Flags().String("wholename", "", "See -path.  This alternative is less portable than -path.")
	rootCmd.Flags().Bool("writable", false, "Matches files which are writable by the current user.")
	rootCmd.Flags().Bool("xdev", false, "Don't descend directories on other filesystems.")
	rootCmd.Flags().String("xtype", "", "The same as -type unless the file is a symbolic link.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"D": carapace.ActionValuesDescribed(
			"exec", "Show diagnostic information relating to -exec, -execdir, -ok and -okdir",
			"opt", "Prints diagnostic information relating to the optimisation of the expression tree",
			"rates", "Prints a summary indicating how often each predicate succeeded or failed.",
			"search", "Navigate the directory tree verbosely.",
			"stat", "Print  messages  as  files are examined with the stat and lstat system calls.",
			"tree", "Show the expression tree in its original and optimised form.",
			"all", "Enable all of the other debug options (but help).",
			"help", "Explain the debugging options.",
		),
		"fstype":    fs.ActionFileSystemTypes(),
		"group":     os.ActionGroups(),
		"perm":      fs.ActionFileModes(),
		"regextype": carapace.ActionValues("findutils-default", "ed", "emacs", "gnu-awk", "grep", "posix-awk", "awk", "posix-basic", "posix-egrep", "egrep", "posix-extended", "posix-minimal-basic", "sed"),
		"type": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return ActionFileTypes().Invoke(c).Filter(c.Parts).ToA()
		}),
		"user": os.ActionUsers(),
		"xtype": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return ActionFileTypes().Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionDirectories(),
	)
}

func ActionFileTypes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"b", "block (buffered) special",
		"c", "character (unbuffered) special",
		"d", "directory",
		"p", "named pipe (FIFO)",
		"f", "regular file",
		"l", "symbolic link",
		"s", "socket",
		"D", "door (Solaris)",
	)
}
