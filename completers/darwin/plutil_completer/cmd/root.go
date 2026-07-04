package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "plutil",
	Short: "property list utility",
	Long:  "https://keith.github.io/xcode-manpages/plutil.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("convert", "convert", "", "Convert the named file to the indicated format")
	rootCmd.Flags().StringP("extract", "extract", "", "Extract the value at the given keypath from the property list")
	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")
	rootCmd.Flags().StringP("insert", "insert", "", "Insert a value into the property list at the given keypath")
	rootCmd.Flags().StringP("key", "key", "", "Set the key for the operation")
	rootCmd.Flags().BoolP("lint", "lint", false, "Verify that the named property list file is valid")
	rootCmd.Flags().StringP("pkey", "p", "", "Print the value at the given keypath from the property list")
	rootCmd.Flags().BoolP("raw", "raw", false, "Use raw output for print/extract")
	rootCmd.Flags().BoolP("reformat", "reformat", false, "Reformat the named property list file")
	rootCmd.Flags().StringP("remove", "remove", "", "Remove a value at the given keypath from the property list")
	rootCmd.Flags().StringP("replace", "replace", "", "Replace a value in the property list at the given keypath")
	rootCmd.Flags().StringP("type", "type", "", "Set the type for the operation")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"convert": carapace.ActionValues("binary1", "xml1", "json", "objc", "swift").StyleF(style.ForKeyword),
		"type":    carapace.ActionValues("string", "array", "dict", "bool", "integer", "float", "date").StyleF(style.ForKeyword),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
