package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "Execute snapshots defined in your project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snapshotCmd).Standalone()

	snapshotCmd.Flags().Bool("defer", false, "Defer to the state variable for resolving unselected nodes")
	snapshotCmd.Flags().Bool("favor-state", false, "Defer to the state variable for resolving unselected nodes")
	snapshotCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	snapshotCmd.Flags().Bool("no-defer", false, "Do not defer to the state variable for resolving unselected nodes")
	snapshotCmd.Flags().Bool("no-favor-state", false, "If defer is set, expect standard defer behaviour")
	snapshotCmd.Flags().String("selector", "", "The selector name to use")
	snapshotCmd.Flags().String("state", "", "Use the given directory as the source for json files to compare")
	snapshotCmd.Flags().StringP("target", "t", "", "Which target to load for the given profile")
	snapshotCmd.Flags().String("threads", "", "Specify number of threads to use while snapshotting tables")
	snapshotCmd.Flags().String("vars", "", "Supply variables to the project")
	addProjectDirFlag(snapshotCmd)
	addSelectionFlags(snapshotCmd)
	addModelsFlag(snapshotCmd)
	addProfileFlag(snapshotCmd)
	rootCmd.AddCommand(snapshotCmd)

	carapace.Gen(snapshotCmd).FlagCompletion(carapace.ActionMap{
		"state": carapace.ActionDirectories(),
	})
}
