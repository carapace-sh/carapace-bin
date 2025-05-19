package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import [OPTIONS] --from <FROM> <PATH>",
	Short: "Import entries from another application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importCmd).Standalone()

	importCmd.Flags().String("from", "", "Application to import from")
	importCmd.Flags().BoolP("help", "h", false, "Print help")
	importCmd.Flags().Bool("merge", false, "Merge into existing database")
	importCmd.Flags().BoolP("version", "V", false, "Print version")
	importCmd.MarkFlagRequired("from")
	rootCmd.AddCommand(importCmd)

	carapace.Gen(importCmd).FlagCompletion(carapace.ActionMap{
		"from": carapace.ActionValues("autojump", "z"),
	})

	carapace.Gen(importCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
