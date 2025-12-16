package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/dbt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "dbt",
	Short: "An ELT tool for managing your SQL transformations and data models",
	Long:  "https://www.getdbt.com/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("cache-selected-only", false, "Pre cache database objects relevant to selected resource only")
	rootCmd.Flags().BoolP("debug", "d", false, "Display debug logging during dbt execution")
	rootCmd.Flags().BoolP("fail-fast", "x", false, "Stop execution upon a first failure")
	rootCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.Flags().String("log-format", "", "Specify the log format")
	rootCmd.Flags().Bool("no-anonymous-usage-stats", false, "Do not send anonymous usage stat to dbt Labs")
	rootCmd.Flags().Bool("no-cache-selected-only", false, "Pre cache all database objects related to the project")
	rootCmd.Flags().Bool("no-partial-parse", false, "Disallow partial parsing")
	rootCmd.Flags().Bool("no-print", false, "Suppress all {{ print() }} macro calls")
	rootCmd.Flags().Bool("no-static-parser", false, "Disables the static parser")
	rootCmd.Flags().Bool("no-use-colors", false, "Do not colorize the output DBT prints to the terminal")
	rootCmd.Flags().Bool("no-version-check", false, "Skip ensuring dbt's version matches the one specified in the dbt_project.yml")
	rootCmd.Flags().Bool("no-write-json", false, "Skip writing the manifest and run_results.json files to disk")
	rootCmd.Flags().Bool("partial-parse", false, "Allow for partial parsing by looking for and writing to a pickle file")
	rootCmd.Flags().String("printer-width", "", "Sets the width of terminal output")
	rootCmd.Flags().BoolP("quiet", "q", false, "Suppress all non-error logging to stdout")
	rootCmd.Flags().StringP("record-timing-info", "r", "", "Output low-level timing stats to the specified file")
	rootCmd.Flags().Bool("use-colors", false, "Colorize the output DBT prints to the terminal")
	rootCmd.Flags().Bool("use-experimental-parser", false, "Enables experimental parsing features")
	rootCmd.Flags().Bool("version", false, "Show version information")
	rootCmd.Flags().Bool("warn-error", false, "If dbt would normally warn, instead raise an exception")
	rootCmd.Flags().String("warn-error-options", "", "If dbt would normally warn, instead raise an exception based on include/exclude configuration")

	rootCmd.PersistentFlags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.PersistentFlags().String("profiles-dir", "", "Which directory to look in for the profiles.yml file")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"log-format":         carapace.ActionValues("text", "json", "default"),
		"profiles-dir":       carapace.ActionDirectories(),
		"record-timing-info": carapace.ActionFiles(),
	})
}

func addProjectDirFlag(cmd *cobra.Command) {
	cmd.Flags().String("project-dir", "", "Which directory to look in for the dbt_project.yml file")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"project-dir": carapace.ActionDirectories(),
	})

	carapace.Gen(cmd).PreInvoke(func(cmd *cobra.Command, flag *pflag.Flag, action carapace.Action) carapace.Action {
		return action.Chdir(cmd.Flag("project-dir").Value.String())
	})
}

func addModelsFlag(cmd *cobra.Command) {
	cmd.Flags().StringSliceP("models", "m", nil, "Specify the nodes to include")

	cmd.Flag("models").Nargs = -1

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"models": dbt.ActionResources(dbt.ResourceOpts{Model: true}).FilterParts(), // TODO test this
	})
}

func addProfileFlag(cmd *cobra.Command) {
	cmd.Flags().String("profile", "", "Which profile to load. Overrides setting in dbt_project.yml")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"profile": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return dbt.ActionProfiles(cmd.Flag("profiles-dir").Value.String())
		}),
	})
}

func addSelectionFlags(cmd *cobra.Command) {
	cmd.Flags().StringSlice("exclude", nil, "Specify the models to exclude")
	cmd.Flags().StringSliceP("select", "s", nil, "Specify the nodes to include")

	cmd.Flag("exclude").Nargs = -1
	cmd.Flag("select").Nargs = -1

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"exclude": dbt.ActionResources(dbt.ResourceOpts{}.Default()).FilterParts(),
		"select":  dbt.ActionResources(dbt.ResourceOpts{}.Default()).FilterParts(),
	})
}
