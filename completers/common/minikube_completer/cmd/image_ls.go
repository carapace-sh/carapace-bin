package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var image_lsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "List images",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_lsCmd).Standalone()

	image_lsCmd.Flags().String("format", "", "Format output. One of: short|table|json|yaml")
	imageCmd.AddCommand(image_lsCmd)

	carapace.Gen(image_lsCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("short", "table", "json", "yaml"),
	})
}
