package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "path",
	Short: "Manipulate and check paths",
	Long:  "https://fishshell.com/docs/current/cmds/path.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("all", false, "return 0 if all paths pass")
	rootCmd.Flags().BoolS("C", "C", false, "center padding")
	rootCmd.Flags().Bool("center", false, "center padding")
	rootCmd.Flags().BoolS("E", "E", false, "remove extension")
	rootCmd.Flags().Bool("no-extension", false, "remove extension")
	rootCmd.Flags().BoolS("R", "R", false, "relative mtime")
	rootCmd.Flags().Bool("relative", false, "relative mtime")
	rootCmd.Flags().BoolS("h", "h", false, "display help")
	rootCmd.Flags().String("key", "", "sort key")
	rootCmd.Flags().BoolS("q", "q", false, "suppress output")
	rootCmd.Flags().Bool("quiet", false, "suppress output")
	rootCmd.Flags().BoolS("r", "r", false, "reverse sort")
	rootCmd.Flags().Bool("reverse", false, "reverse sort")
	rootCmd.Flags().BoolS("u", "u", false, "deduplicate")
	rootCmd.Flags().Bool("unique", false, "deduplicate")
	rootCmd.Flags().BoolS("v", "v", false, "invert filter")
	rootCmd.Flags().Bool("invert", false, "invert filter")
	rootCmd.Flags().BoolS("z", "z", false, "NUL-delimited input")
	rootCmd.Flags().Bool("null-in", false, "NUL-delimited input")
	rootCmd.Flags().BoolS("Z", "Z", false, "NUL-delimited output")
	rootCmd.Flags().Bool("null-out", false, "NUL-delimited output")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"key": carapace.ActionValues("basename", "dirname", "path"),
	})
}
