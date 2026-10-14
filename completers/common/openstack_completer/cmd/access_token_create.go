package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var access_token_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an access token",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(access_token_createCmd).Standalone()

	access_token_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	access_token_createCmd.Flags().String("consumer-key", "", "Consumer key (required)")
	access_token_createCmd.Flags().String("consumer-secret", "", "Consumer secret (required)")
	access_token_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	access_token_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	access_token_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	access_token_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	access_token_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	access_token_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	access_token_createCmd.Flags().String("request-key", "", "Request token to exchange for access token (required)")
	access_token_createCmd.Flags().String("request-secret", "", "Secret associated with <request-key> (required)")
	access_token_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	access_token_createCmd.Flags().String("verifier", "", "Verifier associated with <request-key> (required)")
	access_token_createCmd.MarkFlagRequired("consumer-key")
	access_token_createCmd.MarkFlagRequired("consumer-secret")
	access_token_createCmd.MarkFlagRequired("request-key")
	access_token_createCmd.MarkFlagRequired("request-secret")
	access_token_createCmd.MarkFlagRequired("verifier")
	access_tokenCmd.AddCommand(access_token_createCmd)
}
