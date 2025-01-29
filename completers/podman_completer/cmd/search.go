package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search [options] TERM",
	Short: "Search registry for image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	searchCmd.Flags().String("authfile", "", "Path of the authentication file. Use REGISTRY_AUTH_FILE environment variable to override")
	searchCmd.Flags().String("cert-dir", "", "`Pathname` of a directory containing TLS certificates and keys")
	searchCmd.Flags().Bool("compatible", false, "List stars, official and automated columns (Docker compatibility)")
	searchCmd.Flags().String("creds", "", "`Credentials` (USERNAME:PASSWORD) to use for authenticating to a registry")
	searchCmd.Flags().StringSliceP("filter", "f", []string{}, "Filter output based on conditions provided (default [])")
	searchCmd.Flags().String("format", "", "Change the output format to JSON or a Go template")
	searchCmd.Flags().String("limit", "", "Limit the number of results")
	searchCmd.Flags().Bool("list-tags", false, "List the tags of the input registry")
	searchCmd.Flags().Bool("no-trunc", false, "Do not truncate the output")
	searchCmd.Flags().Bool("tls-verify", false, "Require HTTPS and verify certificates when contacting registries")
	rootCmd.AddCommand(searchCmd)
}
