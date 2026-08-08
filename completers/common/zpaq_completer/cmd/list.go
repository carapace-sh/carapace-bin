package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:    "list",
	Short:  "list the archive contents",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().String("all", "", "list all saved versions")
	listCmd.Flags().BoolP("force", "f", false, "compare files by computing SHA-1 hashes")
	listCmd.Flags().String("key", "", "decrypt with password")
	listCmd.Flags().Bool("noattributes", false, "do not list or compare attributes")
	listCmd.Flags().String("not", "", "do not list files matching the given pattern")
	listCmd.Flags().String("only", "", "only list files matching the given pattern")
	listCmd.Flags().StringP("summary", "s", "", "sort by decreasing size and show N largest files")
	listCmd.Flags().StringP("threads", "t", "", "use N threads")
	listCmd.Flags().String("to", "", "rename external files to respective internal names")
	listCmd.Flags().String("until", "", "ignore part of archive updated after date or version")

	carapace.Gen(listCmd).FlagCompletion(carapace.ActionMap{
		"all":     carapace.ActionValues("2", "3", "4"),
		"key":     carapace.ActionValues(),
		"not":     carapace.ActionFiles(),
		"only":    carapace.ActionFiles(),
		"summary": carapace.ActionValues(),
		"threads": carapace.ActionValues("0", "1", "2", "3", "4", "5", "6", "7", "8"),
		"to":      carapace.ActionFiles(),
		"until":   carapace.ActionValues(),
	})

	carapace.Gen(listCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)

	carapace.Gen(listCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
