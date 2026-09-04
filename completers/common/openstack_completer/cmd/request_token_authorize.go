package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var request_token_authorizeCmd = &cobra.Command{
	Use:   "authorize",
	Short: "Authorize a request token",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(request_token_authorizeCmd).Standalone()

	request_token_authorizeCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	request_token_authorizeCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	request_token_authorizeCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	request_token_authorizeCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	request_token_authorizeCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	request_token_authorizeCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	request_token_authorizeCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	request_token_authorizeCmd.Flags().String("request-key", "", "Request token to authorize (ID only) (required)")
	request_token_authorizeCmd.Flags().String("role", "", "Roles to authorize (name or ID) (repeat option to set multiple values) (required)")
	request_token_authorizeCmd.Flags().String("variable", "", "==SUPPRESS==")
	request_token_authorizeCmd.MarkFlagRequired("request-key")
	request_token_authorizeCmd.MarkFlagRequired("role")
	request_tokenCmd.AddCommand(request_token_authorizeCmd)
}
