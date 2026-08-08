package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var extractCmd = &cobra.Command{
	Use:    "extract",
	Short:  "extract files from archive",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(extractCmd).Standalone()

	extractCmd.Flags().String("all", "", "extract all saved versions in numbered directories")
	extractCmd.Flags().BoolP("force", "f", false, "overwrite existing output files")
	extractCmd.Flags().String("index", "", "create index for archive without extracting")
	extractCmd.Flags().String("key", "", "decrypt with password")
	extractCmd.Flags().Bool("noattributes", false, "ignore saved attributes and extract with defaults")
	extractCmd.Flags().String("not", "", "do not extract files matching the given pattern")
	extractCmd.Flags().String("only", "", "only extract files matching the given pattern")
	extractCmd.Flags().String("repack", "", "store extracted files in new archive")
	extractCmd.Flags().StringP("summary", "s", "", "show only percent completed")
	extractCmd.Flags().Bool("test", false, "do not write to disk")
	extractCmd.Flags().StringP("threads", "t", "", "extract at most N blocks in parallel")
	extractCmd.Flags().String("to", "", "rename internal files to external names")
	extractCmd.Flags().String("until", "", "ignore part of archive updated after date or version")

	carapace.Gen(extractCmd).FlagCompletion(carapace.ActionMap{
		"all":     carapace.ActionValues("2", "3", "4"),
		"index":   carapace.ActionFiles(),
		"key":     carapace.ActionValues(),
		"not":     carapace.ActionFiles(),
		"only":    carapace.ActionFiles(),
		"repack":  carapace.ActionFiles(),
		"summary": carapace.ActionValues(),
		"threads": carapace.ActionValues("0", "1", "2", "3", "4", "5", "6", "7", "8"),
		"to":      carapace.ActionFiles(),
		"until":   carapace.ActionValues(),
	})

	carapace.Gen(extractCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)

	carapace.Gen(extractCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
