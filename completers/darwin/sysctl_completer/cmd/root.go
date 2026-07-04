package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sysctl",
	Short: "read or write kernel state variables",
	Long:  "https://keith.github.io/xcode-manpages/sysctl.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all", "a", false, "List all currently available values")
	rootCmd.Flags().BoolP("binary", "b", false, "Force output in raw, binary format")
	rootCmd.Flags().StringP("bufsize", "B", "", "Set the buffer size for reading variables")
	rootCmd.Flags().BoolP("description", "d", false, "Print the description instead of the value")
	rootCmd.Flags().BoolP("equals", "e", false, "Separate name and value with '='")
	rootCmd.Flags().StringP("file", "f", "", "Read variable names and values from file")
	rootCmd.Flags().BoolP("format", "F", false, "Print the format of the variable")
	rootCmd.Flags().BoolP("hex", "x", false, "Print hex dump of entire value")
	rootCmd.Flags().BoolP("hex-all", "X", false, "Equivalent to -x -a")
	rootCmd.Flags().BoolP("human", "h", false, "Format output for human readability")
	rootCmd.Flags().BoolP("ignore", "i", false, "Ignore unknown OIDs")
	rootCmd.Flags().BoolP("length", "l", false, "Show the length of variables along with values")
	rootCmd.Flags().BoolP("names", "N", false, "Show only variable names, not values")
	rootCmd.Flags().BoolP("no-names", "n", false, "Do not show variable names")
	rootCmd.Flags().BoolP("opaque", "o", false, "Show opaque variables")
	rootCmd.Flags().BoolP("quiet", "q", false, "Suppress some warnings")
	rootCmd.Flags().BoolP("type", "t", false, "Print the type of the variable")
	rootCmd.Flags().BoolP("writable", "W", false, "Display only writable variables")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})
}
