package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fd",
	Short: "find entries in the filesystem",
	Long:  "https://github.com/sharkdp/fd",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("1", "", "", "Limit to one result")
	rootCmd.Flags().BoolP("absolute-path", "a", false, "Show absolute instead of relative paths")
	rootCmd.Flags().StringP("and", "", "", "Add additional required search patterns")
	rootCmd.Flags().StringP("base-directory", "", "", "Change theb base directory")
	rootCmd.Flags().StringP("batch-size", "", "", "Maximum number of arguments for batch commands")
	rootCmd.Flags().BoolP("case-sensitive", "s", false, "Case-sensitive search (default: smart case)")
	rootCmd.Flags().String("changed-before", "", "Filter by file modification time (older than)")
	rootCmd.Flags().String("changed-within", "", "Filter by file modification time (newer than)")
	rootCmd.Flags().StringP("color", "c", "", "When to use colors")
	rootCmd.Flags().StringP("exact-depth", "", "", "When to use colors")
	rootCmd.Flags().StringP("exclude", "E", "", "Exclude entries that match the given glob pattern")
	rootCmd.Flags().StringP("exec", "x", "", "Execute a command for each search result")
	rootCmd.Flags().StringP("exec-batch", "X", "", "Execute a command with all search results at once")
	rootCmd.Flags().StringP("extension", "e", "", "Filter by file extension")
	rootCmd.Flags().BoolP("fixed-strings", "F", false, "fixed string search")
	rootCmd.Flags().BoolP("follow", "L", false, "Follow symbolic links")
	rootCmd.Flags().StringP("format", "", "", "Print results according to template")
	rootCmd.Flags().BoolP("full-path", "p", false, "Search full path (default: file-/dirname only)")
	rootCmd.Flags().BoolP("glob", "g", false, "Glob-based search (default: regular expression)")
	rootCmd.Flags().BoolP("help", "h", false, "Prints help information")
	rootCmd.Flags().BoolP("hidden", "H", false, "Search hidden files and directories")
	rootCmd.Flags().StringP("hyperlink", "", "", "When to add terminal hyperlinks")
	rootCmd.Flags().BoolP("ignore-case", "i", false, "Case-insensitive search (default: smart case)")
	rootCmd.Flags().StringP("ignore-file", "", "", "Add a custom ignore file.")
	rootCmd.Flags().BoolP("list-details", "l", false, "Use a long listing format with file metadata")
	rootCmd.Flags().StringP("max-depth", "d", "", "Set maximum search depth (default: none)")
	rootCmd.Flags().StringP("max-results", "", "", "Limit the number of results")
	rootCmd.Flags().StringP("min-depth", "", "", "Set minimum search depth (default: none)")
	rootCmd.Flags().BoolP("no-ignore", "I", false, "Do not respect .(git|fd)ignore files")
	rootCmd.Flags().BoolP("no-ignore-parent", "", false, "Ignore parent directory's .gitignore")
	rootCmd.Flags().BoolP("no-ignore-vcs", "", false, "Ignore .gitignore files")
	rootCmd.Flags().BoolP("no-require-git", "", false, "Do not require a git repository to respect gitignores")
	rootCmd.Flags().BoolP("one-file-system", "", false, "Stay within the same file system")
	rootCmd.Flags().StringP("owner", "o", "", "Filter by owning user and/or group")
	rootCmd.Flags().StringP("path-separator", "", "", "Set the path separator")
	rootCmd.Flags().BoolP("print0", "0", false, "Separate results by the null character")
	rootCmd.Flags().StringP("prune", "", "", "Do not recurse into matching directories")
	rootCmd.Flags().BoolP("quiet", "q", false, "Quiet mode")
	rootCmd.Flags().StringP("regex", "", "", "Regular expression search")
	rootCmd.Flags().StringP("search-path", "", "", "Provide paths to search")
	rootCmd.Flags().BoolP("show-errors", "", false, "Show filesystem errors")
	rootCmd.Flags().StringP("size", "S", "", "Limit results based on the size of files.")
	rootCmd.Flags().StringP("strip-cwd-prefix", "", "", "Strip current working directory prefix")
	rootCmd.Flags().StringP("threads", "j", "", "Set number of threads")
	rootCmd.Flags().StringP("type", "t", "", "Filter by type")
	rootCmd.Flags().BoolP("unrestricted", "u", false, "Unrestricted Search")
	rootCmd.Flags().BoolP("version", "V", false, "Prints version information")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"color":            carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"exec":             carapace.ActionFiles(),
		"exec-batch":       carapace.ActionFiles(),
		"extension":        fs.ActionFilenameExtensions(),
		"hyperlink":        carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"owner":            os.ActionUserGroup(),
		"strip-cwd-prefix": carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"type": carapace.ActionValues(
			"file",
			"directory",
			"symlink",
			"executable",
			"empty",
			"socket",
			"pipe",
		),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues(),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionDirectories(),
	)
}
