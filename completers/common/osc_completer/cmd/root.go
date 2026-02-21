package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "osc",
	Short: "OpenStack command line interface",
	Long:  "https://github.com/gtema/openstack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	// Connection options
	rootCmd.PersistentFlags().String("os-cloud", "", "Name reference to the clouds.yaml entry for the cloud configuration")
	rootCmd.PersistentFlags().Bool("cloud-config-from-env", false, "Get the cloud config from environment variables")
	rootCmd.PersistentFlags().String("os-cloud-name", "", "Cloud name used when configuration is retrieved from environment variables")
	rootCmd.PersistentFlags().String("os-project-id", "", "Project ID to use instead of the one in connection profile")
	rootCmd.PersistentFlags().String("os-project-name", "", "Project Name to use instead of the one in the connection profile")
	rootCmd.PersistentFlags().String("os-region-name", "", "Region Name to use instead of the one in the connection profile")
	rootCmd.PersistentFlags().String("os-client-config-file", "", "Custom path to the clouds.yaml config file")
	rootCmd.PersistentFlags().String("os-client-secure-file", "", "Custom path to the secure.yaml config file")
	rootCmd.PersistentFlags().String("auth-helper-cmd", "", "External authentication helper command")

	// Output options
	rootCmd.PersistentFlags().StringP("output", "o", "", "Output format")
	rootCmd.PersistentFlags().StringArrayP("fields", "f", nil, "Fields to return in the output")
	rootCmd.PersistentFlags().BoolP("pretty", "p", false, "Pretty print the output")
	rootCmd.PersistentFlags().CountP("verbose", "v", "Verbosity level")
	rootCmd.PersistentFlags().String("table-arrangement", "dynamic", "Output Table arrangement")
	rootCmd.PersistentFlags().Bool("timing", false, "Record HTTP request timings")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"output":            carapace.ActionValues("json", "wide"),
		"table-arrangement": carapace.ActionValues("dynamic", "dynamic-full-width", "disabled"),
		"os-client-config-file": carapace.ActionFiles(),
		"os-client-secure-file": carapace.ActionFiles(),
	})
}
