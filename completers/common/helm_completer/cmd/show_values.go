package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/helm"
	"github.com/carapace-sh/carapace/pkg/condition"
	"github.com/spf13/cobra"
)

var show_valuesCmd = &cobra.Command{
	Use:   "values",
	Short: "show the chart's values",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(show_valuesCmd).Standalone()
	show_valuesCmd.Flags().String("ca-file", "", "verify certificates of HTTPS-enabled servers using this CA bundle")
	show_valuesCmd.Flags().String("cert-file", "", "identify HTTPS client using this SSL certificate file")
	show_valuesCmd.Flags().Bool("devel", false, "use development versions, too. Equivalent to version '>0.0.0-0'. If --version is set, this is ignored")
	show_valuesCmd.Flags().Bool("insecure-skip-tls-verify", false, "skip tls certificate checks for the chart download")
	show_valuesCmd.Flags().String("jsonpath", "", "supply a JSONPath expression to filter the output")
	show_valuesCmd.Flags().String("key-file", "", "identify HTTPS client using this SSL key file")
	show_valuesCmd.Flags().String("keyring", "/home/user/.gnupg/pubring.gpg", "location of public keys used for verification")
	show_valuesCmd.Flags().Bool("pass-credentials", false, "pass credentials to all domains")
	show_valuesCmd.Flags().String("password", "", "chart repository password where to locate the requested chart")
	show_valuesCmd.Flags().String("repo", "", "chart repository url where to locate the requested chart")
	show_valuesCmd.Flags().String("username", "", "chart repository username where to locate the requested chart")
	show_valuesCmd.Flags().Bool("verify", false, "verify the package before using it")
	show_valuesCmd.Flags().String("version", "", "specify a version constraint for the chart version to use. This constraint can be a specific tag (e.g. 1.1.1) or it may reference a valid range (e.g. ^2.0.0). If this is not specified, the latest version is used")
	showCmd.AddCommand(show_valuesCmd)

	carapace.Gen(show_valuesCmd).FlagCompletion(carapace.ActionMap{
		"ca-file":   carapace.ActionFiles(),
		"cert-file": carapace.ActionFiles(),
		"key-file":  carapace.ActionFiles(),
		"keyring":   carapace.ActionFiles(),
		"version": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				if splitted := strings.Split(c.Args[0], "/"); len(splitted) == 2 {
					return helm.ActionChartVersions(helm.ChartVersionOpts{Repo: splitted[0], Chart: splitted[1]})
				}
			}
			return carapace.ActionValues()
		}),
	})

	carapace.Gen(show_valuesCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionFiles(),
			helm.ActionRepositoryCharts().UnlessF(condition.CompletingPath),
		).ToA(),
	)
}
