package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var token_issueCmd = &cobra.Command{
	Use:   "issue",
	Short: "Issue new token",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(token_issueCmd).Standalone()

	token_issueCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	token_issueCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	token_issueCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	token_issueCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	token_issueCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	token_issueCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	token_issueCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	token_issueCmd.Flags().String("variable", "", "==SUPPRESS==")
	tokenCmd.AddCommand(token_issueCmd)
}
