package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/lsfd_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/ps"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lsfd",
	Short: "list file descriptors",
	Long:  "https://man7.org/linux/man-pages/man1/lsfd.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("_drop-privilege", false, "(testing purpose) do setuid(1) just after starting")
	rootCmd.Flags().StringP("counter", "C", "", "define custom counter for --summary output")
	rootCmd.Flags().Bool("debug-filter", false, "dump the internal data structure of filter and exit")
	rootCmd.Flags().Bool("dump-counters", false, "dump counter definitions")
	rootCmd.Flags().StringP("filter", "Q", "", "apply display filter")
	rootCmd.Flags().BoolP("help", "h", false, "display this help")
	rootCmd.Flags().String("hyperlink", "", "print paths as terminal hyperlinks")
	rootCmd.Flags().StringP("inet", "i", "", "list only IPv4 and/or IPv6 sockets")
	rootCmd.Flags().BoolP("json", "J", false, "use JSON output format")
	rootCmd.Flags().BoolP("list-columns", "H", false, "list the available columns")
	rootCmd.Flags().BoolP("noheadings", "n", false, "don't print headings")
	rootCmd.Flags().BoolP("notruncate", "u", false, "don't truncate text in columns")
	rootCmd.Flags().StringP("output", "o", "", "output columns (see --list-columns)")
	rootCmd.Flags().StringP("pid", "p", "", "collect information only specified processes")
	rootCmd.Flags().BoolP("raw", "r", false, "use raw output format")
	rootCmd.Flags().String("summary", "", "print summary information")
	rootCmd.Flags().BoolP("threads", "l", false, "list in threads level")
	rootCmd.Flags().BoolP("version", "V", false, "display version")

	rootCmd.Flag("hyperlink").NoOptDefVal = " "
	rootCmd.Flag("inet").NoOptDefVal = " "
	rootCmd.Flag("summary").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"hyperlink": carapace.ActionValues("auto", "always", "never").StyleF(style.ForKeyword),
		"inet":      carapace.ActionValues("4", "6"),
		"output":    action.ActionColumns().UniqueList(","),
		"pid":       ps.ActionProcessIds().UniqueList(","),
		"summary":   carapace.ActionValues("only", "append", "never").StyleF(style.ForKeyword),
	})
}
