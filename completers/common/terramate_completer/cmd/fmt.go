package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fmtCmd = &cobra.Command{
	Use:   "fmt",
	Short: "Format all files inside dir recursively",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fmtCmd).Standalone()

	fmtCmd.Flags().Bool("check", false, "Lists unformatted files, exit with 0 if all is formatted, 1 otherwise")
	fmtCmd.Flags().Bool("detailed-exit-code", false, "Return an appropriate exit code (0 = ok, 1 = error, 2 = no error but changes were made)")
	rootCmd.AddCommand(fmtCmd)
}
