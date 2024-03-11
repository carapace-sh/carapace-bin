package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net/http"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api <endpoint>",
	Short: "Make an authenticated GitHub API request",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apiCmd).Standalone()

	apiCmd.Flags().String("cache", "", "Cache the response, e.g. \"3600s\", \"60m\", \"1h\"")
	apiCmd.Flags().StringSliceP("field", "F", []string{}, "Add a typed parameter in `key=value` format")
	apiCmd.Flags().StringSliceP("header", "H", []string{}, "Add a HTTP request header in `key:value` format")
	apiCmd.Flags().String("hostname", "", "The GitHub hostname for the request (default \"github.com\")")
	apiCmd.Flags().BoolP("include", "i", false, "Include HTTP response status line and headers in the output")
	apiCmd.Flags().String("input", "", "The `file` to use as body for the HTTP request (use \"-\" to read from standard input)")
	apiCmd.Flags().StringP("jq", "q", "", "Query to select values from the response using jq syntax")
	apiCmd.Flags().StringP("method", "X", "", "The HTTP method for the request")
	apiCmd.Flags().Bool("paginate", false, "Make additional HTTP requests to fetch all pages of results")
	apiCmd.Flags().StringSliceP("preview", "p", []string{}, "GitHub API preview `names` to request (without the \"-preview\" suffix)")
	apiCmd.Flags().StringSliceP("raw-field", "f", []string{}, "Add a string parameter in `key=value` format")
	apiCmd.Flags().Bool("silent", false, "Do not print the response body")
	apiCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	apiCmd.Flags().Bool("verbose", false, "Include full HTTP request and response in the output")
	rootCmd.AddCommand(apiCmd)

	carapace.Gen(apiCmd).FlagCompletion(carapace.ActionMap{
		"header":   http.ActionRequestHeaders(),
		"hostname": action.ActionConfigHosts(),
		"input":    carapace.ActionFiles(),
		"method":   http.ActionRequestMethods(),
		"preview":  action.ActionApiPreviews().UniqueList(","),
	})

	carapace.Gen(apiCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.Batch(
				action.ActionApiV3Paths(apiCmd),
				carapace.ActionValues("graphql"),
			).Invoke(c).Merge().ToA()
		}),
	)
}
