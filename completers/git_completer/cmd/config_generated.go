package cmd

import (
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Get and set repository or global options",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	configCmd.Flags().Bool("add", false, "add a new variable: name value")
	configCmd.Flags().String("blob", "", "read config from given blob object")
	configCmd.Flags().Bool("bool-or-int", false, "value is --bool or --int")
	configCmd.Flags().Bool("bool", false, "value is \"true\" or \"false\"")
	configCmd.Flags().String("default", "", "with --get, use default value when missing entry")
	configCmd.Flags().BoolP("edit", "e", false, "open an editor")
	configCmd.Flags().Bool("expiry-date", false, "value is an expiry date")
	configCmd.Flags().BoolP("file", "f", false, "<file>     use given config file")
	configCmd.Flags().Bool("get-all", false, "get all values: key [value-regex]")
	configCmd.Flags().Bool("get-colorbool", false, "find the color setting: slot [stdout-is-tty]")
	configCmd.Flags().Bool("get-color", false, "find the color configured: slot [default]")
	configCmd.Flags().Bool("get", false, "get value: name [value-regex]")
	configCmd.Flags().Bool("get-regexp", false, "get values for regexp: name-regex [value-regex]")
	configCmd.Flags().Bool("get-urlmatch", false, "get value specific for the URL: section[.var] URL")
	configCmd.Flags().Bool("global", false, "use global config file")
	configCmd.Flags().Bool("includes", false, "respect include directives on lookup")
	configCmd.Flags().Bool("int", false, "value is decimal number")
	configCmd.Flags().BoolP("list", "l", false, "list all")
	configCmd.Flags().Bool("local", false, "use repository config file")
	configCmd.Flags().Bool("name-only", false, "show variable names only")
	configCmd.Flags().Bool("path", false, "value is a path (file or directory name)")
	configCmd.Flags().Bool("remove-section", false, "remove a section: name")
	configCmd.Flags().Bool("rename-section", false, "rename section: old-name new-name")
	configCmd.Flags().Bool("replace-all", false, "replace all matching variables: name value [value_regex]")
	configCmd.Flags().Bool("show-origin", false, "show origin of config (file, standard input, blob, command line)")
	configCmd.Flags().Bool("show-scope", false, "show scope of config (worktree, local, global, system, command)")
	configCmd.Flags().Bool("system", false, "use system config file")
	configCmd.Flags().BoolP("type", "t", false, "<>         value is given this type")
	configCmd.Flags().Bool("unset-all", false, "remove all matches: name [value-regex]")
	configCmd.Flags().Bool("unset", false, "remove a variable: name [value-regex]")
	configCmd.Flags().Bool("worktree", false, "use per-worktree config file")
	configCmd.Flags().BoolP("null", "z", false, "terminate values with NUL byte")
	rootCmd.AddCommand(configCmd)
}
