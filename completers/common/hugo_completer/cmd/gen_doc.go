package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gen_docCmd = &cobra.Command{
	Use:   "doc",
	Short: "Generate Markdown documentation for the Hugo CLI.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gen_docCmd).Standalone()
	gen_docCmd.PersistentFlags().String("dir", "/tmp/hugodoc/", "the directory to write the doc.")
	genCmd.AddCommand(gen_docCmd)

	carapace.Gen(gen_docCmd).FlagCompletion(carapace.ActionMap{
		"dir": carapace.ActionDirectories(),
	})
}
