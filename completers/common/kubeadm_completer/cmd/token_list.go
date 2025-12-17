package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubeadm"
	"github.com/spf13/cobra"
)

var token_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List bootstrap tokens on the server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(token_listCmd).Standalone()

	token_listCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	token_listCmd.Flags().StringP("output", "o", "", "Output format. One of: text|json|yaml|kyaml|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-as-json|jsonpath-file.")
	token_listCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	tokenCmd.AddCommand(token_listCmd)

	carapace.Gen(token_listCmd).FlagCompletion(carapace.ActionMap{
		"output": kubeadm.ActionOutputFormats(),
	})
}
