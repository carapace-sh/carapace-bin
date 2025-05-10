package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golangcilint"
	"github.com/spf13/cobra"
)

var fmtCmd = &cobra.Command{
	Use:   "fmt",
	Short: "Format Go source files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fmtCmd).Standalone()

	fmtCmd.Flags().StringP("config", "c", "", "Read config from file path `PATH`")
	fmtCmd.Flags().BoolP("diff", "d", false, "Display diffs instead of rewriting files")
	fmtCmd.Flags().Bool("diff-colored", false, "Display diffs instead of rewriting files (with colors)")
	fmtCmd.Flags().StringSliceP("enable", "E", nil, "Enable specific formatter")
	fmtCmd.Flags().Bool("no-config", false, "Don't read config file")
	fmtCmd.Flags().Bool("stdin", false, "Use standard input for piping source files")
	rootCmd.AddCommand(fmtCmd)

	carapace.Gen(fmtCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
		"enable": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return golangcilint.ActionFormatters(fmtCmd.Flag("config").Value.String()).UniqueList(",")
		}),
	})
}
