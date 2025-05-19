package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var source_snapshotFreshnessCmd = &cobra.Command{
	Use:   "snapshot-freshness",
	Short: "Snapshots the current freshness of the project's sources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(source_snapshotFreshnessCmd).Standalone()

	source_snapshotFreshnessCmd.Flags().StringP("output", "o", "", "Specify the output path for the json report")
	source_snapshotFreshnessCmd.Flags().String("selector", "", "The selector name to use, as defined in selectors.yml")
	source_snapshotFreshnessCmd.Flags().String("state", "", "Use the given directory as the source for json files to compare with this project.")
	source_snapshotFreshnessCmd.Flags().StringP("target", "t", "", "Which target to load for the given profile")
	source_snapshotFreshnessCmd.Flags().String("threads", "", "Specify number of threads to use.")
	source_snapshotFreshnessCmd.Flags().String("vars", "", "Supply variables to the project.")
	addProjectDirFlag(source_snapshotFreshnessCmd)
	addSelectionFlags(source_snapshotFreshnessCmd)
	addProfileFlag(source_snapshotFreshnessCmd)
	sourceCmd.AddCommand(source_snapshotFreshnessCmd)

	carapace.Gen(source_snapshotFreshnessCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
		"state":  carapace.ActionDirectories(),
	})
}
