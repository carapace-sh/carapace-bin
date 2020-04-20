package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "exa",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
    // DISPLAY OPTIONS
	rootCmd.Flags().BoolP("online", "1", false, "display one entry per line")
	rootCmd.Flags().BoolP("long", "l", false, "display extended file metadata as table")
	rootCmd.Flags().BoolP("grid", "G", false, "display entries as grid (default)")
	rootCmd.Flags().BoolP("across", "x", false, "sort the grid across, rather than downwards")
	rootCmd.Flags().BoolP("recurse", "R", false, "recurse into directories")
	rootCmd.Flags().BoolP("tree", "T", false, "recurse inte directories as tree")
	rootCmd.Flags().BoolP("classify", "F", false, "display type indicator by file names")
	rootCmd.Flags().String("color", "", "when to use terminal colours (always, auto, never)")
	rootCmd.Flags().String("colour", "", "when to use terminal colours (always, auto, never)")
	rootCmd.Flags().String("color-scale", "", "highlight levels of file sizes distinctly")
	rootCmd.Flags().String("colour-scale", "", "highlight levels of file sizes distinctly")

    // FILTERING AND SORTING OPTIONS
    rootCmd.Flags().BoolP("all", "a", false, "show hidden and 'dot' files")
    rootCmd.Flags().BoolP("list-dirs", "d", false, "list directories like regular files")
    rootCmd.Flags().Int32P("level", "L", 0, "limit the depth of recursion")
    rootCmd.Flags().BoolP("reverse", "r", false, "reverse the sort order")
    rootCmd.Flags().StringP("sort", "s", "", "which field to sort by")
    rootCmd.Flags().Bool("group-directories-first", false, "list directories before other files")
    rootCmd.Flags().BoolP("only-dirs", "D", false, "list only directories")
    rootCmd.Flags().StringP("ignore-glob", "I", "", "glob patterns (pipe-separated) of files to ignore")
    rootCmd.Flags().Bool("git-ignore", false, "Ignore files mentioned in '.gitignore'")

    // LONG VIEW OPTIONS
    rootCmd.Flags().BoolP("binary", "b", false, "list file sizes with binary prefixes")
    rootCmd.Flags().BoolP("bytes", "B", false, "list file sizes in bytes, without any prefixes")
    rootCmd.Flags().BoolP("group", "g", false, "list each file's group")
    rootCmd.Flags().BoolP("header", "h", false, "add a header row to each column")
    rootCmd.Flags().BoolP("links", "H", false, "list each file's number of hard links")
    rootCmd.Flags().BoolP("inode", "i", false, "list each file's inode number")
    rootCmd.Flags().BoolP("modified", "m", false, "use the modified timestamp field")
    rootCmd.Flags().BoolP("blocks", "S", false, "show number of file system blocks")
    rootCmd.Flags().StringP("time", "t", "", "which timestamp field to list (modified, accessed, created)")
    rootCmd.Flags().BoolP("accessed", "u", false, "use the accessed timestamp field")
    rootCmd.Flags().BoolP("created", "U", false, "use the created timestamp field")
    rootCmd.Flags().Bool("time-style", false, "use the created timestamp field")
    rootCmd.Flags().Bool("git", false, "list each file's Git status, if tracked or ignored")
    rootCmd.Flags().BoolP("extended", "@", false, "list each file's extended attributes and sizes")

    carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
      "color": carapace.ActionValues("always", "auto", "never"),
      "colour": carapace.ActionValues("always", "auto", "never"),
      "sort" : carapace.ActionValues("name", "Name", "size", "extension", "Extension", "modified", "changed", "accessed", "created", "inode", "type", "none"),
      "time": carapace.ActionValues("modified", "accessed", "created"),
    })

    carapace.Gen(rootCmd).Standalone()
}
