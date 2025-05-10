package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net/http"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api <endpoint>",
	Short: "Make an authenticated request to the GitLab API.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apiCmd).Standalone()

	apiCmd.Flags().StringSliceP("field", "F", nil, "Add a parameter of inferred type. Changes the default HTTP method to \"POST\".")
	apiCmd.Flags().StringSliceP("header", "H", nil, "Add an additional HTTP request header.")
	apiCmd.Flags().String("hostname", "", "The GitLab hostname for the request. Defaults to \"gitlab.com\", or the authenticated host in the current Git directory.")
	apiCmd.Flags().BoolP("include", "i", false, "Include HTTP response headers in the output.")
	apiCmd.Flags().String("input", "", "The file to use as the body for the HTTP request.")
	apiCmd.Flags().StringP("method", "X", "", "The HTTP method for the request.")
	apiCmd.Flags().Bool("paginate", false, "Make additional HTTP requests to fetch all pages of results.")
	apiCmd.Flags().StringSliceP("raw-field", "f", nil, "Add a string parameter.")
	apiCmd.Flags().Bool("silent", false, "Do not print the response body.")
	rootCmd.AddCommand(apiCmd)

	carapace.Gen(apiCmd).FlagCompletion(carapace.ActionMap{
		"header": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return http.ActionRequestHeaderNames().Invoke(c).Suffix(":").ToA()
				case 1:
					return http.ActionRequestHeaderValues(c.Parts[0])
				default:
					return carapace.ActionValues()
				}
			})
		}),
		"hostname": action.ActionConfigHosts(),
		"input":    carapace.ActionFiles(),
		"method":   http.ActionRequestMethods(),
	})

	carapace.Gen(apiCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.Batch(
				action.ActionApiPaths(apiCmd),
				carapace.ActionValues("graphql"),
			).Invoke(c).Merge().ToA()
		}),
	)
}
