package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "eza",
	Short: "a modern replacement for ls",
	Long:  "https://eza.rocks/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("accessed", "u", false, "use the accessed timestamp field")
	rootCmd.Flags().BoolP("across", "x", false, "sort the grid across, rather than downwards")
	rootCmd.Flags().CountP("all", "a", "show hidden and 'dot' files")
	rootCmd.Flags().BoolP("binary", "b", false, "list file sizes with binary prefixes")
	rootCmd.Flags().BoolP("blocksize", "S", false, "show size of allocated file system blocks")
	rootCmd.Flags().BoolP("bytes", "B", false, "list file sizes in bytes, without any prefixes")
	rootCmd.Flags().Bool("changed", false, "use the changed timestamp field")
	rootCmd.Flags().BoolP("classify", "F", false, "display type indicator by file names")
	rootCmd.Flags().String("color", "", "when to use terminal colours (always, auto, never)")
	rootCmd.Flags().Bool("color-scale", false, "highlight levels of file sizes distinctly")
	rootCmd.Flags().String("colour", "", "when to use terminal colours (always, auto, never)")
	rootCmd.Flags().Bool("colour-scale", false, "highlight levels of file sizes distinctly")
	rootCmd.Flags().BoolP("context", "Z", false, "list each file's security context")
	rootCmd.Flags().BoolP("created", "U", false, "use the created timestamp field")
	rootCmd.Flags().BoolP("extended", "@", false, "list each file's extended attributes and sizes")
	rootCmd.Flags().Bool("git", false, "list each file's Git status, if tracked or ignored")
	rootCmd.Flags().Bool("git-ignore", false, "ignore files mentioned in '.gitignore'")
	rootCmd.Flags().Bool("git-repos", false, "list root of git-tree status")
	rootCmd.Flags().BoolP("grid", "G", false, "display entries as a grid (default)")
	rootCmd.Flags().BoolP("group", "g", false, "list each file's group")
	rootCmd.Flags().Bool("group-directories-first", false, "list directories before other files")
	rootCmd.Flags().BoolP("header", "h", false, "add a header row to each column")
	rootCmd.Flags().BoolP("help", "?", false, "show list of command-line options")
	rootCmd.Flags().Bool("hyperlink", false, "display entries as hyperlinks")
	rootCmd.Flags().Bool("icons", false, "display icons")
	rootCmd.Flags().StringP("ignore-glob", "I", "", "glob patterns (pipe-separated) of files to ignore")
	rootCmd.Flags().BoolP("inode", "i", false, "list each file's inode number")
	rootCmd.Flags().StringP("level", "L", "", "limit the depth of recursion")
	rootCmd.Flags().BoolP("links", "H", false, "list each file's number of hard links")
	rootCmd.Flags().BoolP("list-dirs", "d", false, "list directories as files; don't list their contents")
	rootCmd.Flags().BoolP("long", "l", false, "display extended file metadata as a table")
	rootCmd.Flags().BoolP("modified", "m", false, "use the modified timestamp field")
	rootCmd.Flags().BoolP("mounts", "M", false, "show mount details (Linux and MacOS only)")
	rootCmd.Flags().Bool("no-filesize", false, "suppress the filesize field")
	rootCmd.Flags().Bool("no-git", false, "suppress Git status (always overrides --git, --git-repos, --git-repos-no-status)")
	rootCmd.Flags().Bool("no-icons", false, "don't display icons (always overrides --icons)")
	rootCmd.Flags().Bool("no-permissions", false, "suppress the permissions field")
	rootCmd.Flags().Bool("no-quotes", false, "don't quote file names with spaces")
	rootCmd.Flags().Bool("no-time", false, "suppress the time field")
	rootCmd.Flags().Bool("no-user", false, "suppress the user field")
	rootCmd.Flags().BoolP("numeric", "n", false, "list numeric user and group IDs")
	rootCmd.Flags().BoolP("octal-permissions", "o", false, "list each file's permission in octal format")
	rootCmd.Flags().BoolP("oneline", "1", false, "display one entry per line")
	rootCmd.Flags().BoolP("only-dirs", "D", false, "list only directories")
	rootCmd.Flags().BoolP("only-files", "f", false, "list only files")
	rootCmd.Flags().BoolP("recurse", "R", false, "recurse into directories")
	rootCmd.Flags().BoolP("reverse", "r", false, "reverse the sort order")
	rootCmd.Flags().StringP("sort", "s", "", "which field to sort by")
	rootCmd.Flags().StringP("time", "t", "", "which timestamp field to list (modified, accessed, created)")
	rootCmd.Flags().String("time-style", "", "how to format timestamps (default, iso, long-iso, full-iso)")
	rootCmd.Flags().BoolP("tree", "T", false, "recurse into directories as a tree")
	rootCmd.Flags().BoolP("version", "v", false, "show version of eza")
	rootCmd.Flags().StringP("width", "w", "", "set screen width in columns")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"color":      carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"colour":     carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"sort":       carapace.ActionValues("name", "Name", "size", "extension", "Extension", "modified", "changed", "accessed", "created", "inode", "type", "none"),
		"time":       carapace.ActionValues("modified", "accessed", "created"),
		"time-style": carapace.ActionValues("default", "iso", "long-iso", "full-iso"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
