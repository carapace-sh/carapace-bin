package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sortCmd = &cobra.Command{
	Use:   "sort",
	Short: "Sort paths",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sortCmd).Standalone()

	sortCmd.Flags().BoolS("Z", "Z", false, "NUL-delimited output")
	sortCmd.Flags().String("key", "", "sort key")
	sortCmd.Flags().Bool("null-in", false, "NUL-delimited input")
	sortCmd.Flags().Bool("null-out", false, "NUL-delimited output")
	sortCmd.Flags().BoolS("q", "q", false, "suppress output")
	sortCmd.Flags().Bool("quiet", false, "suppress output")
	sortCmd.Flags().BoolP("reverse", "r", false, "reverse sort order")
	sortCmd.Flags().BoolP("unique", "u", false, "deduplicate")
	sortCmd.Flags().BoolS("z", "z", false, "NUL-delimited input")

	rootCmd.AddCommand(sortCmd)

	carapace.Gen(sortCmd).FlagCompletion(carapace.ActionMap{
		"key": carapace.ActionValues("basename", "dirname", "path"),
	})
}
