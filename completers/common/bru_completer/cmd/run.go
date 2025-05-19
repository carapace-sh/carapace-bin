package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a request",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().Bool("bail", false, "Stop execution after a failure of a request, test, or assertion")
	runCmd.Flags().String("cacert", "", "CA certificate to verify peer against")
	runCmd.Flags().String("env", "", "Environment variables")
	runCmd.Flags().String("env-var", "", "Overwrite a single environment variable, multiple usages possible")
	runCmd.Flags().StringP("format", "f", "", "Format of the file results")
	runCmd.Flags().BoolP("help", "h", false, "Show help")
	runCmd.Flags().Bool("insecure", false, "Allow insecure server connections")
	runCmd.Flags().StringP("output", "o", "", "Path to write file results to")
	runCmd.Flags().BoolS("r", "r", false, "Indicates a recursive run")
	runCmd.Flags().Bool("tests-only", false, "Only run requests that have a test")
	runCmd.Flags().Bool("version", false, "Show version number")
	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"cacert":  carapace.ActionFiles(),
		"env":     env.ActionNameValues(false),
		"env-var": carapace.ActionValues(), // TODO
		"format":  carapace.ActionValues("json", "junit", "html"),
		"output":  carapace.ActionFiles(),
	})

	carapace.Gen(runCmd).PositionalCompletion(
		carapace.ActionFiles(".bru"),
	)
}
