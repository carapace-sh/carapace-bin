package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var extractCmd = &cobra.Command{
	Use:     "extract archive files...",
	Short:   "extract most recent versions of files",
	Aliases: []string{"x"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(extractCmd).Standalone()

	extractCmd.Flags().StringS("all", "all", "", "extract versions in N [4] digit directories")
	extractCmd.Flags().BoolS("force", "force", false, "overwrite existing output files")
	extractCmd.Flags().StringS("index", "index", "", "create index F for archive")
	extractCmd.Flags().StringS("repack", "repack", "", "extract to new archive F with key X (default: none)")
	extractCmd.Flags().StringS("summary", "summary", "", "if N > 0 show brief progress")
	extractCmd.Flags().BoolS("test", "test", false, "verify but do not write files")
	extractCmd.Flag("all").NoOptDefVal = " "
	rootCmd.AddCommand(extractCmd)
}
