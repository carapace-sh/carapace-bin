package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secret_inspectCmd = &cobra.Command{
	Use:   "inspect [options] SECRET [SECRET...]",
	Short: "Inspect a secret",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secret_inspectCmd).Standalone()

	secret_inspectCmd.Flags().StringP("format", "f", "", "Format inspect output using Go template")
	secret_inspectCmd.Flags().Bool("pretty", false, "Print inspect output in human-readable format")
	secret_inspectCmd.Flags().Bool("showsecret", false, "Display the secret")
	secretCmd.AddCommand(secret_inspectCmd)
}
