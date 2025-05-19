package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parses the project and provides information on performance",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(parseCmd).Standalone()

	parseCmd.Flags().Bool("compile", false, "Compile")
	parseCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	parseCmd.Flags().String("log-path", "", "Configure the 'log-path'")
	parseCmd.Flags().Bool("no-version-check", false, "Skip ensuring dbt's version matches the one specified in the dbt_project.yml")
	parseCmd.Flags().StringP("target", "t", "", "Which target to load for the given profile")
	parseCmd.Flags().String("target-path", "", "Configure the 'target-path'")
	parseCmd.Flags().String("threads", "", "Specify number of threads to use while executing models")
	parseCmd.Flags().String("vars", "", "Supply variables to the project")
	parseCmd.Flags().Bool("write-manifest", false, "Write manifest")
	addProjectDirFlag(parseCmd)
	addProfileFlag(parseCmd)
	rootCmd.AddCommand(parseCmd)

	carapace.Gen(parseCmd).FlagCompletion(carapace.ActionMap{
		"log-path":    carapace.ActionFiles(),
		"target-path": carapace.ActionDirectories(),
	})
}
