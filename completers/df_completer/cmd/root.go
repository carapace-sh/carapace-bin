package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/df_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "df",
	Short: "report file system disk space usage",
	Long:  "https://man7.org/linux/man-pages/man1/df.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all", "a", false, "include pseudo, duplicate, inaccessible file systems")
	rootCmd.Flags().StringP("block-size", "B", "", "scale sizes by SIZE before printing them")
	rootCmd.Flags().StringP("exclude-type", "x", "", "limit listing to file systems not of type TYPE")
	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().BoolP("human-readable", "h", false, "print sizes in powers of 1024 (e.g., 1023M)")
	rootCmd.Flags().BoolP("inodes", "i", false, "list inode information instead of block usage")
	rootCmd.Flags().BoolS("k", "k", false, "like --block-size=1K")
	rootCmd.Flags().BoolP("local", "l", false, "limit listing to local file systems")
	rootCmd.Flags().Bool("no-sync", false, "do not invoke sync before getting usage info (default)")
	rootCmd.Flags().String("output", "", "use the output format defined by FIELD_LIST, or print all fields if FIELD_LIST is omitted.")
	rootCmd.Flags().BoolP("portability", "P", false, "use the POSIX output format")
	rootCmd.Flags().BoolP("print-type", "T", false, "print file system type")
	rootCmd.Flags().BoolP("si", "H", false, "print sizes in powers of 1000 (e.g., 1.1G)")
	rootCmd.Flags().Bool("sync", false, "invoke sync before getting usage info")
	rootCmd.Flags().Bool("total", false, "elide all entries insignificant to available space, and produce a grand total")
	rootCmd.Flags().StringP("type", "t", "", "limit listing to file systems of type TYPE")
	rootCmd.Flags().BoolS("v", "v", false, "(ignored)")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	rootCmd.Flag("output").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"exclude-type": action.ActionFilesystemTypes(),
		"output":       action.ActionColumns().UniqueList(","),
		"type":         action.ActionFilesystemTypes(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
