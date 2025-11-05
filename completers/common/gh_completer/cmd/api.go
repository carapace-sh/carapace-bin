package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net/http"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var apiCmd = &cobra.Command{
	Use:   "api <endpoint>",
	Short: "Make an authenticated GitHub API request",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apiCmd).Standalone()

	apiCmd.Flags().String("cache", "", "Cache the response, e.g. \"3600s\", \"60m\", \"1h\"")
	apiCmd.Flags().StringSliceP("field", "F", nil, "Add a typed parameter in `key=value` format (use \"@<path>\" or \"@-\" to read value from file or stdin)")
	apiCmd.Flags().StringSliceP("header", "H", nil, "Add a HTTP request header in `key:value` format")
	apiCmd.Flags().String("hostname", "", "The GitHub hostname for the request (default \"github.com\")")
	apiCmd.Flags().BoolP("include", "i", false, "Include HTTP response status line and headers in the output")
	apiCmd.Flags().String("input", "", "The `file` to use as body for the HTTP request (use \"-\" to read from standard input)")
	apiCmd.Flags().StringP("jq", "q", "", "Query to select values from the response using jq syntax")
	apiCmd.Flags().StringP("method", "X", "", "The HTTP method for the request")
	apiCmd.Flags().Bool("paginate", false, "Make additional HTTP requests to fetch all pages of results")
	apiCmd.Flags().StringSliceP("preview", "p", nil, "Opt into GitHub API previews (names should omit '-preview')")
	apiCmd.Flags().StringSliceP("raw-field", "f", nil, "Add a string parameter in `key=value` format")
	apiCmd.Flags().Bool("silent", false, "Do not print the response body")
	apiCmd.Flags().Bool("slurp", false, "Use with \"--paginate\" to return an array of all pages of either JSON arrays or objects")
	apiCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	apiCmd.Flags().Bool("verbose", false, "Include full HTTP request and response in the output")
	rootCmd.AddCommand(apiCmd)

	carapace.Gen(apiCmd).FlagCompletion(carapace.ActionMap{
		"field": carapace.ActionMultiPartsN("=", 2, func(c carapace.Context) carapace.Action {
			if len(c.Parts) == 0 {
				return carapace.ActionValues()
			}
			return carapace.ActionFiles().Prefix("@")
		}),
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

	carapace.Gen(apiCmd).PreInvoke(func(cmd *cobra.Command, _ *pflag.Flag, action carapace.Action) carapace.Action {
		if f := cmd.Flag("hostname"); f.Changed {
			return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
				c.Setenv("GH_HOST", f.Value.String()) // TODO provide Action.SetEnv or similar in carapace
				return action.Invoke(c).ToA()
			})
		}
		return action
	})
}
