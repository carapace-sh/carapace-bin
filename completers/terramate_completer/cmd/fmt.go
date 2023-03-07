package cmd

import (
	"github.com/rsteube/carapace"
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
	rootCmd.AddCommand(fmtCmd)
}
