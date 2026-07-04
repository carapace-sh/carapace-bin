package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tar",
	Short: "manipulate tape archives",
	Long:  "https://keith.github.io/xcode-manpages/tar.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("J", "J", false, "Use xz compression")
	rootCmd.Flags().BoolS("O", "O", false, "Extract to stdout")
	rootCmd.Flags().BoolS("P", "P", false, "Do not strip leading /")
	rootCmd.Flags().BoolS("W", "W", false, "Verify archive")
	rootCmd.Flags().BoolS("Z", "Z", false, "Use compress compression")
	rootCmd.Flags().BoolS("c", "c", false, "Create archive")
	rootCmd.Flags().BoolS("j", "j", false, "Use bzip2 compression")
	rootCmd.Flags().BoolS("k", "k", false, "Do not overwrite existing files")
	rootCmd.Flags().BoolS("p", "p", false, "Preserve permissions")
	rootCmd.Flags().BoolS("r", "r", false, "Append files to archive")
	rootCmd.Flags().BoolS("t", "t", false, "List archive contents")
	rootCmd.Flags().BoolS("u", "u", false, "Update archive")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose")
	rootCmd.Flags().BoolS("x", "x", false, "Extract from archive")
	rootCmd.Flags().BoolS("z", "z", false, "Use gzip compression")

	rootCmd.Flags().StringS("C", "C", "", "Change to directory")
	rootCmd.Flags().StringS("f", "f", "", "Archive file")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"f": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
