package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var secret_inspectCmd = &cobra.Command{
	Use:   "inspect [OPTIONS] SECRET [SECRET...]",
	Short: "Display detailed information on one or more secrets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secret_inspectCmd).Standalone()

	secret_inspectCmd.Flags().StringP("format", "f", "", "Format output using a custom template:")
	secret_inspectCmd.Flags().Bool("pretty", false, "Print the information in a human friendly format")
	secretCmd.AddCommand(secret_inspectCmd)

	carapace.Gen(secret_inspectCmd).PositionalAnyCompletion(docker.ActionSecrets())
}
