package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:    "add",
	Short:  "append changes in files to archive",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addCmd).Standalone()

	addCmd.Flags().BoolP("force", "f", false, "add files even if the last-modified date has not changed")
	addCmd.Flags().String("fragment", "", "set the dedupe fragment size range")
	addCmd.Flags().String("index", "", "create archive suffix for remote backup")
	addCmd.Flags().String("key", "", "encrypt with password")
	addCmd.Flags().StringP("method", "m", "", "select a compression method")
	addCmd.Flags().Bool("noattributes", false, "do not save attributes or permissions")
	addCmd.Flags().String("not", "", "do not add files matching the given pattern")
	addCmd.Flags().String("only", "", "only add files matching the given pattern")
	addCmd.Flags().StringP("summary", "s", "", "show only percent completed")
	addCmd.Flags().StringP("threads", "t", "", "add at most N blocks in parallel")
	addCmd.Flags().String("to", "", "rename external files to respective internal names")
	addCmd.Flags().String("until", "", "ignore part of archive updated after date or version")

	carapace.Gen(addCmd).FlagCompletion(carapace.ActionMap{
		"fragment": carapace.ActionValues("0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"),
		"index":    carapace.ActionFiles(),
		"key":      carapace.ActionValues(),
		"method": carapace.ActionValuesDescribed(
			"0", "store with deduplication but no compression",
			"1", "default, recommended for backups",
			"2", "slower compression, fast decompression",
			"3", "higher compression",
			"4", "higher compression",
			"5", "highest compression",
			"x", "experimental journaling mode",
			"s", "experimental streaming mode",
		),
		"not":     carapace.ActionFiles(),
		"only":    carapace.ActionFiles(),
		"summary": carapace.ActionValues(),
		"threads": carapace.ActionValues("0", "1", "2", "3", "4", "5", "6", "7", "8"),
		"to":      carapace.ActionFiles(),
		"until":   carapace.ActionValues(),
	})

	carapace.Gen(addCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)

	carapace.Gen(addCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
