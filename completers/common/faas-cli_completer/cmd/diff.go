package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/faas-cli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Diff stack.yaml against deployed functions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diffCmd).Standalone()

	diffCmd.Flags().String("env", "local", "Environment diff mode: local, remote, or all")
	diffCmd.Flags().Bool("envsubst", true, "Substitute environment variables in stack.yaml file")
	diffCmd.Flags().StringP("gateway", "g", "http://127.0.0.1:8080", "Gateway URL starting with http(s)://")
	diffCmd.Flags().StringP("namespace", "n", "", "Namespace override for the diff")
	diffCmd.Flags().String("tag", "latest", "Override latest tag on function Docker image, accepts 'digest', 'sha', 'branch', or 'describe', or 'latest'")
	diffCmd.Flags().Bool("tls-no-verify", false, "Disable TLS validation")
	diffCmd.Flags().StringP("token", "k", "", "Pass a JWT token to use instead of basic auth")
	rootCmd.AddCommand(diffCmd)

	carapace.Gen(diffCmd).FlagCompletion(carapace.ActionMap{
		"env":       carapace.ActionValues("all", "local", "remote"),
		"namespace": action.ActionNamespaces(),
		"tag":       carapace.ActionValues("branch", "describe", "digest", "latest", "sha"),
	})
}
