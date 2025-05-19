package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/helm"
	"github.com/carapace-sh/carapace/pkg/condition"
	"github.com/spf13/cobra"
)

var show_readmeCmd = &cobra.Command{
	Use:   "readme",
	Short: "show the chart's README",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(show_readmeCmd).Standalone()
	show_readmeCmd.Flags().String("ca-file", "", "verify certificates of HTTPS-enabled servers using this CA bundle")
	show_readmeCmd.Flags().String("cert-file", "", "identify HTTPS client using this SSL certificate file")
	show_readmeCmd.Flags().Bool("devel", false, "use development versions, too. Equivalent to version '>0.0.0-0'. If --version is set, this is ignored")
	show_readmeCmd.Flags().Bool("insecure-skip-tls-verify", false, "skip tls certificate checks for the chart download")
	show_readmeCmd.Flags().String("key-file", "", "identify HTTPS client using this SSL key file")
	show_readmeCmd.Flags().String("keyring", "/home/user/.gnupg/pubring.gpg", "location of public keys used for verification")
	show_readmeCmd.Flags().Bool("pass-credentials", false, "pass credentials to all domains")
	show_readmeCmd.Flags().String("password", "", "chart repository password where to locate the requested chart")
	show_readmeCmd.Flags().String("repo", "", "chart repository url where to locate the requested chart")
	show_readmeCmd.Flags().String("username", "", "chart repository username where to locate the requested chart")
	show_readmeCmd.Flags().Bool("verify", false, "verify the package before using it")
	show_readmeCmd.Flags().String("version", "", "specify a version constraint for the chart version to use. This constraint can be a specific tag (e.g. 1.1.1) or it may reference a valid range (e.g. ^2.0.0). If this is not specified, the latest version is used")
	showCmd.AddCommand(show_readmeCmd)

	carapace.Gen(show_readmeCmd).FlagCompletion(carapace.ActionMap{
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

	carapace.Gen(show_readmeCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionFiles(),
			helm.ActionRepositoryCharts().UnlessF(condition.CompletingPath),
		).ToA(),
	)
}
