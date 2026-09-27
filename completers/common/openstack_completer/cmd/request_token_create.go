package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var request_token_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a request token",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(request_token_createCmd).Standalone()

	request_token_createCmd.Flags().StringP("column", "c", "", "specify the column(s) to include, can be repeated to show multiple columns")
	request_token_createCmd.Flags().String("consumer-key", "", "Consumer key (required)")
	request_token_createCmd.Flags().String("consumer-secret", "", "Consumer secret (required)")
	request_token_createCmd.Flags().String("domain", "", "Domain owning <project> (name or ID)")
	request_token_createCmd.Flags().Bool("fit-width", false, "Fit the table to the display width.")
	request_token_createCmd.Flags().StringP("format", "f", "", "the output format, defaults to table")
	request_token_createCmd.Flags().String("max-width", "", "Maximum display width, <1 to disable.")
	request_token_createCmd.Flags().Bool("noindent", false, "whether to disable indenting the JSON")
	request_token_createCmd.Flags().String("prefix", "", "add a prefix to all variable names")
	request_token_createCmd.Flags().Bool("print-empty", false, "Print empty table if there is no data to show.")
	request_token_createCmd.Flags().String("project", "", "Project that consumer wants to access (name or ID) (required)")
	request_token_createCmd.Flags().String("variable", "", "==SUPPRESS==")
	request_token_createCmd.MarkFlagRequired("consumer-key")
	request_token_createCmd.MarkFlagRequired("consumer-secret")
	request_token_createCmd.MarkFlagRequired("project")
	request_tokenCmd.AddCommand(request_token_createCmd)
}
