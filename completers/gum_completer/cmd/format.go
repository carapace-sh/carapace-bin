package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var formatCmd = &cobra.Command{
	Use:   "format",
	Short: "Format a string using a template",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(formatCmd).Standalone()

	formatCmd.Flags().StringP("type", "t", "", "Format to use (markdown,template,code,emoji)")
	rootCmd.AddCommand(formatCmd)

	carapace.Gen(formatCmd).FlagCompletion(carapace.ActionMap{
		"type": carapace.ActionValues("markdown", "template", "code", "emoji"),
	})
}
