package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zpaq",
	Short: "Incremental Journaling Backup Utility and Archiver",
	Long:  "https://mattmahoney.net/dc/zpaq.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().BoolS("f", "f", false, "-force")
	rootCmd.PersistentFlags().BoolS("force", "force", false, "")
	rootCmd.PersistentFlags().StringS("key", "key", "", "access encrypted archive with password X")
	rootCmd.PersistentFlags().BoolS("noattributes", "noattributes", false, "ignore file attributes or permissions")
	rootCmd.PersistentFlags().StringSliceS("not", "not", nil, "exclude (* and ? match any string or char)")
	rootCmd.PersistentFlags().StringSliceS("only", "only", nil, "include only matches (default: *)")
	rootCmd.PersistentFlags().BoolS("s1", "s1", false, "-summary 1")
	rootCmd.PersistentFlags().BoolS("sN", "sN", false, "-summary N")
	rootCmd.PersistentFlags().StringS("summary", "summary", "", "")
	rootCmd.PersistentFlags().BoolS("t1", "t1", false, "-threads 1")
	rootCmd.PersistentFlags().BoolS("tN", "tN", false, "-threads N")
	rootCmd.PersistentFlags().StringS("threads", "threads", "", "use N threads (default: 0 = all cores)")
	rootCmd.PersistentFlags().StringSliceS("to", "to", nil, "rename files... to P... or all to P/all")
	rootCmd.PersistentFlags().StringS("until", "until", "", "roll back archive to N'th update (-N from end) or set date N")

	rootCmd.Flag("not").Nargs = -1
	rootCmd.Flag("only").Nargs = -1
	rootCmd.Flag("to").Nargs = -1

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"not":  carapace.ActionFiles(),
		"only": carapace.ActionFiles(),
		"to":   carapace.ActionFiles(),
	})
}
