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

	rootCmd.Flags().BoolS("1", "1", false, "Limit to one result")
	rootCmd.Flags().BoolP("absolute-path", "a", false, "Show absolute instead of relative paths")
	rootCmd.Flags().String("and", "", "Add additional required search patterns")
	rootCmd.Flags().String("base-directory", "", "Change theb base directory")
	rootCmd.Flags().String("batch-size", "", "Maximum number of arguments for batch commands")
	rootCmd.Flags().BoolP("case-sensitive", "s", false, "Case-sensitive search (default: smart case)")
	rootCmd.Flags().String("changed-after", "", "alias for --changed-within")
	rootCmd.Flags().String("changed-before", "", "Filter by file modification time (older than)")
	rootCmd.Flags().String("changed-within", "", "Filter by file modification time (newer than)")
	rootCmd.Flags().StringP("color", "c", "", "When to use colors")
	rootCmd.Flags().String("exact-depth", "", "When to use colors")
	rootCmd.Flags().StringP("exclude", "E", "", "Exclude entries that match the given glob pattern")
	rootCmd.Flags().StringP("exec", "x", "", "Execute a command for each search result")
	rootCmd.Flags().StringP("exec-batch", "X", "", "Execute a command with all search results at once")
	rootCmd.Flags().StringP("extension", "e", "", "Filter by file extension")
	rootCmd.Flags().BoolP("fixed-strings", "F", false, "fixed string search")
	rootCmd.Flags().BoolP("follow", "L", false, "Follow symbolic links")
	rootCmd.Flags().String("format", "", "Print results according to template")
	rootCmd.Flags().BoolP("full-path", "p", false, "Search full path (default: file-/dirname only)")
	rootCmd.Flags().BoolP("glob", "g", false, "Glob-based search (default: regular expression)")
	rootCmd.Flags().BoolP("help", "h", false, "Prints help information")
	rootCmd.Flags().BoolP("hidden", "H", false, "Search hidden files and directories")
	rootCmd.Flags().String("hyperlink", "", "When to add terminal hyperlinks")
	rootCmd.Flags().BoolP("ignore-case", "i", false, "Case-insensitive search (default: smart case)")
	rootCmd.Flags().String("ignore-file", "", "Add a custom ignore file.")
	rootCmd.Flags().BoolP("list-details", "l", false, "Use a long listing format with file metadata")
	rootCmd.Flags().StringP("max-depth", "d", "", "Set maximum search depth (default: none)")
	rootCmd.Flags().String("max-results", "", "Limit the number of results")
	rootCmd.Flags().String("min-depth", "", "Set minimum search depth (default: none)")
	rootCmd.Flags().BoolP("no-ignore", "I", false, "Do not respect .(git|fd)ignore files")
	rootCmd.Flags().Bool("no-ignore-parent", false, "Ignore parent directory's .gitignore")
	rootCmd.Flags().Bool("no-ignore-vcs", false, "Ignore .gitignore files")
	rootCmd.Flags().Bool("no-require-git", false, "Do not require a git repository to respect gitignores")
	rootCmd.Flags().Bool("one-file-system", false, "Stay within the same file system")
	rootCmd.Flags().StringP("owner", "o", "", "Filter by owning user and/or group")
	rootCmd.Flags().String("path-separator", "", "Set the path separator")
	rootCmd.Flags().BoolP("print0", "0", false, "Separate results by the null character")
	rootCmd.Flags().Bool("prune", false, "Do not recurse into matching directories")
	rootCmd.Flags().BoolP("quiet", "q", false, "Quiet mode")
	rootCmd.Flags().String("regex", "", "Regular expression search")
	rootCmd.Flags().StringSliceP("search-path", "", nil, "Provide paths to search")
	rootCmd.Flags().Bool("show-errors", false, "Show filesystem errors")
	rootCmd.Flags().StringP("size", "S", "", "Limit results based on the size of files.")
	rootCmd.Flags().String("strip-cwd-prefix", "", "Strip current working directory prefix")
	rootCmd.Flags().StringP("threads", "j", "", "Set number of threads")
	rootCmd.Flags().StringP("type", "t", "", "Filter by type")
	rootCmd.Flags().BoolP("unrestricted", "u", false, "Unrestricted Search")
	rootCmd.Flags().BoolP("version", "V", false, "Prints version information")

	rootCmd.MarkFlagsMutuallyExclusive("changed-after", "changed-within")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"base-directory": carapace.ActionDirectories(),
		"color":          carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"exec":           carapace.ActionFiles(),
		"exec-batch":     carapace.ActionFiles(),
		"extension":      fs.ActionFilenameExtensions(),
		"hyperlink":      carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"owner":          os.ActionUserGroup(),
		"path-separator": carapace.ActionValuesDescribed(
			"/", "Unix",
			`\`, "Windows",
		),
		"search-path":      carapace.ActionDirectories(),
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
