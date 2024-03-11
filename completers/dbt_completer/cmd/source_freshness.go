package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var source_freshnessCmd = &cobra.Command{
	Use:   "freshness",
	Short: "source freshness",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(source_freshnessCmd).Standalone()

	source_freshnessCmd.Flags().StringP("output", "o", "", "Specify the output path for the json report.")
	source_freshnessCmd.Flags().String("selector", "", "The selector name to use, as defined in selectors.yml")
	source_freshnessCmd.Flags().String("state", "", "Use the given directory as the source for json files to compare with this project.")
	source_freshnessCmd.Flags().StringP("target", "t", "", "Which target to load for the given profile")
	source_freshnessCmd.Flags().String("threads", "", "Specify number of threads to use.")
	source_freshnessCmd.Flags().String("vars", "", "Supply variables to the project.")
	addProjectDirFlag(source_freshnessCmd)
	addSelectionFlags(source_freshnessCmd)
	addProfileFlag(source_freshnessCmd)
	sourceCmd.AddCommand(source_freshnessCmd)

	carapace.Gen(source_freshnessCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
		"state":  carapace.ActionDirectories(),
	})
}
