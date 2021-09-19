package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/packer_completer/cmd/action"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build image(s) from template",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCmd).Standalone()

	buildCmd.Flags().Bool("color", false, "Enable color output")
	buildCmd.Flags().Bool("debug", false, "Debug mode enabled for builds.")
	buildCmd.Flags().String("except", "", "Run all builds and post-processors other than these.")
	buildCmd.Flags().Bool("force", false, "Force a build to continue if artifacts exist, deletes existing artifacts.")
	buildCmd.Flags().Bool("machine-readable", false, "Produce machine-readable output.")
	buildCmd.Flags().String("on-error", "", "If the build fails do: clean up (default), abort, ask, or run-cleanup-provisioner.")
	buildCmd.Flags().String("only", "", "Build only the specified builds.")
	buildCmd.Flags().String("parallel-builds", "", "Number of builds to run in parallel. 1 disables parallelization. 0 means no limit (Default: 0)")
	buildCmd.Flags().Bool("timestamp-ui", false, "Enable prefixing of each ui output with an RFC3339 timestamp.")
	buildCmd.Flags().StringArray("var", []string{}, "Variable for templates, can be used multiple times.")
	buildCmd.Flags().String("var-file", "", "JSON or HCL2 file containing user variables.")
	rootCmd.AddCommand(buildCmd)

	carapace.Gen(buildCmd).FlagCompletion(carapace.ActionMap{
		// TODO except/only
		"on-error": carapace.ActionValues("cleanup", "abort", "ask", "run-cleanup-provisioner"),
		"var": carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			if len(c.Parts) == 0 && len(c.Args) > 0 {
				return action.ActionVariables(c.Args[0]).Invoke(c).Suffix("=").ToA()
			}
			return carapace.ActionValues()
		}),
		"var-file": carapace.ActionFiles(".json", ".hcl"),
	})

	carapace.Gen(buildCmd).PositionalCompletion(
		carapace.ActionFiles(".json", ".hcl"),
	)
}
