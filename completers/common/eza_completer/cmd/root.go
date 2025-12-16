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

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("absolute", false, "display entries with their absolute path")
	rootCmd.Flags().BoolP("accessed", "u", false, "use the accessed timestamp field")
	rootCmd.Flags().BoolP("across", "x", false, "sort the grid across, rather than downwards")
	rootCmd.Flags().CountP("all", "a", "show hidden and 'dot' files")
	rootCmd.Flags().CountP("almost-all", "A", "equivalent to --all")
	rootCmd.Flags().BoolP("binary", "b", false, "list file sizes with binary prefixes")
	rootCmd.Flags().BoolP("blocksize", "S", false, "show size of allocated file system blocks")
	rootCmd.Flags().BoolP("bytes", "B", false, "list file sizes in bytes, without any prefixes")
	rootCmd.Flags().Bool("changed", false, "use the changed timestamp field")
	rootCmd.Flags().StringP("classify", "F", "", "display type indicator by file names")
	rootCmd.Flags().String("color", "", "when to use terminal colours")
	rootCmd.Flags().String("color-scale", "", "highlight levels of 'field' distinctly")
	rootCmd.Flags().String("color-scale-mode", "", "use gradient or fixed colors in --color-scale")
	rootCmd.Flags().String("colour", "", "when to use terminal colours")
	rootCmd.Flags().String("colour-scale", "", "highlight levels of 'field' distinctly")
	rootCmd.Flags().String("colour-scale-mode", "", "use gradient or fixed colors in --color-scale")
	rootCmd.Flags().BoolP("context", "Z", false, "list each file's security context")
	rootCmd.Flags().BoolP("created", "U", false, "use the created timestamp field")
	rootCmd.Flags().BoolP("dereference", "X", false, "dereference symbolic links when displaying information")
	rootCmd.Flags().BoolP("extended", "@", false, "list each file's extended attributes and sizes")
	rootCmd.Flags().BoolP("flags", "O", false, "list file flags")
	rootCmd.Flags().Bool("git", false, "list each file's Git status, if tracked or ignored")
	rootCmd.Flags().Bool("git-ignore", false, "ignore files mentioned in '.gitignore'")
	rootCmd.Flags().Bool("git-repos", false, "list root of git-tree status")
	rootCmd.Flags().BoolP("grid", "G", false, "display entries as a grid")
	rootCmd.Flags().BoolP("group", "g", false, "list each file's group")
	rootCmd.Flags().Bool("group-directories-first", false, "list directories before other files")
	rootCmd.Flags().BoolP("header", "h", false, "add a header row to each column")
	rootCmd.Flags().Bool("help", false, "show list of command-line options")
	rootCmd.Flags().Bool("hyperlink", false, "display entries as hyperlinks")
	rootCmd.Flags().String("icons", "", "when to display icons")
	rootCmd.Flags().StringP("ignore-glob", "I", "", "glob patterns (pipe-separated) of files to ignore")
	rootCmd.Flags().BoolP("inode", "i", false, "list each file's inode number")
	rootCmd.Flags().StringP("level", "L", "", "limit the depth of recursion")
	rootCmd.Flags().BoolP("links", "H", false, "list each file's number of hard links")
	rootCmd.Flags().BoolP("list-dirs", "d", false, "list directories as files; don't list their contents")
	rootCmd.Flags().BoolP("long", "l", false, "display extended file metadata as a table")
	rootCmd.Flags().BoolP("modified", "m", false, "use the modified timestamp field")
	rootCmd.Flags().BoolP("mounts", "M", false, "show mount details")
	rootCmd.Flags().Bool("no-filesize", false, "suppress the filesize field")
	rootCmd.Flags().Bool("no-git", false, "suppress Git status")
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
	rootCmd.Flags().Bool("smart-group", false, "only show group if it has a different name from owner")
	rootCmd.Flags().StringP("sort", "s", "", "which field to sort by")
	rootCmd.Flags().Bool("stdin", false, "read file names from stdin, one per line or other separator ")
	rootCmd.Flags().StringP("time", "t", "", "which timestamp field to list")
	rootCmd.Flags().String("time-style", "", "how to format timestamps")
	rootCmd.Flags().Bool("total-size", false, "show the size of a directory as the size of all")
	rootCmd.Flags().BoolP("tree", "T", false, "recurse into directories as a tree")
	rootCmd.Flags().BoolP("version", "v", false, "show version of eza")
	rootCmd.Flags().StringP("width", "w", "", "set screen width in columns")

	rootCmd.MarkFlagsMutuallyExclusive("color", "colour")
	rootCmd.MarkFlagsMutuallyExclusive("color-scale", "colour-scale")
	rootCmd.MarkFlagsMutuallyExclusive("color-scale-mode", "colour-scale-mode")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"classify":          carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"color":             carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"color-scale":       carapace.ActionValues("all", "age", "size"),
		"color-scale-mode":  carapace.ActionValues("fixed", "gradient"),
		"colour":            carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"colour-scale":      carapace.ActionValues("all", "age", "size"),
		"colour-scale-mode": carapace.ActionValues("fixed", "gradient"),
		"icons":             carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"sort":              carapace.ActionValues("name", "Name", "size", "extension", "Extension", "modified", "changed", "accessed", "created", "inode", "type", "none"),
		"time":              carapace.ActionValues("modified", "accessed", "created"),
		"time-style":        carapace.ActionValues("default", "iso", "long-iso", "full-iso"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
